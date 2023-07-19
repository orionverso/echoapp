package pipeline

import (
	"castor/construct/highlevel/repository"
	"castor/stack/computesave"
	"castor/stack/environment"
	"fmt"
	"log"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodestarconnections"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/pipelines"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type PipelineDeployFargateWriteToSaveObjectProps struct {
	awscdk.StackProps
	awscdk.StageProps
	pipelines.AddStageOpts
	FargateWriteToSaveObjectProps_FIRST_ENV  computesave.FargateWriteToSaveObjectProps
	FargateWriteToSaveObjectProps_SECOND_ENV computesave.FargateWriteToSaveObjectProps
	awscodestarconnections.CfnConnectionProps
	pipelines.ConnectionSourceOptions
	RepositoryProps                string
	BranchProps                    string
	CodeBuildSynthProps            pipelines.CodeBuildStepProps
	CodeBuildPushImageProps        pipelines.CodeBuildStepProps
	CodeBuildPushImageCrossAccount pipelines.CodeBuildStepProps
	pipelines.CodePipelineProps
	PromoteToProductionProps      pipelines.ManualApprovalStepProps
	PushImagePolicyStatementProps awsiam.PolicyStatementProps
	CrossAccountRole              awsiam.RoleProps
}

type pipelineDeployFargateWriteToSaveObject struct {
	awscdk.Stack
	codepipeline pipelines.CodePipeline
}

type PipelineDeployFargateWriteToSaveObject interface {
	awscdk.Stack
	CodePipeline() pipelines.CodePipeline
}

func NewPipelineDeployFargateWriteToSaveObject(scope constructs.Construct, id *string, props *PipelineDeployFargateWriteToSaveObjectProps) PipelineDeployFargateWriteToSaveObject {
	var sprops *PipelineDeployFargateWriteToSaveObjectProps = &PipelineDeployFargateWriteToSaveObjectProps{}

	if id == nil {
		log.Panicln("parameter id is required, but nil was provided")
	}

	if props != nil {
		sprops = props
	}

	stack := awscdk.NewStack(scope, id, &sprops.StackProps)

	conn := awscodestarconnections.NewCfnConnection(stack, jsii.String("CodeStarConnectionToGitHub"), &sprops.CfnConnectionProps)

	sprops.AddConnectionArn(conn.AttrConnectionArn())

	GithubRepository := pipelines.CodePipelineSource_Connection(&sprops.RepositoryProps, &sprops.BranchProps, &sprops.ConnectionSourceOptions)

	sprops.AddRemoteRepositoryToSynthStep(GithubRepository)

	template := pipelines.NewCodeBuildStep(jsii.String("SynthStep"), &sprops.CodeBuildSynthProps)

	sprops.AddTemplateToCodePipeline(template)

	pipe := pipelines.NewCodePipeline(stack, jsii.String("CodePipeline"), &sprops.CodePipelineProps)

	stagedev := awscdk.NewStage(stack, jsii.String("ComputeSaveStage"), &sprops.StageProps)

	computesave.NewFargateWriteToSaveObject(stagedev, jsii.String("ComputeSave"), &sprops.FargateWriteToSaveObjectProps_FIRST_ENV)

	stagedevdeployment := pipe.AddStage(stagedev, &sprops.AddStageOpts)

	pushimage := pipelines.NewCodeBuildStep(jsii.String("PushImageToEcr"), &sprops.CodeBuildPushImageProps)

	promotedecision := pipelines.NewManualApprovalStep(jsii.String("PromoteToProduction"), &sprops.PromoteToProductionProps)

	sprops.AddPostStepsToStackStep(0, stagedevdeployment, pushimage)

	stagedevdeployment.AddPost(promotedecision)
	// change environment
	sprops.StageProps.Env.Account = environment.StageProps_PROD.Env.Account
	sprops.StackProps.Env.Account = environment.StackProps_PROD.Env.Account

	stageprod := awscdk.NewStage(stack, jsii.String("ComputeSaveStage-Prod"), &sprops.StageProps)

	aux := awscdk.NewStack(stageprod, jsii.String("ComputeSaveAuxiliar"), &sprops.StackProps)

	awsiam.NewRole(aux, jsii.String("AssumeRoleToPushImageProdAccount"), &sprops.CrossAccountRole)

	computesave.NewFargateWriteToSaveObject(stageprod, jsii.String("ComputeSave-Prod"), &sprops.FargateWriteToSaveObjectProps_SECOND_ENV)

	stageproddeployment := pipe.AddStage(stageprod, &sprops.AddStageOpts)

	pushimagecross := pipelines.NewCodeBuildStep(jsii.String("PushImageToEcrProd"), &sprops.CodeBuildPushImageCrossAccount)

	sprops.AddPostStepsToStackStep(1, stageproddeployment, pushimagecross)

	var component PipelineDeployFargateWriteToSaveObject = &pipelineDeployFargateWriteToSaveObject{
		Stack:        stack,
		codepipeline: pipe,
	}

	return component
}

// PROPS
func (props *PipelineDeployFargateWriteToSaveObjectProps) AddConnectionArn(arn *string) {
	props.ConnectionSourceOptions.ConnectionArn = arn
}

func (props *PipelineDeployFargateWriteToSaveObjectProps) AddRemoteRepositoryToSynthStep(remoterepository pipelines.CodePipelineSource) {
	props.CodeBuildSynthProps.Input = remoterepository
}

func (props *PipelineDeployFargateWriteToSaveObjectProps) AddTemplateToCodePipeline(template pipelines.CodeBuildStep) {
	props.CodePipelineProps.Synth = template
}

func (props *PipelineDeployFargateWriteToSaveObjectProps) AddEnvironmentVariableToCodeBuildStep(key *string, value *string, step *pipelines.CodeBuildStepProps) {
	if step.BuildEnvironment == nil {
		var buildenvironment *awscodebuild.BuildEnvironment = &awscodebuild.BuildEnvironment{}
		step.BuildEnvironment = buildenvironment
	}
	var vars map[string]*awscodebuild.BuildEnvironmentVariable = *step.BuildEnvironment.EnvironmentVariables
	vars[*key] = &awscodebuild.BuildEnvironmentVariable{
		Value: *value,
	}
}

func (props *PipelineDeployFargateWriteToSaveObjectProps) AddPostStepsToStackStep(stackposition int, stagedeployment pipelines.StageDeployment, steps ...pipelines.Step) {
	var stacks []pipelines.StackDeployment = *stagedeployment.Stacks()
	var stack pipelines.StackDeployment = stacks[stackposition]
	stack.AddStackSteps(&[]pipelines.Step{}, &[]pipelines.Step{}, &steps)
}

// IMPLEMENTATION

func (mo *pipelineDeployFargateWriteToSaveObject) CodePipeline() pipelines.CodePipeline {
	return mo.codepipeline
}

// SETTINGS
// DEVELOPMENT
var PipelineDeployFargateWriteToSaveObjectProps_DEV PipelineDeployFargateWriteToSaveObjectProps = PipelineDeployFargateWriteToSaveObjectProps{
	StackProps: environment.StackProps_DEV,

	StageProps: environment.StageProps_DEV,

	AddStageOpts: pipelines.AddStageOpts{},

	FargateWriteToSaveObjectProps_FIRST_ENV:  computesave.FargateWriteToSaveObjectProps_DEV,
	FargateWriteToSaveObjectProps_SECOND_ENV: computesave.FargateWriteToSaveObjectProps_PROD,

	CfnConnectionProps: awscodestarconnections.CfnConnectionProps{
		ConnectionName: jsii.String("GithubConnection"),
		ProviderType:   jsii.String("GitHub"),
	},

	ConnectionSourceOptions: pipelines.ConnectionSourceOptions{
		TriggerOnPush: jsii.Bool(true),
	},

	RepositoryProps: "orionverso/echoapp",

	BranchProps: "main",

	CodeBuildSynthProps: pipelines.CodeBuildStepProps{
		Commands: jsii.Strings("npm install -g aws-cdk", "cdk synth"),
		BuildEnvironment: &awscodebuild.BuildEnvironment{
			EnvironmentVariables: &map[string]*awscodebuild.BuildEnvironmentVariable{
				"CDK_DEV_REGION": {
					Value: aws.ToString(environment.StackProps_DEV.Env.Region),
				},
				"CDK_DEV_ACCOUNT": {
					Value: aws.ToString(environment.StackProps_DEV.Env.Account),
				},
				"CDK_PROD_REGION": {
					Value: aws.ToString(environment.StackProps_PROD.Env.Region),
				},
				"CDK_PROD_ACCOUNT": {
					Value: aws.ToString(environment.StackProps_PROD.Env.Account),
				},
			},
		},
	},

	CodeBuildPushImageProps: pipelines.CodeBuildStepProps{
		Commands: jsii.Strings(
			"aws ecr get-login-password --region $CDK_REGION | docker login --username AWS --password-stdin $CDK_ACCOUNT.dkr.ecr.$CDK_REGION.amazonaws.com",
			"cd asset/docker/webserver",
			"docker build -t $REPOSITORY_NAME .",
			"docker tag $REPOSITORY_NAME:latest $CDK_ACCOUNT.dkr.ecr.$CDK_REGION.amazonaws.com/$REPOSITORY_NAME:latest",
			"docker push $CDK_ACCOUNT.dkr.ecr.$CDK_REGION.amazonaws.com/$REPOSITORY_NAME:latest",
		),

		BuildEnvironment: &awscodebuild.BuildEnvironment{
			Privileged: jsii.Bool(true), // Run Docker inside CodeBuild container
			EnvironmentVariables: &map[string]*awscodebuild.BuildEnvironmentVariable{
				"CDK_REGION": {
					Value: aws.ToString(environment.StackProps_DEV.Env.Region),
				},
				"CDK_ACCOUNT": {
					Value: aws.ToString(environment.StackProps_DEV.Env.Account),
				},
				"REPOSITORY_NAME": {
					Value: aws.ToString(repository.EcrRepoProps_DEV.RepositoryName),
				},
			},
		}, RolePolicyStatements: &[]awsiam.PolicyStatement{
			awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
				Actions:   jsii.Strings("ecr:GetAuthorizationToken"),
				Resources: jsii.Strings("*"),
			}),
			awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
				Actions: jsii.Strings(
					"ecr:GetAuthorizationToken",
					"ecr:BatchCheckLayerAvailability",
					"ecr:GetDownloadUrlForLayer",
					"ecr:GetRepositoryPolicy",
					"ecr:DescribeRepositories",
					"ecr:ListImages",
					"ecr:DescribeImages",
					"ecr:BatchGetImage",
					"ecr:InitiateLayerUpload",
					"ecr:UploadLayerPart",
					"ecr:CompleteLayerUpload",
					"ecr:PutImage",
				),
				Resources: jsii.Strings(fmt.Sprint("arn:aws:ecr:", *environment.StackProps_DEV.Env.Region, ":", *environment.StackProps_DEV.Env.Account, ":repository/echoapp-dev"), fmt.Sprint("arn:aws:ecr:", *environment.StackProps_DEV.Env.Region, ":", *environment.StackProps_DEV.Env.Account, ":repository/echoapp-dev/*")),
			}),
			awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
				Actions:   jsii.Strings("sts:AssumeRole"),
				Resources: jsii.Strings(fmt.Sprint("arn:aws:iam::", *environment.StackProps_PROD.Env.Account, ":role/", "pushimagecrossaccount")),
			}),
		},
	},

	CodePipelineProps: pipelines.CodePipelineProps{
		PipelineName:      jsii.String("DeployComputeSave"),
		CrossAccountKeys:  jsii.Bool(true),
		CodeBuildDefaults: &pipelines.CodeBuildOptions{},
	},

	CrossAccountRole: awsiam.RoleProps{
		RoleName:  jsii.String("pushimagecrossaccount"),
		AssumedBy: awsiam.NewAccountPrincipal(*environment.StackProps_DEV.Env.Account),
		InlinePolicies: &map[string]awsiam.PolicyDocument{
			"PushImageRole": awsiam.NewPolicyDocument(&awsiam.PolicyDocumentProps{
				Statements: &[]awsiam.PolicyStatement{
					awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
						Actions:   jsii.Strings("ecr:GetAuthorizationToken"),
						Resources: jsii.Strings("*"),
					}),
					awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
						Actions: jsii.Strings(
							"ecr:GetAuthorizationToken",
							"ecr:BatchCheckLayerAvailability",
							"ecr:GetDownloadUrlForLayer",
							"ecr:GetRepositoryPolicy",
							"ecr:DescribeRepositories",
							"ecr:ListImages",
							"ecr:DescribeImages",
							"ecr:BatchGetImage",
							"ecr:InitiateLayerUpload",
							"ecr:UploadLayerPart",
							"ecr:CompleteLayerUpload",
							"ecr:PutImage",
						),
						Resources: jsii.Strings(fmt.Sprint("arn:aws:ecr:", *environment.StackProps_PROD.Env.Region, ":", *environment.StackProps_PROD.Env.Account, ":repository/echoapp-prod"), fmt.Sprint("arn:aws:ecr:", *environment.StackProps_PROD.Env.Region, ":", *environment.StackProps_PROD.Env.Account, ":repository/echoapp-prod/*")),
					}),
				},
			}),
		},
	},
	CodeBuildPushImageCrossAccount: pipelines.CodeBuildStepProps{
		Commands: jsii.Strings(
			"cache=\"/tmp/creds\"",
			"aws sts assume-role --role-arn $PUSH_ROLE_ARN --role-session-name pushing-prod > $cache", // test
			"export AWS_ACCESS_KEY_ID=$(cat $cache | jq -r '.Credentials.AccessKeyId')",
			"export AWS_SECRET_ACCESS_KEY=$(cat $cache | jq -r '.Credentials.SecretAccessKey')",
			"export AWS_SESSION_TOKEN=$(cat $cache | jq -r '.Credentials.SessionToken')",
			"aws ecr get-login-password --region $CDK_REGION | docker login --username AWS --password-stdin $CDK_ACCOUNT.dkr.ecr.$CDK_REGION.amazonaws.com",
			"cd asset/docker/webserver",
			"docker build -t $REPOSITORY_NAME .",
			"docker tag $REPOSITORY_NAME:latest $CDK_ACCOUNT.dkr.ecr.$CDK_REGION.amazonaws.com/$REPOSITORY_NAME:latest",
			"docker push $CDK_ACCOUNT.dkr.ecr.$CDK_REGION.amazonaws.com/$REPOSITORY_NAME:latest",
		),

		BuildEnvironment: &awscodebuild.BuildEnvironment{
			Privileged: jsii.Bool(true), // Run Docker inside CodeBuild container
			EnvironmentVariables: &map[string]*awscodebuild.BuildEnvironmentVariable{
				"CDK_REGION": {
					Value: aws.ToString(environment.StackProps_PROD.Env.Region),
				},
				"CDK_ACCOUNT": {
					Value: aws.ToString(environment.StackProps_PROD.Env.Account),
				},
				"REPOSITORY_NAME": {
					Value: aws.ToString(repository.EcrRepoProps_PROD.RepositoryName),
				},
				"PUSH_ROLE_ARN": {
					Value: fmt.Sprint("arn:aws:iam::", *environment.StackProps_PROD.Env.Account, ":role/", "pushimagecrossaccount"),
				},
			},
		}, RolePolicyStatements: &[]awsiam.PolicyStatement{
			awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
				Actions:   jsii.Strings("sts:AssumeRole"),
				Resources: jsii.Strings(fmt.Sprint("arn:aws:iam::", *environment.StackProps_PROD.Env.Account, ":role/", "pushimagecrossaccount")),
			}),
		},
	},
}

// PRODUCTION
var PipelineDeployFargateWriteToSaveObjectProps_PROD PipelineDeployFargateWriteToSaveObjectProps = PipelineDeployFargateWriteToSaveObjectProps{
	StackProps: environment.StackProps_PROD,

	StageProps: environment.StageProps_PROD,

	AddStageOpts: pipelines.AddStageOpts{},

	FargateWriteToSaveObjectProps_FIRST_ENV: computesave.FargateWriteToSaveObjectProps_PROD,

	CfnConnectionProps: awscodestarconnections.CfnConnectionProps{
		ConnectionName: jsii.String("GithubConnection"),
		ProviderType:   jsii.String("GitHub"),
	},

	ConnectionSourceOptions: pipelines.ConnectionSourceOptions{
		TriggerOnPush: jsii.Bool(true),
	},

	RepositoryProps: "orionverso/echoapp",

	BranchProps: "main",

	CodeBuildSynthProps: pipelines.CodeBuildStepProps{
		Commands:         jsii.Strings("npm install -g aws-cdk", "cdk synth"),
		BuildEnvironment: &awscodebuild.BuildEnvironment{},
	},

	CodePipelineProps: pipelines.CodePipelineProps{
		PipelineName:      jsii.String("DeployComputeSave"),
		CrossAccountKeys:  jsii.Bool(true),
		CodeBuildDefaults: &pipelines.CodeBuildOptions{},
	},
}

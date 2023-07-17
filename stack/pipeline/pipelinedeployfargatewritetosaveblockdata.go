package pipeline

import (
	"castor/construct/highlevel/repository"
	"castor/stack/computesave"
	"castor/stack/environment"
	"log"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodestarconnections"
	"github.com/aws/aws-cdk-go/awscdk/v2/pipelines"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type PipelineDeployFargateWriteToSaveBlockDataProps struct {
	awscdk.StackProps
	awscdk.StageProps
	pipelines.AddStageOpts
	FargateWriteToSaveBlockDataProps_FIRST_ENV computesave.FargateWriteToSaveBlockDataProps
	awscodestarconnections.CfnConnectionProps
	pipelines.ConnectionSourceOptions
	RepositoryProps         string
	BranchProps             string
	CodeBuildSynthProps     pipelines.CodeBuildStepProps
	CodeBuildPushImageProps pipelines.CodeBuildStepProps
	pipelines.CodePipelineProps
}

type pipelineDeployFargateWriteToSaveBlockData struct {
	awscdk.Stack
	codepipeline pipelines.CodePipeline
}

type PipelineDeployFargateWriteToSaveBlockData interface {
	awscdk.Stack
	CodePipeline() pipelines.CodePipeline
}

func NewPipelineDeployFargateWriteToSaveBlockData(scope constructs.Construct, id *string, props *PipelineDeployFargateWriteToSaveBlockDataProps) PipelineDeployFargateWriteToSaveBlockData {
	var sprops *PipelineDeployFargateWriteToSaveBlockDataProps = &PipelineDeployFargateWriteToSaveBlockDataProps{}

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

	stage := awscdk.NewStage(stack, jsii.String("ComputeSaveStage"), &sprops.StageProps)

	computesave.NewFargateWriteToSaveBlockData(stage, jsii.String("ComputeSave"), &sprops.FargateWriteToSaveBlockDataProps_FIRST_ENV)

	pipelines.NewCodeBuildStep(jsii.String("PushImageToEcr"), &sprops.CodeBuildPushImageProps)

	var component PipelineDeployFargateWriteToSaveBlockData = &pipelineDeployFargateWriteToSaveBlockData{
		Stack:        stack,
		codepipeline: pipe,
	}

	return component
}

// PROPS
func (props *PipelineDeployFargateWriteToSaveBlockDataProps) AddConnectionArn(arn *string) {
	props.ConnectionSourceOptions.ConnectionArn = arn
}

func (props *PipelineDeployFargateWriteToSaveBlockDataProps) AddRemoteRepositoryToSynthStep(remoterepository pipelines.CodePipelineSource) {
	props.CodeBuildSynthProps.Input = remoterepository
}

func (props *PipelineDeployFargateWriteToSaveBlockDataProps) AddTemplateToCodePipeline(template pipelines.CodeBuildStep) {
	props.CodePipelineProps.Synth = template
}

func (props *PipelineDeployFargateWriteToSaveBlockDataProps) AddEnvironmentVariableToCodeBuildStep(step pipelines.CodeBuildStep, key *string, value *string) {
	var vars map[string]*awscodebuild.BuildEnvironmentVariable = *step.BuildEnvironment().EnvironmentVariables
	vars[aws.ToString(key)] = &awscodebuild.BuildEnvironmentVariable{
		Value: aws.ToString(value),
	}
}

// IMPLEMENTATION

func (mo *pipelineDeployFargateWriteToSaveBlockData) CodePipeline() pipelines.CodePipeline {
	return mo.codepipeline
}

// SETTINGS
// DEVELOPMENT
var PipelineDeployFargateWriteToSaveBlockDataProps_DEV PipelineDeployFargateWriteToSaveBlockDataProps = PipelineDeployFargateWriteToSaveBlockDataProps{
	StackProps: environment.StackProps_DEV,

	StageProps: environment.StageProps_DEV,

	AddStageOpts: pipelines.AddStageOpts{},

	FargateWriteToSaveBlockDataProps_FIRST_ENV: computesave.FargateWriteToSaveBlockDataProps_DEV,

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

	CodeBuildPushImageProps: pipelines.CodeBuildStepProps{
		Commands: jsii.Strings(
			"cd asset/docker/webserver",
			"cache=\"/tmp/creds\"",
			"aws sts assume-role --role-arn $PUSH_ROLE_ARN --role-session-name pushingimage > $cache", // test
			"export AWS_ACCESS_KEY_ID=$(cat $cache | jq -r '.Credentials.AccessKeyId')",
			"export AWS_SECRET_ACCESS_KEY=$(cat $cache | jq -r '.Credentials.SecretAccessKey')",
			"export AWS_SESSION_TOKEN=$(cat $cache | jq -r '.Credentials.SessionToken')",
			"aws ecr get-login-password --region $CDK_REGION | docker login --username AWS --password-stdin $CDK_ACCOUNT.dkr.ecr.$CDK_REGION.amazonaws.com",
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
				/*
					"PUSH_ROLE_ARN": &awscodebuild.BuildEnvironmentVariable{
						//Value: At runtime
					},
				*/
			},
		},
	},

	CodePipelineProps: pipelines.CodePipelineProps{
		PipelineName:      jsii.String("DeployComputeSave"),
		CrossAccountKeys:  jsii.Bool(true),
		CodeBuildDefaults: &pipelines.CodeBuildOptions{},
	},
}

// PRODUCTION
var PipelineDeployFargateWriteToSaveBlockDataProps_PROD PipelineDeployFargateWriteToSaveBlockDataProps = PipelineDeployFargateWriteToSaveBlockDataProps{
	StackProps: environment.StackProps_PROD,

	StageProps: environment.StageProps_PROD,

	AddStageOpts: pipelines.AddStageOpts{},

	FargateWriteToSaveBlockDataProps_FIRST_ENV: computesave.FargateWriteToSaveBlockDataProps_PROD,

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

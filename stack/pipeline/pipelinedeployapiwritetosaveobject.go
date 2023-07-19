package pipeline

import (
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

type PipelineDeployApiWriteToSaveObjectProps struct {
	awscdk.StackProps
	awscdk.StageProps
	pipelines.AddStageOpts
	ApiWriteToSaveObjectProps_FIRST_ENV  computesave.ApiWriteToSaveObjectProps
	ApiWriteToSaveObjectProps_SECOND_ENV computesave.ApiWriteToSaveObjectProps
	awscodestarconnections.CfnConnectionProps
	pipelines.ConnectionSourceOptions
	RepositoryProps                  string
	BranchProps                      string
	CodeBuildSynthProps              pipelines.CodeBuildStepProps
	PromoteToProductionDecisionProps pipelines.ManualApprovalStepProps
	pipelines.CodePipelineProps
}

type pipelineDeployApiWriteToSaveObject struct {
	awscdk.Stack
	codepipeline pipelines.CodePipeline
}

type PipelineDeployApiWriteToSaveObject interface {
	awscdk.Stack
	CodePipeline() pipelines.CodePipeline
}

func NewPipelineDeployApiWriteToSaveObject(scope constructs.Construct, id *string, props *PipelineDeployApiWriteToSaveObjectProps) PipelineDeployApiWriteToSaveObject {
	var sprops *PipelineDeployApiWriteToSaveObjectProps = &PipelineDeployApiWriteToSaveObjectProps{}

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

	computesave.NewApiWriteToSaveObject(stagedev, jsii.String("ComputeSave"), &sprops.ApiWriteToSaveObjectProps_FIRST_ENV)

	stagedevdeployment := pipe.AddStage(stagedev, &sprops.AddStageOpts)

	promotedecision := pipelines.NewManualApprovalStep(jsii.String("PromoteToProduction"), &sprops.PromoteToProductionDecisionProps)

	stagedevdeployment.AddPost(promotedecision)
	// change environment
	sprops.StageProps = environment.StageProps_PROD
	sprops.StackProps = environment.StackProps_PROD

	stageprod := awscdk.NewStage(stack, jsii.String("ComputeSaveStage-Prod"), &sprops.StageProps)

	computesave.NewApiWriteToSaveObject(stageprod, jsii.String("ComputeSave"), &sprops.ApiWriteToSaveObjectProps_SECOND_ENV)

	pipe.AddStage(stageprod, &sprops.AddStageOpts)

	var component PipelineDeployApiWriteToSaveObject = &pipelineDeployApiWriteToSaveObject{
		Stack:        stack,
		codepipeline: pipe,
	}

	return component
}

// PROPS
func (props *PipelineDeployApiWriteToSaveObjectProps) AddConnectionArn(arn *string) {
	props.ConnectionSourceOptions.ConnectionArn = arn
}

func (props *PipelineDeployApiWriteToSaveObjectProps) AddRemoteRepositoryToSynthStep(remoterepository pipelines.CodePipelineSource) {
	props.CodeBuildSynthProps.Input = remoterepository
}

func (props *PipelineDeployApiWriteToSaveObjectProps) AddTemplateToCodePipeline(template pipelines.CodeBuildStep) {
	props.CodePipelineProps.Synth = template
}

// IMPLEMENTATION

func (mo *pipelineDeployApiWriteToSaveObject) CodePipeline() pipelines.CodePipeline {
	return mo.codepipeline
}

// SETTINGS
// DEVELOPMENT
var PipelineDeployApiWriteToSaveObjectProps_DEV PipelineDeployApiWriteToSaveObjectProps = PipelineDeployApiWriteToSaveObjectProps{
	StackProps: environment.StackProps_DEV,

	StageProps: environment.StageProps_DEV,

	AddStageOpts: pipelines.AddStageOpts{},

	ApiWriteToSaveObjectProps_FIRST_ENV: computesave.ApiWriteToSaveObjectProps_DEV,

	ApiWriteToSaveObjectProps_SECOND_ENV: computesave.ApiWriteToSaveObjectProps_PROD,

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
		Commands: jsii.Strings("npm install -g aws-cdk", "cd asset/lambda/echo",
			"./compile.sh handler echolambda.go", "cd ../../../", "cdk synth"),
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

	PromoteToProductionDecisionProps: pipelines.ManualApprovalStepProps{},

	CodePipelineProps: pipelines.CodePipelineProps{
		PipelineName:      jsii.String("DeployComputeSave"),
		CrossAccountKeys:  jsii.Bool(true),
		CodeBuildDefaults: &pipelines.CodeBuildOptions{},
	},
}

// PRODUCTION
var PipelineDeployApiWriteToSaveObjectProps_PROD PipelineDeployApiWriteToSaveObjectProps = PipelineDeployApiWriteToSaveObjectProps{
	StackProps: environment.StackProps_PROD,

	StageProps: environment.StageProps_PROD,

	AddStageOpts: pipelines.AddStageOpts{},

	ApiWriteToSaveObjectProps_FIRST_ENV: computesave.ApiWriteToSaveObjectProps_PROD,

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
		Commands: jsii.Strings("npm install -g aws-cdk", "cd asset/lambda/echo",
			"./compile.sh handler echolambda.go", "cd ../../../", "cdk synth"),
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

	CodePipelineProps: pipelines.CodePipelineProps{
		PipelineName:      jsii.String("DeployComputeSave"),
		CrossAccountKeys:  jsii.Bool(true),
		CodeBuildDefaults: &pipelines.CodeBuildOptions{},
	},
}

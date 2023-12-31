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

type PipelineDeployApiWriteToSaveBlockDataProps struct {
	awscdk.StackProps
	awscdk.StageProps
	pipelines.AddStageOpts
	ApiWriteToSaveBlockDataProps_FIRST_ENV  computesave.ApiWriteToSaveBlockDataProps
	ApiWriteToSaveBlockDataProps_SECOND_ENV computesave.ApiWriteToSaveBlockDataProps
	awscodestarconnections.CfnConnectionProps
	pipelines.ConnectionSourceOptions
	RepositoryProps                  string
	BranchProps                      string
	CodeBuildSynthProps              pipelines.CodeBuildStepProps
	PromoteToProductionDecisionProps pipelines.ManualApprovalStepProps
	pipelines.CodePipelineProps
}

type pipelineDeployApiWriteToSaveBlockData struct {
	awscdk.Stack
	codepipeline pipelines.CodePipeline
}

type PipelineDeployApiWriteToSaveBlockData interface {
	awscdk.Stack
	CodePipeline() pipelines.CodePipeline
}

func NewPipelineDeployApiWriteToSaveBlockData(scope constructs.Construct, id *string, props *PipelineDeployApiWriteToSaveBlockDataProps) PipelineDeployApiWriteToSaveBlockData {
	var sprops *PipelineDeployApiWriteToSaveBlockDataProps = &PipelineDeployApiWriteToSaveBlockDataProps{}

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

	computesave.NewApiWriteToSaveBlockData(stagedev, jsii.String("ComputeSave"), &sprops.ApiWriteToSaveBlockDataProps_FIRST_ENV)

	stagedevdeployment := pipe.AddStage(stagedev, &sprops.AddStageOpts)

	promotedecision := pipelines.NewManualApprovalStep(jsii.String("PromoteToProduction"), &sprops.PromoteToProductionDecisionProps)

	stagedevdeployment.AddPost(promotedecision)
	// change environment
	sprops.StageProps = environment.StageProps_PROD
	sprops.StackProps = environment.StackProps_PROD

	stageprod := awscdk.NewStage(stack, jsii.String("ComputeSaveStage-Prod"), &sprops.StageProps)

	computesave.NewApiWriteToSaveBlockData(stageprod, jsii.String("ComputeSave"), &sprops.ApiWriteToSaveBlockDataProps_SECOND_ENV)

	pipe.AddStage(stageprod, &sprops.AddStageOpts)

	var component PipelineDeployApiWriteToSaveBlockData = &pipelineDeployApiWriteToSaveBlockData{
		Stack:        stack,
		codepipeline: pipe,
	}

	return component
}

// PROPS
func (props *PipelineDeployApiWriteToSaveBlockDataProps) AddConnectionArn(arn *string) {
	props.ConnectionSourceOptions.ConnectionArn = arn
}

func (props *PipelineDeployApiWriteToSaveBlockDataProps) AddRemoteRepositoryToSynthStep(remoterepository pipelines.CodePipelineSource) {
	props.CodeBuildSynthProps.Input = remoterepository
}

func (props *PipelineDeployApiWriteToSaveBlockDataProps) AddTemplateToCodePipeline(template pipelines.CodeBuildStep) {
	props.CodePipelineProps.Synth = template
}

// IMPLEMENTATION

func (mo *pipelineDeployApiWriteToSaveBlockData) CodePipeline() pipelines.CodePipeline {
	return mo.codepipeline
}

// SETTINGS
// DEVELOPMENT
var PipelineDeployApiWriteToSaveBlockDataProps_DEV PipelineDeployApiWriteToSaveBlockDataProps = PipelineDeployApiWriteToSaveBlockDataProps{
	StackProps: environment.StackProps_DEV,

	StageProps: environment.StageProps_DEV,

	AddStageOpts: pipelines.AddStageOpts{},

	ApiWriteToSaveBlockDataProps_FIRST_ENV: computesave.ApiWriteToSaveBlockDataProps_DEV,

	ApiWriteToSaveBlockDataProps_SECOND_ENV: computesave.ApiWriteToSaveBlockDataProps_PROD,

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
var PipelineDeployApiWriteToSaveBlockDataProps_PROD PipelineDeployApiWriteToSaveBlockDataProps = PipelineDeployApiWriteToSaveBlockDataProps{
	StackProps: environment.StackProps_PROD,

	StageProps: environment.StageProps_PROD,

	AddStageOpts: pipelines.AddStageOpts{},

	ApiWriteToSaveBlockDataProps_FIRST_ENV: computesave.ApiWriteToSaveBlockDataProps_PROD,

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

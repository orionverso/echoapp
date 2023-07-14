package pipeline

import (
	computesave "castor/stage/computesave/alfa"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodestarconnections"
	"github.com/aws/aws-cdk-go/awscdk/v2/pipelines"
	"github.com/aws/constructs-go/constructs/v10"
)

type DeployStackIds interface {
	Stack() *string
	CfnConnection() *string
	CodeBuildSynth() *string
	CodePipeline() *string
	StackCollection_FIRST_DEPLOY() computesave.StackCollectionIds
	StackCollection_SECOND_DEPLOY() computesave.StackCollectionIds
}

type DeployStackProps interface {
	Stack() *awscdk.StackProps
	CfnConnection() *awscodestarconnections.CfnConnectionProps
	ConnectionSourceOptions() *pipelines.ConnectionSourceOptions
	Repository() *string
	Branch() *string
	CodeBuildSynth() *pipelines.CodeBuildStepProps
	CodePipeline() *pipelines.CodePipelineProps
	// connections
	AddConnectionArn(*string)
	AddRemoteRepositoryToSynthStep(pipelines.CodePipelineSource)
	AddTemplateToCodePipeline(pipelines.CodeBuildStep)
	StackCollection_FIRST_DEPLOY() computesave.StackCollectionProps
	StackCollection_SECOND_DEPLOY() computesave.StackCollectionProps
}

type DeployStack interface{}

func NewDeployStack(scope constructs.Construct, id DeployStackIds, props DeployStackProps) DeployStack {
	var sprops DeployStackProps = &DeployModelProps_DEFAULT
	var sid DeployStackIds = &DeployModelIds_DEFAULT

	if props != nil {
		sprops = props
	}

	if id != nil {
		sid = id
	}

	stack := awscdk.NewStack(scope, sid.Stack(), sprops.Stack())

	conn := awscodestarconnections.NewCfnConnection(stack, id.CfnConnection(), sprops.CfnConnection())

	sprops.AddConnectionArn(conn.AttrConnectionArn())

	GithubRepository := pipelines.CodePipelineSource_Connection(sprops.Repository(), sprops.Branch(), sprops.ConnectionSourceOptions())

	sprops.AddRemoteRepositoryToSynthStep(GithubRepository)

	template := pipelines.NewCodeBuildStep(id.CodeBuildSynth(), sprops.CodeBuildSynth())

	sprops.AddTemplateToCodePipeline(template)

	pipe := pipelines.NewCodePipeline(stack, id.CodePipeline(), sprops.CodePipeline())

	FirstDeploy := computesave.NewStackCollection(stack, sid.StackCollection_FIRST_DEPLOY(), sprops.StackCollection_FIRST_DEPLOY()) // aca esta el error

	pipe.AddStage(FirstDeploy.Stage(), &pipelines.AddStageOpts{})

	SecondDeploy := computesave.NewStackCollection(stack, sid.StackCollection_SECOND_DEPLOY(), sprops.StackCollection_SECOND_DEPLOY())

	pipe.AddStage(SecondDeploy.Stage(), &pipelines.AddStageOpts{})

	var component DeployStack = &DeployModel{
		codepipeline: pipe,
	}

	return component
}

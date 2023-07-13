package pipeline

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodestarconnections"
	"github.com/aws/aws-cdk-go/awscdk/v2/pipelines"
	"github.com/aws/constructs-go/constructs/v10"
)

type DeployStackIds interface {
	Stack() *string
	CfnConnection() *string
	CodeBuildSynth() *string
	CodePipeline() *string
}

type DeployStackProps interface {
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

	stack := constructs.NewConstruct(scope, sid.Stack())

	conn := awscodestarconnections.NewCfnConnection(stack, id.CfnConnection(), sprops.CfnConnection())

	sprops.AddConnectionArn(conn.AttrConnectionArn())

	GithubRepository := pipelines.CodePipelineSource_Connection(sprops.Repository(), sprops.Branch(), sprops.ConnectionSourceOptions())

	sprops.AddRemoteRepositoryToSynthStep(GithubRepository)

	template := pipelines.NewCodeBuildStep(id.CodeBuildSynth(), sprops.CodeBuildSynth())

	sprops.AddTemplateToCodePipeline(template)

	pipe := pipelines.NewCodePipeline(stack, id.CodePipeline(), sprops.CodePipeline())

	var component DeployStack = &DeployModel{
		codepipeline: pipe,
	}

	return component
}

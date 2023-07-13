package pipeline

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodestarconnections"
	"github.com/aws/aws-cdk-go/awscdk/v2/pipelines"
	"github.com/aws/jsii-runtime-go"
)

type DeployModelIds struct {
	StackId          *string
	CfnConnectionId  *string
	CodeBuildSynthId *string
	CodePipelineId   *string
}

func (id *DeployModelIds) Stack() *string {
	return id.StackId
}

func (id *DeployModelIds) CfnConnection() *string {
	return id.CfnConnectionId
}

func (id *DeployModelIds) CodeBuildSynth() *string {
	return id.CodeBuildSynthId
}

func (id *DeployModelIds) CodePipeline() *string {
	return id.CodePipelineId
}

type DeployModelProps struct {
	*awscodestarconnections.CfnConnectionProps
	ConnectionSourceOptionsProps *pipelines.ConnectionSourceOptions
	RepositoryProps              *string
	BranchProps                  *string
	CodeBuildSynthProps          *pipelines.CodeBuildStepProps
	*pipelines.CodePipelineProps
}

func (props *DeployModelProps) CfnConnection() *awscodestarconnections.CfnConnectionProps {
	return props.CfnConnectionProps
}

func (props *DeployModelProps) ConnectionSourceOptions() *pipelines.ConnectionSourceOptions {
	return props.ConnectionSourceOptionsProps
}

func (props *DeployModelProps) Repository() *string {
	return props.RepositoryProps
}

func (props *DeployModelProps) Branch() *string {
	return props.BranchProps
}

func (props *DeployModelProps) CodeBuildSynth() *pipelines.CodeBuildStepProps {
	return props.CodeBuildSynthProps
}

func (props *DeployModelProps) CodePipeline() *pipelines.CodePipelineProps {
	return props.CodePipelineProps
}

func (props *DeployModelProps) AddConnectionArn(arn *string) {
	props.ConnectionSourceOptionsProps.ConnectionArn = arn
}

func (props *DeployModelProps) AddRemoteRepositoryToSynthStep(remoterepository pipelines.CodePipelineSource) {
	props.CodeBuildSynthProps.Input = remoterepository
}

func (props *DeployModelProps) AddTemplateToCodePipeline(template pipelines.CodeBuildStep) {
	props.CodePipelineProps.Synth = template
}

type DeployModel struct {
	codepipeline pipelines.CodePipeline
}

// SETTINGS
var DeployModelIds_DEFAULT DeployModelIds = DeployModelIds{
	StackId:          jsii.String("DeployModel-default"),
	CfnConnectionId:  jsii.String("CodeStarConnectionToGitHub-default"),
	CodeBuildSynthId: jsii.String("SynthStep-default"),
	CodePipelineId:   jsii.String("EchoModel-default"),
}

var DeployModelProps_DEFAULT DeployModelProps = DeployModelProps{
	CfnConnectionProps: &awscodestarconnections.CfnConnectionProps{
		ConnectionName: jsii.String("GithubConnection"),
		ProviderType:   jsii.String("GitHub"),
	},

	ConnectionSourceOptionsProps: &pipelines.ConnectionSourceOptions{
		TriggerOnPush: jsii.Bool(true),
	},

	RepositoryProps: jsii.String("orionverso/echoapp"),

	BranchProps: jsii.String("main"),

	CodeBuildSynthProps: &pipelines.CodeBuildStepProps{
		Commands:         jsii.Strings("npm install -g aws-cdk", "cdk synth"),
		BuildEnvironment: &awscodebuild.BuildEnvironment{},
	},

	CodePipelineProps: &pipelines.CodePipelineProps{
		PipelineName:      jsii.String("DeployFuncionalityModel-dev"),
		CrossAccountKeys:  jsii.Bool(true),
		CodeBuildDefaults: &pipelines.CodeBuildOptions{},
	},
}

package main

import (
	"castor/stack/computesave"

	"github.com/aws/aws-cdk-go/awscdk/v2"

	// "github.com/aws/aws-cdk-go/awscdk/v2/awssqs"

	"github.com/aws/jsii-runtime-go"
)

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)
	// Choose and uncomment one implementation

	// PIPELINES
	// pipeline.NewPipelineDeployFargateWriteToSaveBlockData(app, jsii.String("ComputeSavePipeline"), &pipeline.PipelineDeployFargateWriteToSaveBlockDataProps_DEV)
	// pipeline.NewPipelineDeployFargateWriteToSaveObject(app, jsii.String("ComputeSavePipeline"), &pipeline.PipelineDeployFargateWriteToSaveObjectProps_DEV)
	// pipeline.NewPipelineDeployApiWriteToSaveBlockData(app, jsii.String("ComputeSavePipeline"), &pipeline.PipelineDeployApiWriteToSaveBlockDataProps_DEV)
	// pipeline.NewPipelineDeployApiWriteToSaveObject(app, jsii.String("ComputeSavePipeline"), &pipeline.PipelineDeployApiWriteToSaveObjectProps_DEV)

	// STACKS DIRECTLY
	// computesave.NewFargateWriteToSaveBlockData(app, jsii.String("ComputeSavePipeline"), &computesave.FargateWriteToSaveBlockDataProps_DEV)
	// computesave.NewFargateWriteToSaveObject(app, jsii.String("ComputeSavePipeline"), &computesave.FargateWriteToSaveObjectProps_DEV)
	//	computesave.NewApiWriteToSaveBlockData(app, jsii.String("ComputeSavePipeline"), &computesave.ApiWriteToSaveBlockDataProps_DEV)
	computesave.NewApiWriteToSaveObject(app, jsii.String("ComputeSavePipeline"), &computesave.ApiWriteToSaveObjectProps_DEV)

	app.Synth(nil)
}

package main

import (
	"castor/stack/pipeline"

	"github.com/aws/aws-cdk-go/awscdk/v2"

	// "github.com/aws/aws-cdk-go/awscdk/v2/awssqs"

	"github.com/aws/jsii-runtime-go"
)

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	pipeline.NewPipelineDeployApiWriteToSaveBlockData(app, jsii.String("ComputeSavePipeline"), &pipeline.PipelineDeployApiWriteToSaveBlockDataProps_DEV)

	app.Synth(nil)
}

package pipeline

import (
	"fmt"
	"log"
	"testing"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	"github.com/aws/jsii-runtime-go"
)

// DEFAULT
func TestSynthPipelineDeployFargateWriteToSaveBlockData_DEFAULT(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Fail()

			log.Println("THE COMPONENT IS NOT SYNTHEABLE")
			fmt.Println("----CHECK CDK ERROR---")
			fmt.Println(err)
			fmt.Println("----CHECK CDK ERROR---")
			// debug.PrintStack() //+info
		} else {
			log.Println("THE COMPONENT IS SYNTHEABLE")
		}
	}()

	defer jsii.Close()

	stack := awscdk.NewStack(nil, nil, nil)

	NewPipelineDeployFargateWriteToSaveBlockData(stack, jsii.String("TestSynth_DEFAULT"), nil)

	assertions.Template_FromStack(stack, nil)
}

// DEVELOPMENT
func TestSynthPipelineDeployFargateWriteToSaveBlockData_DEV(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Fail()

			log.Println("THE COMPONENT IS NOT SYNTHEABLE")
			fmt.Println("----CHECK CDK ERROR---")
			fmt.Println(err)
			fmt.Println("----CHECK CDK ERROR---")
			// debug.PrintStack() //+info
		} else {
			log.Println("THE COMPONENT IS SYNTHEABLE")
		}
	}()

	defer jsii.Close()

	stack := awscdk.NewStack(nil, nil, nil)

	NewPipelineDeployFargateWriteToSaveBlockData(stack, jsii.String("TestSynth_DEV"), &PipelineDeployFargateWriteToSaveBlockDataProps_DEV)

	assertions.Template_FromStack(stack, &TemplateParsingOption_DEV)
}

// PRODUCTION
func TestSynthPipelineDeployFargateWriteToSaveBlockData_PROD(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Fail()

			log.Println("THE COMPONENT IS NOT SYNTHEABLE")
			fmt.Println("----CHECK CDK ERROR---")
			fmt.Println(err)
			fmt.Println("----CHECK CDK ERROR---")
			// debug.PrintStack() //+info
		} else {
			log.Println("THE COMPONENT IS SYNTHEABLE")
		}
	}()

	defer jsii.Close()

	stack := awscdk.NewStack(nil, nil, nil)

	NewPipelineDeployFargateWriteToSaveBlockData(stack, jsii.String("TestSynth_PROD"), &PipelineDeployFargateWriteToSaveBlockDataProps_PROD)

	assertions.Template_FromStack(stack, &TemplateParsingOption_PROD)
}

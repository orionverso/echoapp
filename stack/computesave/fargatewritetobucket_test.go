package computesave

import (
	"fmt"
	"log"
	"testing"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	"github.com/aws/jsii-runtime-go"
)

// DEFAULT
func TestSynthFargateWriteToBucket_DEFAULT(t *testing.T) {
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

	NewFargateWriteToBucket(stack, jsii.String("TestSynth_DEFAULT"), nil)

	assertions.Template_FromStack(stack, nil)
}

// DEVELOPMENT
func TestSynthFargateWriteToBucket_DEV(t *testing.T) {
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

	NewFargateWriteToBucket(stack, jsii.String("TestSynth_DEV"), &FargateWriteToBucketProps_DEV)

	assertions.Template_FromStack(stack, &TemplateParsingOption_DEV)
}

var TemplateParsingOption_DEV assertions.TemplateParsingOptions = assertions.TemplateParsingOptions{
	SkipCyclicalDependenciesCheck: jsii.Bool(false),
}

// PRODUCTION
func TestSynthFargateWriteToBucket_PROD(t *testing.T) {
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

	NewFargateWriteToBucket(stack, jsii.String("TestSynth_PROD"), &FargateWriteToBucketProps_PROD)

	assertions.Template_FromStack(stack, &TemplateParsingOption_PROD)
}

var TemplateParsingOption_PROD assertions.TemplateParsingOptions = assertions.TemplateParsingOptions{
	SkipCyclicalDependenciesCheck: jsii.Bool(false),
}
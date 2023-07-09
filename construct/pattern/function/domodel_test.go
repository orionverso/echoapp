package function

import (
	"fmt"
	"log"
	"testing"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	"github.com/aws/jsii-runtime-go"
)

// MODEL
func TestSynthDoModel_DEFAULT(t *testing.T) {
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

	NewDoAction(stack, &DoModelIds_DEFAULT, &DoModelProps_DEFAULT)

	assertions.Template_FromStack(stack, &TemplateParsingOption_DEFAULT)
}

var TemplateParsingOption_DEFAULT assertions.TemplateParsingOptions = assertions.TemplateParsingOptions{
	SkipCyclicalDependenciesCheck: jsii.Bool(false),
}

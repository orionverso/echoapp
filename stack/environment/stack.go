package environment

import (
	"fmt"
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

var StackProps_DEV awscdk.StackProps = awscdk.StackProps{
	Env: DevEnv(),
}

var StackProps_PROD awscdk.StackProps = awscdk.StackProps{
	Env: ProdEnv(),
}

func ProdEnv() *awscdk.Environment {
	region := os.Getenv("CDK_PROD_REGION")
	account := os.Getenv("CDK_PROD_ACCOUNT")

	fmt.Println("PROD_REGION", region)
	fmt.Println("PROD_ACCOUNT", account)

	return &awscdk.Environment{
		Region:  jsii.String(region),
		Account: jsii.String(account),
	}
}

func DevEnv() *awscdk.Environment {
	region := os.Getenv("CDK_DEV_REGION")
	account := os.Getenv("CDK_DEV_ACCOUNT")

	fmt.Println("DEV_REGION", region)
	fmt.Println("DEV_ACCOUNT", account)

	return &awscdk.Environment{
		Region:  jsii.String(region),
		Account: jsii.String(account),
	}
}

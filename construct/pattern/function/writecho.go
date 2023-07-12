package function

import (
	"fmt"
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/jsii-runtime-go"
)

type WriteEchoIds struct {
	DoModelIds
}

type WriteEchoProps struct {
	DoModelProps
}

// SETTINGS
// DEVELOPMENT
var WriteEchoIds_DEV WriteEchoIds = WriteEchoIds{
	DoModelIds: DoModelIds{
		ConstructId:    jsii.String("WriteEcho-dev"),
		FunctionId:     jsii.String("function-dev"),
		SuccessQueueId: jsii.String("successqueue-dev"),
		FailureQueueId: jsii.String("failurequeue-dev"),
	},
}

var WriteEchoProps_DEV WriteEchoProps = WriteEchoProps{
	DoModelProps: DoModelProps{
		FunctionProps: &awslambda.FunctionProps{
			Runtime:      awslambda.Runtime_GO_1_X(),
			Handler:      jsii.String("handler"),
			Code:         awslambda.Code_FromAsset(jsii.String(fmt.Sprint(os.Getenv("ASSET_DIR"), "/echo/handler.zip")), nil),
			FunctionName: jsii.String("WriteEcho-dev"),
			Description:  jsii.String("This function write the echo request from api gateway to storage solution"),
			Architecture: awslambda.Architecture_X86_64(),
			MemorySize:   jsii.Number[float64](128),
			Timeout:      awscdk.Duration_Seconds(jsii.Number[float64](2)),
		},
		SuccessQueueProps: &awssqs.QueueProps{
			QueueName:     jsii.String("WriteEchoSuccessQueue"),
			RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
		},
		FailureQueueProps: &awssqs.QueueProps{
			QueueName:     jsii.String("WriteEchoFailureQueue"),
			RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
		},

		AllowReceiveMessageFromFunctionProps: &awsiam.PolicyStatementProps{
			Actions: jsii.Strings(
				"sqs:GetQueueAttributes",
				"sqs:GetQueueUrl",
				"sqs:SendMessage"),
			Effect: awsiam.Effect_ALLOW,
			// Principals: At runtime ,
			// Resource: At runtime
		},
	},
}

// PRODUCTION
var WriteEchoIds_PROD WriteEchoIds = WriteEchoIds{
	DoModelIds: DoModelIds{
		ConstructId:    jsii.String("WriteEcho-prod"),
		FunctionId:     jsii.String("function-prod"),
		SuccessQueueId: jsii.String("successqueue-prod"),
		FailureQueueId: jsii.String("failurequeue-prod"),
	},
}

var WriteEchoProps_PROD WriteEchoProps = WriteEchoProps{
	DoModelProps: DoModelProps{
		FunctionProps: &awslambda.FunctionProps{
			Runtime:      awslambda.Runtime_GO_1_X(),
			Handler:      jsii.String("handler"),
			Code:         awslambda.Code_FromAsset(jsii.String(fmt.Sprint(os.Getenv("ASSET_DIR"), "/echo/handler.zip")), nil),
			FunctionName: jsii.String("WriteEcho-prod"),
			Description:  jsii.String("This function write the echo request from api gateway to storage solution"),
			Architecture: awslambda.Architecture_X86_64(),
			MemorySize:   jsii.Number[float64](512),
			Timeout:      awscdk.Duration_Seconds(jsii.Number[float64](5)),
		},
		SuccessQueueProps: &awssqs.QueueProps{
			QueueName:     jsii.String("WriteEchoSuccessQueue"),
			RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
		},
		FailureQueueProps: &awssqs.QueueProps{
			QueueName:     jsii.String("WriteEchoFailureQueue"),
			RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
		},

		AllowReceiveMessageFromFunctionProps: &awsiam.PolicyStatementProps{
			Actions: jsii.Strings(
				"sqs:GetQueueAttributes",
				"sqs:GetQueueUrl",
				"sqs:SendMessage"),
			Effect: awsiam.Effect_ALLOW,
			// Principals: At runtime ,
			// Resource: At runtime
		},
	},
}

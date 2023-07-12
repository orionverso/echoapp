package lambdarestapi

import (
	"castor/construct/pattern/function"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/jsii-runtime-go"
)

type ReceiveEchoWriteEchoIds struct {
	ReceiveRequestDoActionModelIds
}

type ReceiveEchoWriteEchoProps struct {
	ReceiveRequestDoActionModelProps
}

// SETTINGS
// DEVELOPMENT
var ReceiveEchoWriteEchoIds_DEV ReceiveEchoWriteEchoIds = ReceiveEchoWriteEchoIds{
	ReceiveRequestDoActionModelIds: ReceiveRequestDoActionModelIds{
		ConstructId:     jsii.String("ReceiveEchoWriteEcho-dev"),
		LambdaRestApiId: jsii.String("lambdarestapi-dev"),
		DoActionIds:     &function.WriteEchoIds_DEV,
		LogGroupId:      jsii.String("loggroup-dev"),
	},
}

var ReceiveEchoWriteEchoProps_DEV ReceiveEchoWriteEchoProps = ReceiveEchoWriteEchoProps{
	ReceiveRequestDoActionModelProps: ReceiveRequestDoActionModelProps{
		LambdaRestApiProps: &awsapigateway.LambdaRestApiProps{
			CloudWatchRole: jsii.Bool(true),
			Description:    jsii.String("This is a apigateway with lambda proxy. The api will receive the body request and offer to proxy"),
			EndpointTypes:  &[]awsapigateway.EndpointType{awsapigateway.EndpointType_REGIONAL},
			FailOnWarnings: jsii.Bool(true),
			RestApiName:    jsii.String("EchoApi"),
			DeployOptions: &awsapigateway.StageOptions{
				CachingEnabled: jsii.Bool(true),
				MetricsEnabled: jsii.Bool(true),
				LoggingLevel:   awsapigateway.MethodLoggingLevel_ERROR,
				// AccessLogDestination: At runtime,
			},
		},
		LogGroupProps: &awslogs.LogGroupProps{
			LogGroupName:  jsii.String("EchoApiLogs"),
			RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
		},
		DoActionProps: &function.WriteEchoProps_DEV,
	},
}

// PRODUCTION
var ReceiveEchoWriteEchoIds_PROD ReceiveEchoWriteEchoIds = ReceiveEchoWriteEchoIds{
	ReceiveRequestDoActionModelIds: ReceiveRequestDoActionModelIds{
		ConstructId:     jsii.String("ReceiveEchoWriteEcho-prod"),
		LambdaRestApiId: jsii.String("lambdarestapi-prod"),
		DoActionIds:     &function.WriteEchoIds_PROD,
		LogGroupId:      jsii.String("loggroup-prod"),
	},
}

var ReceiveEchoWriteEchoProps_PROD ReceiveEchoWriteEchoProps = ReceiveEchoWriteEchoProps{
	ReceiveRequestDoActionModelProps: ReceiveRequestDoActionModelProps{
		LambdaRestApiProps: &awsapigateway.LambdaRestApiProps{
			CloudWatchRole: jsii.Bool(true),
			Description:    jsii.String("This is a apigateway with lambda proxy. The api will receive the body request and offer to proxy"),
			EndpointTypes:  &[]awsapigateway.EndpointType{awsapigateway.EndpointType_REGIONAL},
			FailOnWarnings: jsii.Bool(true),
			RestApiName:    jsii.String("EchoApi"),
			DeployOptions: &awsapigateway.StageOptions{
				CachingEnabled: jsii.Bool(true),
				MetricsEnabled: jsii.Bool(true),
				LoggingLevel:   awsapigateway.MethodLoggingLevel_ERROR,
				// AccessLogDestination: At runtime,
			},
		},
		LogGroupProps: &awslogs.LogGroupProps{
			LogGroupName:  jsii.String("EchoApiLogs"),
			RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
		},
		DoActionProps: &function.WriteEchoProps_PROD,
	},
}

package lambdarestapi

import (
	"castor/construct/pattern/function"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/jsii-runtime-go"
)

type ReceiveEchoIds struct {
	ReceiveModelIds
}

type ReceiveEchoProps struct {
	ReceiveModelProps
}

// SETTINGS
// DEVELOPMENT
var ReceiveEchoIds_DEV ReceiveEchoIds = ReceiveEchoIds{
	ReceiveModelIds: ReceiveModelIds{
		ConstructId:     jsii.String("ReceiveEcho-resource-receivemodel-default"),
		LambdaRestApiId: jsii.String("ReceiveEcho-resource-lambdarestapi-default"),
		DoActionIds:     &function.WriteEchoIds_DEV,
		LogGroupId:      jsii.String("ReceiveEcho-resource-loggroup-default"),
	},
}

var ReceiveEchoProps_DEV ReceiveEchoProps = ReceiveEchoProps{
	ReceiveModelProps: ReceiveModelProps{
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
var ReceiveEchoIds_PROD ReceiveEchoIds = ReceiveEchoIds{
	ReceiveModelIds: ReceiveModelIds{
		ConstructId:     jsii.String("ReceiveEcho-resource-receivemodel-default"),
		LambdaRestApiId: jsii.String("ReceiveEcho-resource-lambdarestapi-default"),
		DoActionIds:     &function.WriteEchoIds_PROD,
		LogGroupId:      jsii.String("ReceiveEcho-resource-loggroup-default"),
	},
}

var ReceiveEchoProps_PROD ReceiveEchoProps = ReceiveEchoProps{
	ReceiveModelProps: ReceiveModelProps{
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

package lambdarestapi

import (
	"castor/construct/pattern/function"
	"log"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type ApiGatewayWithLambdaProxyIntegratedProps struct {
	awsapigateway.LambdaRestApiProps
	awslogs.LogGroupProps
	function.FunctionWithSqsDestinationsProps
}

type apiGatewayWithLambdaProxyIntegrated struct {
	constructs.Construct
	lambdarestapi          awsapigateway.LambdaRestApi
	function               function.FunctionWithSqsDestinations
	loggroup               awslogs.LogGroup
	loggrouplogdestination awsapigateway.LogGroupLogDestination
}

type ApiGatewayWithLambdaProxyIntegrated interface {
	LambdaRestApi() awsapigateway.LambdaRestApi
	FunctionWithSqsDestinations() function.FunctionWithSqsDestinations
	LogGroup() awslogs.LogGroup
	LogGroupLogDestination() awsapigateway.LogGroupLogDestination
}

func NewApiGatewayWithLambdaProxyIntegrated(scope constructs.Construct, id *string, props *ApiGatewayWithLambdaProxyIntegratedProps) ApiGatewayWithLambdaProxyIntegrated {
	var sprops *ApiGatewayWithLambdaProxyIntegratedProps = &ApiGatewayWithLambdaProxyIntegratedProps{}

	if id == nil {
		log.Panicln("parameter id is required, but nil was provided")
	}

	if props != nil {
		sprops = props
	}

	this := constructs.NewConstruct(scope, id)

	fn := function.NewFunctionWithSqsDestinations(this, jsii.String("FunctionWithSqsDestinations"), &sprops.FunctionWithSqsDestinationsProps)

	loggroup := awslogs.NewLogGroup(this, jsii.String("LogGroup"), &sprops.LogGroupProps)
	loggrouplogdestination := awsapigateway.NewLogGroupLogDestination(loggroup)

	sprops.AddHandlerToLambdaRestApi(fn.Function())
	sprops.AddAccessLogDestinationToLambdaRestApi(loggrouplogdestination)

	resource := awsapigateway.NewLambdaRestApi(this, jsii.String("ApiGatewayWithLambdaProxy"), &sprops.LambdaRestApiProps)

	var component ApiGatewayWithLambdaProxyIntegrated = &apiGatewayWithLambdaProxyIntegrated{
		Construct:              this,
		lambdarestapi:          resource,
		function:               fn,
		loggroup:               loggroup,
		loggrouplogdestination: loggrouplogdestination,
	}

	return component
}

// PROPS
func (props *ApiGatewayWithLambdaProxyIntegratedProps) AddHandlerToLambdaRestApi(fn awslambda.Function) {
	props.LambdaRestApiProps.Handler = fn
}

func (props *ApiGatewayWithLambdaProxyIntegratedProps) AddAccessLogDestinationToLambdaRestApi(log awsapigateway.LogGroupLogDestination) {
	props.LambdaRestApiProps.DeployOptions.AccessLogDestination = log
}

// IMPLEMENTATION
func (mo *apiGatewayWithLambdaProxyIntegrated) LambdaRestApi() awsapigateway.LambdaRestApi {
	return mo.lambdarestapi
}

func (mo *apiGatewayWithLambdaProxyIntegrated) FunctionWithSqsDestinations() function.FunctionWithSqsDestinations {
	return mo.function
}

func (mo *apiGatewayWithLambdaProxyIntegrated) LogGroup() awslogs.LogGroup {
	return mo.loggroup
}

func (mo *apiGatewayWithLambdaProxyIntegrated) LogGroupLogDestination() awsapigateway.LogGroupLogDestination {
	return mo.loggrouplogdestination
}

// SETTINGS
// DEVELOPMENT
var ApiGatewayWithLambdaProxyIntegratedProps_DEV ApiGatewayWithLambdaProxyIntegratedProps = ApiGatewayWithLambdaProxyIntegratedProps{
	LambdaRestApiProps: awsapigateway.LambdaRestApiProps{
		CloudWatchRole: jsii.Bool(true),
		Description:    jsii.String("This is a apigateway with lambda proxy. The api will receive the body request and offer to proxy"),
		EndpointTypes:  &[]awsapigateway.EndpointType{awsapigateway.EndpointType_REGIONAL},
		FailOnWarnings: jsii.Bool(true),
		RestApiName:    jsii.String("EchoApi"),
		DeployOptions: &awsapigateway.StageOptions{
			MetricsEnabled: jsii.Bool(true),
			LoggingLevel:   awsapigateway.MethodLoggingLevel_ERROR,
			StageName:      jsii.String("dev"),
			// AccessLogDestination: At runtime,
		},
	},
	LogGroupProps: awslogs.LogGroupProps{
		LogGroupName:  jsii.String("EchoApiLogs"),
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	},
	FunctionWithSqsDestinationsProps: function.FunctionWithSqsDestinationsProps_DEV,
}

// PRODUCTION
var ApiGatewayWithLambdaProxyIntegratedProps_PROD ApiGatewayWithLambdaProxyIntegratedProps = ApiGatewayWithLambdaProxyIntegratedProps{
	LambdaRestApiProps: awsapigateway.LambdaRestApiProps{
		CloudWatchRole: jsii.Bool(true),
		Description:    jsii.String("This is a apigateway with lambda proxy. The api will receive the body request and offer to proxy"),
		EndpointTypes:  &[]awsapigateway.EndpointType{awsapigateway.EndpointType_REGIONAL},
		FailOnWarnings: jsii.Bool(true),
		RestApiName:    jsii.String("EchoApi"),
		DeployOptions: &awsapigateway.StageOptions{
			CachingEnabled: jsii.Bool(true),
			MetricsEnabled: jsii.Bool(true),
			LoggingLevel:   awsapigateway.MethodLoggingLevel_ERROR,
			StageName:      jsii.String("prod"),
			// AccessLogDestination: At runtime,
		},
	},
	LogGroupProps: awslogs.LogGroupProps{
		LogGroupName:  jsii.String("EchoApiLogs"),
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	},
	FunctionWithSqsDestinationsProps: function.FunctionWithSqsDestinationsProps_PROD,
}

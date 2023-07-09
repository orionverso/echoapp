package lambdarestapi

import (
	"castor/construct/pattern/function"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type ReceiveModelIds struct {
	ConstructId     *string
	LambdaRestApiId *string
	LogGroupId      *string
	function.DoActionIds
}

func (id *ReceiveModelIds) Construct() *string {
	return id.ConstructId
}

func (id *ReceiveModelIds) LambdaRestApi() *string {
	return id.LambdaRestApiId
}

func (id *ReceiveModelIds) DoAction() function.DoActionIds {
	return id.DoActionIds
}

func (id *ReceiveModelIds) LogGroup() *string {
	return id.LogGroupId
}

type ReceiveModelProps struct {
	*awsapigateway.LambdaRestApiProps
	*awslogs.LogGroupProps
	function.DoActionProps
}

func (props *ReceiveModelProps) LambdaRestApi() *awsapigateway.LambdaRestApiProps {
	return props.LambdaRestApiProps
}

func (props *ReceiveModelProps) DoAction() function.DoActionProps {
	return props.DoActionProps
}

func (props *ReceiveModelProps) LogGroup() *awslogs.LogGroupProps {
	return props.LogGroupProps
}

// connections
func (props *ReceiveModelProps) AddHandlerToLambdaRestApi(fn awslambda.Function, api *awsapigateway.LambdaRestApiProps) {
	api.Handler = fn
}

func (props *ReceiveModelProps) AddAccessLogDestinationToLambdaRestApi(log awsapigateway.LogGroupLogDestination, api *awsapigateway.LambdaRestApiProps) {
	api.DeployOptions.AccessLogDestination = log
}

type ReceiveModel struct {
	constructs.Construct
	lambdarestapi awsapigateway.LambdaRestApi
}

func (mo *ReceiveModel) LambdaRestApi() awsapigateway.LambdaRestApi {
	return mo.lambdarestapi
}

// SETTINGS
var ReceiveModelIds_DEFAULT ReceiveModelIds = ReceiveModelIds{
	ConstructId:     jsii.String("MODEL-resource-receivemodel-default"),
	LambdaRestApiId: jsii.String("MODEL-resource-lambdarestapi-default"),
	DoActionIds:     &function.DoModelIds_DEFAULT,
	LogGroupId:      jsii.String("MODEL-resource-loggroup-default"),
}

var ReceiveModelProps_DEFAULT ReceiveModelProps = ReceiveModelProps{
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
	DoActionProps: &function.DoModelProps_DEFAULT,
}

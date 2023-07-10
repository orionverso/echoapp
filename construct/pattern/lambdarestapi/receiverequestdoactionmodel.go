package lambdarestapi

import (
	"castor/construct/pattern/function"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type ReceiveRequestDoActionModelIds struct {
	ConstructId     *string
	LambdaRestApiId *string
	LogGroupId      *string
	function.DoActionIds
}

func (id *ReceiveRequestDoActionModelIds) Construct() *string {
	return id.ConstructId
}

func (id *ReceiveRequestDoActionModelIds) LambdaRestApi() *string {
	return id.LambdaRestApiId
}

func (id *ReceiveRequestDoActionModelIds) DoAction() function.DoActionIds {
	return id.DoActionIds
}

func (id *ReceiveRequestDoActionModelIds) LogGroup() *string {
	return id.LogGroupId
}

type ReceiveRequestDoActionModelProps struct {
	*awsapigateway.LambdaRestApiProps
	*awslogs.LogGroupProps
	function.DoActionProps
}

func (props *ReceiveRequestDoActionModelProps) LambdaRestApi() *awsapigateway.LambdaRestApiProps {
	return props.LambdaRestApiProps
}

func (props *ReceiveRequestDoActionModelProps) DoAction() function.DoActionProps {
	return props.DoActionProps
}

func (props *ReceiveRequestDoActionModelProps) LogGroup() *awslogs.LogGroupProps {
	return props.LogGroupProps
}

// connections
func (props *ReceiveRequestDoActionModelProps) AddHandlerToLambdaRestApi(fn awslambda.Function, api *awsapigateway.LambdaRestApiProps) {
	api.Handler = fn
}

func (props *ReceiveRequestDoActionModelProps) AddAccessLogDestinationToLambdaRestApi(log awsapigateway.LogGroupLogDestination, api *awsapigateway.LambdaRestApiProps) {
	api.DeployOptions.AccessLogDestination = log
}

type ReceiveRequestDoActionModel struct {
	constructs.Construct
	lambdarestapi awsapigateway.LambdaRestApi
	doaction      awsiam.IRole
}

func (mo *ReceiveRequestDoActionModel) LambdaRestApi() awsapigateway.LambdaRestApi {
	return mo.lambdarestapi
}

func (mo *ReceiveRequestDoActionModel) DoActionRole() awsiam.IRole {
	return mo.doaction
}

// SETTINGS
var ReceiveRequestDoActionModelIds_DEFAULT ReceiveRequestDoActionModelIds = ReceiveRequestDoActionModelIds{
	ConstructId:     jsii.String("MODEL-resource-receivemodel-default"),
	LambdaRestApiId: jsii.String("MODEL-resource-lambdarestapi-default"),
	DoActionIds:     &function.DoModelIds_DEFAULT,
	LogGroupId:      jsii.String("MODEL-resource-loggroup-default"),
}

var ReceiveRequestDoActionModelProps_DEFAULT ReceiveRequestDoActionModelProps = ReceiveRequestDoActionModelProps{
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

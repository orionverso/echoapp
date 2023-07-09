package lambdarestapi

import (
	"castor/construct/pattern/function"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/constructs-go/constructs/v10"
)

type ReceiveRequestIds interface {
	Construct() *string
	LambdaRestApi() *string
	DoAction() function.DoActionIds
	LogGroup() *string
}

type ReceiveRequestProps interface {
	LambdaRestApi() *awsapigateway.LambdaRestApiProps
	DoAction() function.DoActionProps
	LogGroup() *awslogs.LogGroupProps
	// connections
	AddHandlerToLambdaRestApi(awslambda.Function, *awsapigateway.LambdaRestApiProps)
	AddAccessLogDestinationToLambdaRestApi(awsapigateway.LogGroupLogDestination, *awsapigateway.LambdaRestApiProps)
}

type ReceiveRequest interface {
	LambdaRestApi() awsapigateway.LambdaRestApi
}

func NewReceiveRequest(scope constructs.Construct, id ReceiveRequestIds, props ReceiveRequestProps) ReceiveRequest {
	var sprops ReceiveRequestProps = &ReceiveModelProps_DEFAULT
	var sid ReceiveRequestIds = &ReceiveModelIds_DEFAULT

	if props != nil {
		sprops = props
	}

	if id != nil {
		sid = id
	}

	fn := function.NewDoAction(scope, sid.DoAction(), sprops.DoAction())

	this := constructs.NewConstruct(scope, sid.Construct())

	logs := awslogs.NewLogGroup(this, sid.LogGroup(), sprops.LogGroup())
	logsdestinations := awsapigateway.NewLogGroupLogDestination(logs)

	sprops.AddHandlerToLambdaRestApi(fn.Function(), sprops.LambdaRestApi())
	sprops.AddAccessLogDestinationToLambdaRestApi(logsdestinations, sprops.LambdaRestApi())

	resource := awsapigateway.NewLambdaRestApi(this, sid.LambdaRestApi(), sprops.LambdaRestApi())

	var component ReceiveRequest = &ReceiveModel{
		Construct:     this,
		lambdarestapi: resource,
	}

	return component
}

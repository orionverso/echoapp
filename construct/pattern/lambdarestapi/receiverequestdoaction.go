package lambdarestapi

import (
	"castor/construct/pattern/function"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/constructs-go/constructs/v10"
)

type ReceiveRequestDoActionIds interface {
	Construct() *string
	LambdaRestApi() *string
	DoAction() function.DoActionIds
	LogGroup() *string
}

type ReceiveRequestDoActionProps interface {
	LambdaRestApi() *awsapigateway.LambdaRestApiProps
	DoAction() function.DoActionProps
	LogGroup() *awslogs.LogGroupProps
	// connections
	AddHandlerToLambdaRestApi(awslambda.Function, *awsapigateway.LambdaRestApiProps)
	AddAccessLogDestinationToLambdaRestApi(awsapigateway.LogGroupLogDestination, *awsapigateway.LambdaRestApiProps)
}

type ReceiveRequestDoAction interface {
	LambdaRestApi() awsapigateway.LambdaRestApi
	DoAction() function.DoAction
	LogGroup() awslogs.LogGroup
	LogGroupLogDestination() awsapigateway.LogGroupLogDestination
}

func NewReceiveRequestDoAction(scope constructs.Construct, id ReceiveRequestDoActionIds, props ReceiveRequestDoActionProps) ReceiveRequestDoAction {
	var sprops ReceiveRequestDoActionProps = &ReceiveRequestDoActionModelProps_DEFAULT
	var sid ReceiveRequestDoActionIds = &ReceiveRequestDoActionModelIds_DEFAULT

	if props != nil {
		sprops = props
	}

	if id != nil {
		sid = id
	}

	doaction := function.NewDoAction(scope, sid.DoAction(), sprops.DoAction())

	this := constructs.NewConstruct(scope, sid.Construct())

	loggroup := awslogs.NewLogGroup(this, sid.LogGroup(), sprops.LogGroup())
	loggrouplogdestination := awsapigateway.NewLogGroupLogDestination(loggroup)

	sprops.AddHandlerToLambdaRestApi(doaction.Function(), sprops.LambdaRestApi())
	sprops.AddAccessLogDestinationToLambdaRestApi(loggrouplogdestination, sprops.LambdaRestApi())

	resource := awsapigateway.NewLambdaRestApi(this, sid.LambdaRestApi(), sprops.LambdaRestApi())

	var component ReceiveRequestDoAction = &ReceiveRequestDoActionModel{
		Construct:              this,
		lambdarestapi:          resource,
		doaction:               doaction,
		loggroup:               loggroup,
		loggrouplogdestination: loggrouplogdestination,
	}

	return component
}

package function

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambdadestinations"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
)

type DoActionIds interface {
	Construct() *string
	Function() *string
	SuccessQueue() *string
	FailureQueue() *string
}

type DoActionProps interface {
	Function() *awslambda.FunctionProps
	SuccessQueue() *awssqs.QueueProps
	FailureQueue() *awssqs.QueueProps
	AllowReceiveMessageFromFunction() *awsiam.PolicyStatementProps
	// Connections
	AddOnSuccessDestination(awslambda.IDestination, *awslambda.FunctionProps)
	AddOnFailureDestination(awslambda.IDestination, *awslambda.FunctionProps)
	AddPrincipalToPolicy(awsiam.IPrincipal, *awsiam.PolicyStatementProps)
	AddResourceToPolicy(*string, *awsiam.PolicyStatementProps)
}

type DoAction interface {
	Function() awslambda.Function
}

func NewDoAction(scope constructs.Construct, id DoActionIds, props DoActionProps) DoAction {
	var sprops DoActionProps = &DoModelProps_DEFAULT
	var sid DoActionIds = &DoModelIds_DEFAULT

	if props != nil {
		sprops = props
	}

	if id != nil {
		sid = id
	}

	this := constructs.NewConstruct(scope, sid.Construct())

	success := awssqs.NewQueue(this, sid.SuccessQueue(), sprops.SuccessQueue())
	failure := awssqs.NewQueue(this, sid.FailureQueue(), sprops.FailureQueue())

	sprops.AddResourceToPolicy(success.QueueArn(), sprops.AllowReceiveMessageFromFunction())
	sprops.AddResourceToPolicy(failure.QueueArn(), sprops.AllowReceiveMessageFromFunction())

	destinationOnSuccess := awslambdadestinations.NewSqsDestination(success)
	destinationOnFailure := awslambdadestinations.NewSqsDestination(failure)

	sprops.AddOnSuccessDestination(destinationOnSuccess, sprops.Function())
	sprops.AddOnFailureDestination(destinationOnFailure, sprops.Function())

	fn := awslambda.NewFunction(this, sid.Function(), sprops.Function())

	sprops.AddPrincipalToPolicy(fn.Role(), sprops.AllowReceiveMessageFromFunction())

	success.GrantSendMessages(fn)
	failure.GrantSendMessages(fn)

	successpolicy := awsiam.NewPolicyStatement(sprops.AllowReceiveMessageFromFunction())
	failurepolicy := awsiam.NewPolicyStatement(sprops.AllowReceiveMessageFromFunction())

	success.AddToResourcePolicy(successpolicy)
	failure.AddToResourcePolicy(failurepolicy)

	var component DoAction = &DoModel{
		Construct: this,
		function:  fn,
	}

	return component
}

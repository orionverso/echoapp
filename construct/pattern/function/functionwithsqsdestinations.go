package function

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambdadestinations"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type FunctionWithSqsDestinationsProps struct {
	awslambda.FunctionProps
	SuccessQueueProps                awssqs.QueueProps
	FailureQueueProps                awssqs.QueueProps
	SuccessQueuePolicyStatementProps awsiam.PolicyStatementProps
	FailureQueuePolicyStatementProps awsiam.PolicyStatementProps
}

type functionWithSqsDestinations struct {
	constructs.Construct
	function           awslambda.Function
	successqueue       awssqs.Queue
	failurequeue       awssqs.Queue
	successqueuepolicy awsiam.PolicyStatement
	failurequeuepolicy awsiam.PolicyStatement
}

type FunctionWithSqsDestinations interface {
	Function() awslambda.Function
	SuccessQueue() awssqs.Queue
	FailureQueue() awssqs.Queue
	SuccessQueuePolicy() awsiam.PolicyStatement
	FailureQueuePolicy() awsiam.PolicyStatement
}

func NewFunctionWithSqsDestinations(scope constructs.Construct, id *string, props *FunctionWithSqsDestinationsProps) FunctionWithSqsDestinations {
	var sprops *FunctionWithSqsDestinationsProps = &FunctionWithSqsDestinationsProps{}

	if id == nil {
		log.Panicln("parameter id is required, but nil was provided")
	}

	if props != nil {
		sprops = props
	}

	this := constructs.NewConstruct(scope, id)

	successqueue := awssqs.NewQueue(this, jsii.String("SuccessQueue"), &sprops.SuccessQueueProps)
	failurequeue := awssqs.NewQueue(this, jsii.String("FailureQueue"), &sprops.FailureQueueProps)

	sprops.AddResourceToPolicy(successqueue.QueueArn(), &sprops.SuccessQueuePolicyStatementProps)
	sprops.AddResourceToPolicy(failurequeue.QueueArn(), &sprops.FailureQueuePolicyStatementProps)

	destinationOnSuccess := awslambdadestinations.NewSqsDestination(successqueue)
	destinationOnFailure := awslambdadestinations.NewSqsDestination(failurequeue)

	sprops.AddOnSuccessDestination(destinationOnSuccess)
	sprops.AddOnFailureDestination(destinationOnFailure)

	fn := awslambda.NewFunction(this, jsii.String("LambdaFunction"), &sprops.FunctionProps)

	sprops.AddPrincipalToPolicy(fn.Role(), &sprops.SuccessQueuePolicyStatementProps)
	sprops.AddPrincipalToPolicy(fn.Role(), &sprops.FailureQueuePolicyStatementProps)

	successqueue.GrantSendMessages(fn)
	failurequeue.GrantSendMessages(fn)

	successqueuepolicy := awsiam.NewPolicyStatement(&sprops.SuccessQueuePolicyStatementProps)
	failurequeuepolicy := awsiam.NewPolicyStatement(&sprops.FailureQueuePolicyStatementProps)

	successqueue.AddToResourcePolicy(successqueuepolicy)
	failurequeue.AddToResourcePolicy(failurequeuepolicy)

	var component FunctionWithSqsDestinations = &functionWithSqsDestinations{
		Construct:          this,
		function:           fn,
		successqueue:       successqueue,
		failurequeue:       failurequeue,
		successqueuepolicy: successqueuepolicy,
		failurequeuepolicy: failurequeuepolicy,
	}

	return component
}

// PROPS
func (props *FunctionWithSqsDestinationsProps) AddOnSuccessDestination(dst awslambda.IDestination) {
	props.FunctionProps.OnSuccess = dst
}

func (props *FunctionWithSqsDestinationsProps) AddOnFailureDestination(dst awslambda.IDestination) {
	props.FunctionProps.OnFailure = dst
}

func (props *FunctionWithSqsDestinationsProps) AddPrincipalToPolicy(principal awsiam.IPrincipal, sts *awsiam.PolicyStatementProps) {
	if sts.Principals == nil {
		sts.Principals = &[]awsiam.IPrincipal{}
	}
	*sts.Principals = append(*sts.Principals, principal)
}

func (props *FunctionWithSqsDestinationsProps) AddResourceToPolicy(resource *string, sts *awsiam.PolicyStatementProps) {
	if sts.Resources == nil {
		sts.Resources = &[]*string{}
	}
	*sts.Resources = append(*sts.Resources, resource)
}

// IMPLEMENTATION
func (mo *functionWithSqsDestinations) Function() awslambda.Function {
	return mo.function
}

func (mo *functionWithSqsDestinations) SuccessQueue() awssqs.Queue {
	return mo.successqueue
}

func (mo *functionWithSqsDestinations) FailureQueue() awssqs.Queue {
	return mo.failurequeue
}

func (mo *functionWithSqsDestinations) SuccessQueuePolicy() awsiam.PolicyStatement {
	return mo.successqueuepolicy
}

func (mo *functionWithSqsDestinations) FailureQueuePolicy() awsiam.PolicyStatement {
	return mo.failurequeuepolicy
}

// SETTINGS
// DEVELOPMENT
var FunctionWithSqsDestinationsProps_DEV FunctionWithSqsDestinationsProps = FunctionWithSqsDestinationsProps{
	FunctionProps: awslambda.FunctionProps{
		Runtime:      awslambda.Runtime_GO_1_X(),
		Handler:      jsii.String("handler"),
		Code:         awslambda.Code_FromAsset(jsii.String(fmt.Sprint(os.Getenv("ASSET_DIR"), "/echo/handler.zip")), nil),
		FunctionName: jsii.String("FunctionWithSqsDestinations-dev"),
		Description:  jsii.String("This function write the echo request from api gateway to storage solution"),
		Architecture: awslambda.Architecture_X86_64(),
		MemorySize:   jsii.Number[float64](128),
		Timeout:      awscdk.Duration_Seconds(jsii.Number[float64](2)),
	},
	SuccessQueueProps: awssqs.QueueProps{
		QueueName:     jsii.String("FunctionWithSqsDestinationsSuccessQueue"),
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	},
	FailureQueueProps: awssqs.QueueProps{
		QueueName:     jsii.String("FunctionWithSqsDestinationsFailureQueue"),
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	},

	SuccessQueuePolicyStatementProps: awsiam.PolicyStatementProps{
		Actions: jsii.Strings(
			"sqs:GetQueueAttributes",
			"sqs:GetQueueUrl",
			"sqs:SendMessage"),
		Effect: awsiam.Effect_ALLOW,
		// Principals: At runtime ,
		// Resource: At runtime
	},

	FailureQueuePolicyStatementProps: awsiam.PolicyStatementProps{
		Actions: jsii.Strings(
			"sqs:GetQueueAttributes",
			"sqs:GetQueueUrl",
			"sqs:SendMessage"),
		Effect: awsiam.Effect_ALLOW,
		// Principals: At runtime ,
		// Resource: At runtime
	},
}

// PRODUCTION
var FunctionWithSqsDestinationsProps_PROD FunctionWithSqsDestinationsProps = FunctionWithSqsDestinationsProps{
	FunctionProps: awslambda.FunctionProps{
		Runtime:      awslambda.Runtime_GO_1_X(),
		Handler:      jsii.String("handler"),
		Code:         awslambda.Code_FromAsset(jsii.String(fmt.Sprint(os.Getenv("ASSET_DIR"), "/echo/handler.zip")), nil),
		FunctionName: jsii.String("FunctionWithSqsDestinations-prod"),
		Description:  jsii.String("This function write the echo request from api gateway to storage solution"),
		Architecture: awslambda.Architecture_X86_64(),
		MemorySize:   jsii.Number[float64](512),
		Timeout:      awscdk.Duration_Seconds(jsii.Number[float64](5)),
	},
	SuccessQueueProps: awssqs.QueueProps{
		QueueName:     jsii.String("FunctionWithSqsDestinationsSuccessQueue"),
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	},
	FailureQueueProps: awssqs.QueueProps{
		QueueName:     jsii.String("FunctionWithSqsDestinationsFailureQueue"),
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	},

	SuccessQueuePolicyStatementProps: awsiam.PolicyStatementProps{
		Actions: jsii.Strings(
			"sqs:GetQueueAttributes",
			"sqs:GetQueueUrl",
			"sqs:SendMessage"),
		Effect: awsiam.Effect_ALLOW,
		// Principals: At runtime ,
		// Resource: At runtime
	},

	FailureQueuePolicyStatementProps: awsiam.PolicyStatementProps{
		Actions: jsii.Strings(
			"sqs:GetQueueAttributes",
			"sqs:GetQueueUrl",
			"sqs:SendMessage"),
		Effect: awsiam.Effect_ALLOW,
		// Principals: At runtime ,
		// Resource: At runtime
	},
}

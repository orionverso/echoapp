package function

import (
	"fmt"
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type DoModelIds struct {
	ConstructId    *string
	FunctionId     *string
	SuccessQueueId *string
	FailureQueueId *string
}

func (id DoModelIds) Construct() *string {
	return id.ConstructId
}

func (id *DoModelIds) Function() *string {
	return id.FunctionId
}

func (id *DoModelIds) SuccessQueue() *string {
	return id.SuccessQueueId
}

func (id *DoModelIds) FailureQueue() *string {
	return id.FailureQueueId
}

type DoModelProps struct {
	*awslambda.FunctionProps
	SuccessQueueProps                    *awssqs.QueueProps
	FailureQueueProps                    *awssqs.QueueProps
	AllowReceiveMessageFromFunctionProps *awsiam.PolicyStatementProps
}

func (props *DoModelProps) Function() *awslambda.FunctionProps {
	return props.FunctionProps
}

func (props *DoModelProps) SuccessQueue() *awssqs.QueueProps {
	return props.SuccessQueueProps
}

func (props *DoModelProps) FailureQueue() *awssqs.QueueProps {
	return props.FailureQueueProps
}

func (props *DoModelProps) AddOnSuccessDestination(dst awslambda.IDestination, fn *awslambda.FunctionProps) {
	fn.OnSuccess = dst
}

func (props *DoModelProps) AddOnFailureDestination(dst awslambda.IDestination, fn *awslambda.FunctionProps) {
	props.FunctionProps.OnFailure = dst
}

func (props *DoModelProps) AllowReceiveMessageFromFunction() *awsiam.PolicyStatementProps {
	return props.AllowReceiveMessageFromFunctionProps
}

func (props *DoModelProps) AddPrincipalToPolicy(principal awsiam.IPrincipal, policy *awsiam.PolicyStatementProps) {
	var principals []awsiam.IPrincipal
	principals = append(principals, principal)
	policy.Principals = &principals
}

func (props *DoModelProps) AddResourceToPolicy(resource *string, policy *awsiam.PolicyStatementProps) {
	policy.Resources = jsii.Strings(*resource)
}

type DoModel struct {
	constructs.Construct
	function           awslambda.Function
	successqueue       awssqs.Queue
	failurequeue       awssqs.Queue
	successqueuepolicy awsiam.PolicyStatement
	failurequeuepolicy awsiam.PolicyStatement
}

func (mo *DoModel) Function() awslambda.Function {
	return mo.function
}

func (mo *DoModel) SuccessQueue() awssqs.Queue {
	return mo.successqueue
}

func (mo *DoModel) FailureQueue() awssqs.Queue {
	return mo.failurequeue
}

func (mo *DoModel) SuccessQueuePolicy() awsiam.PolicyStatement {
	return mo.successqueuepolicy
}

func (mo *DoModel) FailureQueuePolicy() awsiam.PolicyStatement {
	return mo.failurequeuepolicy
}

// SETTINGS
var DoModelIds_DEFAULT DoModelIds = DoModelIds{
	ConstructId:    jsii.String("MODEL-resource-domodel-default"),
	FunctionId:     jsii.String("MODEL-resource-function-default"),
	SuccessQueueId: jsii.String("MODEL-resource-successqueue-default"),
	FailureQueueId: jsii.String("MODEL-resource-failurequeue-default"),
}

var DoModelProps_DEFAULT DoModelProps = DoModelProps{
	FunctionProps: &awslambda.FunctionProps{
		Runtime:      awslambda.Runtime_GO_1_X(),
		Handler:      jsii.String("handler"),
		Code:         awslambda.Code_FromAsset(jsii.String(fmt.Sprint(os.Getenv("ASSET_DIR"), "/echo/handler.zip")), nil),
		FunctionName: jsii.String("DoModel"),
		Description:  jsii.String("This function is a model to create any function. It's extended with sqs queues destinations on failure and success. It have minimal configuration to can works"),
	},
	SuccessQueueProps: &awssqs.QueueProps{
		QueueName:     jsii.String("DoModelSuccessQueue"),
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	},
	FailureQueueProps: &awssqs.QueueProps{
		QueueName:     jsii.String("DoModelFailureQueue"),
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
}

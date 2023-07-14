package computesave

import (
	"castor/construct/pattern/lambdarestapi"
	"castor/stack/environment"
	"log"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type ApiWriteToTableProps struct {
	awscdk.StackProps
	lambdarestapi.ApiGatewayWithLambdaProxyIntegratedProps
	awsdynamodb.TableProps
}

type apiWriteToTable struct {
	awscdk.Stack
	writer  lambdarestapi.ApiGatewayWithLambdaProxyIntegrated
	storage awsdynamodb.Table
}

type ApiWriteToTable interface {
	awscdk.Stack
	ApiProxyIntegrated() lambdarestapi.ApiGatewayWithLambdaProxyIntegrated
	Table() awsdynamodb.Table
}

func NewApiWriteToTable(scope constructs.Construct, id *string, props *ApiWriteToTableProps) ApiWriteToTable {
	var sprops *ApiWriteToTableProps = &ApiWriteToTableProps{}

	if id == nil {
		log.Panicln("parameter id is required, but nil was provided")
	}

	if props != nil {
		sprops = props
	}

	stack := awscdk.NewStack(scope, id, &sprops.StackProps)

	api := lambdarestapi.NewApiGatewayWithLambdaProxyIntegrated(stack, jsii.String("Writer"), &sprops.ApiGatewayWithLambdaProxyIntegratedProps)

	tb := awsdynamodb.NewTable(stack, jsii.String("Storage"), &sprops.TableProps)

	fn := api.FunctionWithSqsDestinations().Function()

	tb.GrantWriteData(fn.Role())

	fn.AddEnvironment(
		jsii.String("STORAGE_SERVICE"),
		jsii.String("DYNAMODB"),
		&awslambda.EnvironmentOptions{})

	fn.AddEnvironment(
		jsii.String("DESTINATION"),
		tb.TableName(),
		&awslambda.EnvironmentOptions{})

	var component ApiWriteToTable = &apiWriteToTable{
		Stack:   stack,
		writer:  api,
		storage: tb,
	}

	return component
}

// IMPLEMENTATION
func (mo *apiWriteToTable) ApiProxyIntegrated() lambdarestapi.ApiGatewayWithLambdaProxyIntegrated {
	return mo.writer
}

func (mo *apiWriteToTable) Table() awsdynamodb.Table {
	return mo.storage
}

// SETTINGS
// DEVELOPMENT
var ApiWriteToTableProps_DEV ApiWriteToTableProps = ApiWriteToTableProps{
	StackProps:                               environment.StackProps_DEV,
	ApiGatewayWithLambdaProxyIntegratedProps: lambdarestapi.ApiGatewayWithLambdaProxyIntegratedProps_DEV,
	TableProps: awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("id"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	},
}

// PRODUCTION
var ApiWriteToTableProps_PROD ApiWriteToTableProps = ApiWriteToTableProps{
	StackProps:                               environment.StackProps_PROD,
	ApiGatewayWithLambdaProxyIntegratedProps: lambdarestapi.ApiGatewayWithLambdaProxyIntegratedProps_PROD,
	TableProps: awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("id"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	},
}

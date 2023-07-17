package computesave

import (
	"castor/construct/highlevel/table"
	"castor/construct/pattern/lambdarestapi"
	"castor/stack/environment"
	"log"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type ApiWriteToSaveBlockDataProps struct {
	awscdk.StackProps
	lambdarestapi.ApiGatewayWithLambdaProxyIntegratedProps
	table.SaveBlockDataProps
}

type apiWriteToSaveBlockData struct {
	awscdk.Stack
	writer  lambdarestapi.ApiGatewayWithLambdaProxyIntegrated
	storage table.SaveBlockData
}

type ApiWriteToSaveBlockData interface {
	awscdk.Stack
	ApiProxyIntegrated() lambdarestapi.ApiGatewayWithLambdaProxyIntegrated
	SaveBlockData() table.SaveBlockData
}

func NewApiWriteToSaveBlockData(scope constructs.Construct, id *string, props *ApiWriteToSaveBlockDataProps) ApiWriteToSaveBlockData {
	var sprops *ApiWriteToSaveBlockDataProps = &ApiWriteToSaveBlockDataProps{}

	if id == nil {
		log.Panicln("parameter id is required, but nil was provided")
	}

	if props != nil {
		sprops = props
	}

	this := constructs.NewConstruct(scope, id)

	awscdk.NewStack(this, stateful, &sprops.StackProps) // empty stack //Delete fargate stateful

	stackless := awscdk.NewStack(this, stateless, &sprops.StackProps)

	api := lambdarestapi.NewApiGatewayWithLambdaProxyIntegrated(stackless, serverless, &sprops.ApiGatewayWithLambdaProxyIntegratedProps)

	sv := table.NewSaveBlockData(stackless, database, &sprops.SaveBlockDataProps)

	tb := sv.Table()

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

	var component ApiWriteToSaveBlockData = &apiWriteToSaveBlockData{
		Stack:   stackless,
		writer:  api,
		storage: sv,
	}

	return component
}

// IMPLEMENTATION
func (mo *apiWriteToSaveBlockData) ApiProxyIntegrated() lambdarestapi.ApiGatewayWithLambdaProxyIntegrated {
	return mo.writer
}

func (mo *apiWriteToSaveBlockData) SaveBlockData() table.SaveBlockData {
	return mo.storage
}

// SETTINGS
// DEVELOPMENT
var ApiWriteToSaveBlockDataProps_DEV ApiWriteToSaveBlockDataProps = ApiWriteToSaveBlockDataProps{
	StackProps:                               environment.StackProps_DEV,
	ApiGatewayWithLambdaProxyIntegratedProps: lambdarestapi.ApiGatewayWithLambdaProxyIntegratedProps_DEV,
	SaveBlockDataProps:                       table.SaveBlockDataProps_DEV,
}

// PRODUCTION
var ApiWriteToSaveBlockDataProps_PROD ApiWriteToSaveBlockDataProps = ApiWriteToSaveBlockDataProps{
	StackProps:                               environment.StackProps_PROD,
	ApiGatewayWithLambdaProxyIntegratedProps: lambdarestapi.ApiGatewayWithLambdaProxyIntegratedProps_PROD,
	SaveBlockDataProps:                       table.SaveBlockDataProps_PROD,
}

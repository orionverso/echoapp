package computesave

import (
	"castor/construct/highlevel/bucket"
	"castor/construct/pattern/lambdarestapi"
	"castor/stack/environment"
	"log"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type ApiWriteToSaveObjectProps struct {
	awscdk.StackProps
	lambdarestapi.ApiGatewayWithLambdaProxyIntegratedProps
	bucket.SaveObjectProps
}

type apiWriteToSaveObject struct {
	awscdk.Stack
	writer  lambdarestapi.ApiGatewayWithLambdaProxyIntegrated
	storage bucket.SaveObject
}

type ApiWriteToSaveObject interface {
	awscdk.Stack
	ApiProxyIntegrated() lambdarestapi.ApiGatewayWithLambdaProxyIntegrated
	SaveObject() bucket.SaveObject
}

func NewApiWriteToSaveObject(scope constructs.Construct, id *string, props *ApiWriteToSaveObjectProps) ApiWriteToSaveObject {
	var sprops *ApiWriteToSaveObjectProps = &ApiWriteToSaveObjectProps{}

	if id == nil {
		log.Panicln("parameter id is required, but nil was provided")
	}

	if props != nil {
		sprops = props
	}

	stack := awscdk.NewStack(scope, id, &sprops.StackProps)

	api := lambdarestapi.NewApiGatewayWithLambdaProxyIntegrated(stack, jsii.String("Writer"), &sprops.ApiGatewayWithLambdaProxyIntegratedProps)

	sv := bucket.NewSaveObject(stack, jsii.String("Storage"), &sprops.SaveObjectProps)

	fn := api.FunctionWithSqsDestinations().Function()

	bk := sv.Bucket()

	bk.GrantWrite(fn.Role(), jsii.String("*"), jsii.Strings("*"))

	fn.AddEnvironment(
		jsii.String("STORAGE_SERVICE"),
		jsii.String("S3"),
		&awslambda.EnvironmentOptions{})

	fn.AddEnvironment(
		jsii.String("DESTINATION"),
		bk.BucketName(),
		&awslambda.EnvironmentOptions{})

	var component ApiWriteToSaveObject = &apiWriteToSaveObject{
		Stack:   stack,
		writer:  api,
		storage: sv,
	}

	return component
}

// IMPLEMENTATION
func (mo *apiWriteToSaveObject) ApiProxyIntegrated() lambdarestapi.ApiGatewayWithLambdaProxyIntegrated {
	return mo.writer
}

func (mo *apiWriteToSaveObject) SaveObject() bucket.SaveObject {
	return mo.storage
}

// SETTINGS
// DEVELOPMENT
var ApiWriteToSaveObjectProps_DEV ApiWriteToSaveObjectProps = ApiWriteToSaveObjectProps{
	StackProps:                               environment.StackProps_DEV,
	ApiGatewayWithLambdaProxyIntegratedProps: lambdarestapi.ApiGatewayWithLambdaProxyIntegratedProps_DEV,
	SaveObjectProps:                          bucket.SaveObjectProps_DEV,
}

// PRODUCTION
var ApiWriteToSaveObjectProps_PROD ApiWriteToSaveObjectProps = ApiWriteToSaveObjectProps{
	StackProps:                               environment.StackProps_PROD,
	ApiGatewayWithLambdaProxyIntegratedProps: lambdarestapi.ApiGatewayWithLambdaProxyIntegratedProps_PROD,
	SaveObjectProps:                          bucket.SaveObjectProps_PROD,
}

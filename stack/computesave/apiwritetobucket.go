package computesave

import (
	"castor/construct/highlevel/bucket"
	"castor/construct/pattern/lambdarestapi"
	"castor/stack/environment"
	"log"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type ApiWriteToBucketProps struct {
	awscdk.StackProps
	lambdarestapi.ApiGatewayWithLambdaProxyIntegratedProps
	bucket.SaveFileProps
}

type apiWriteToBucket struct {
	awscdk.Stack
	writer  lambdarestapi.ApiGatewayWithLambdaProxyIntegrated
	storage bucket.SaveFile
}

type ApiWriteToBucket interface {
	awscdk.Stack
	ApiProxyIntegrated() lambdarestapi.ApiGatewayWithLambdaProxyIntegrated
	Bucket() awss3.Bucket
}

func NewApiWriteToBucket(scope constructs.Construct, id *string, props *ApiWriteToBucketProps) ApiWriteToBucket {
	var sprops *ApiWriteToBucketProps = &ApiWriteToBucketProps{}

	if id == nil {
		log.Panicln("parameter id is required, but nil was provided")
	}

	if props != nil {
		sprops = props
	}

	stack := awscdk.NewStack(scope, id, &sprops.StackProps)

	api := lambdarestapi.NewApiGatewayWithLambdaProxyIntegrated(stack, jsii.String("Writer"), &sprops.ApiGatewayWithLambdaProxyIntegratedProps)

	sv := bucket.NewSaveFile(stack, jsii.String("Storage"), &sprops.SaveFileProps)

	fn := api.FunctionWithSqsDestinations().Function()

	bk := sv.Bucket()

	bk.Bucket().GrantWrite(fn.Role(), jsii.String("*"), jsii.Strings("*"))

	fn.AddEnvironment(
		jsii.String("STORAGE_SERVICE"),
		jsii.String("S3"),
		&awslambda.EnvironmentOptions{})

	fn.AddEnvironment(
		jsii.String("DESTINATION"),
		bk.BucketName(),
		&awslambda.EnvironmentOptions{})

	var component ApiWriteToBucket = &apiWriteToBucket{
		Stack:   stack,
		writer:  api,
		storage: sv,
	}

	return component
}

// IMPLEMENTATION
func (mo *apiWriteToBucket) ApiProxyIntegrated() lambdarestapi.ApiGatewayWithLambdaProxyIntegrated {
	return mo.writer
}

func (mo *apiWriteToBucket) Bucket() awss3.Bucket {
	return mo.storage
}

// SETTINGS
// DEVELOPMENT
var ApiWriteToBucketProps_DEV ApiWriteToBucketProps = ApiWriteToBucketProps{
	StackProps:                               environment.StackProps_DEV,
	ApiGatewayWithLambdaProxyIntegratedProps: lambdarestapi.ApiGatewayWithLambdaProxyIntegratedProps_DEV,
	BucketProps: awss3.BucketProps{
		AutoDeleteObjects: jsii.Bool(true),
		RemovalPolicy:     awscdk.RemovalPolicy_DESTROY,
	},
}

// PRODUCTION
var ApiWriteToBucketProps_PROD ApiWriteToBucketProps = ApiWriteToBucketProps{
	StackProps:                               environment.StackProps_PROD,
	ApiGatewayWithLambdaProxyIntegratedProps: lambdarestapi.ApiGatewayWithLambdaProxyIntegratedProps_PROD,
	BucketProps: awss3.BucketProps{
		AutoDeleteObjects: jsii.Bool(true),
		RemovalPolicy:     awscdk.RemovalPolicy_DESTROY,
	},
}

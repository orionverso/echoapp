package computesave

import (
	fargate "castor/construct/pattern/applicationloadbalancedfargateservice"
	"castor/stack/environment"
	"log"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type FargateWriteToBucketProps struct {
	awscdk.StackProps
	fargate.FargateEcrImageProps
	awss3.BucketProps
}

type fargateWriteToBucket struct {
	awscdk.Stack
	writer  fargate.FargateEcrImage
	storage awss3.Bucket
}

type FargateWriteToBucket interface {
	awscdk.Stack
	Fargate() fargate.FargateEcrImage
	Bucket() awss3.Bucket
}

func NewFargateWriteToBucket(scope constructs.Construct, id *string, props *FargateWriteToBucketProps) FargateWriteToBucket {
	var sprops *FargateWriteToBucketProps = &FargateWriteToBucketProps{}

	if id == nil {
		log.Panicln("parameter id is required, but nil was provided")
	}

	if props != nil {
		sprops = props
	}

	stack := awscdk.NewStack(scope, id, &sprops.StackProps)

	fg := fargate.NewFargateEcrImage(stack, jsii.String("Writer"), &sprops.FargateEcrImageProps)

	fgrole := fg.ApplicationLoadBalancedFargateService().TaskDefinition().TaskRole()

	bk := awss3.NewBucket(stack, jsii.String("Storage"), &sprops.BucketProps)

	bk.GrantWrite(fgrole, jsii.String("*"), jsii.Strings("*"))

	defaultcontainer := fg.ApplicationLoadBalancedFargateService().TaskDefinition().DefaultContainer()
	defaultcontainer.AddEnvironment(jsii.String("STORAGE_SERVICE"), jsii.String("S3"))
	defaultcontainer.AddEnvironment(jsii.String("DESTINATION"), bk.BucketName())

	var component FargateWriteToBucket = &fargateWriteToBucket{
		Stack:   stack,
		writer:  fg,
		storage: bk,
	}

	return component
}

// IMPLEMENTATION
func (mo fargateWriteToBucket) Fargate() fargate.FargateEcrImage {
	return mo.writer
}

func (mo fargateWriteToBucket) Bucket() awss3.Bucket {
	return mo.storage
}

// SETTINGS
// DEVELOPMENT
var FargateWriteToBucketProps_DEV FargateWriteToBucketProps = FargateWriteToBucketProps{
	StackProps:           environment.StackProps_DEV,
	FargateEcrImageProps: fargate.FargateEcrImageProps_DEV,
	BucketProps: awss3.BucketProps{
		AutoDeleteObjects: jsii.Bool(true),
		RemovalPolicy:     awscdk.RemovalPolicy_DESTROY,
	},
}

// PRODUCTION
var FargateWriteToBucketProps_PROD FargateWriteToBucketProps = FargateWriteToBucketProps{
	StackProps:           environment.StackProps_PROD,
	FargateEcrImageProps: fargate.FargateEcrImageProps_PROD,
	BucketProps: awss3.BucketProps{
		AutoDeleteObjects: jsii.Bool(true),
		RemovalPolicy:     awscdk.RemovalPolicy_DESTROY,
	},
}

package computesave

import (
	"castor/construct/highlevel/bucket"
	fargate "castor/construct/pattern/applicationloadbalancedfargateservice"
	"castor/stack/environment"
	"log"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type FargateWriteToSaveObjectProps struct {
	awscdk.StackProps
	fargate.FargateEcrImageProps
	bucket.SaveObjectProps
}

type fargateWriteToSaveObject struct {
	awscdk.Stack
	writer  fargate.FargateEcrImage
	storage bucket.SaveObject
}

type FargateWriteToSaveObject interface {
	awscdk.Stack
	Fargate() fargate.FargateEcrImage
	SaveObject() bucket.SaveObject
}

func NewFargateWriteToSaveObject(scope constructs.Construct, id *string, props *FargateWriteToSaveObjectProps) FargateWriteToSaveObject {
	var sprops *FargateWriteToSaveObjectProps = &FargateWriteToSaveObjectProps{}

	if id == nil {
		log.Panicln("parameter id is required, but nil was provided")
	}

	if props != nil {
		sprops = props
	}

	stack := awscdk.NewStack(scope, id, &sprops.StackProps)

	fg := fargate.NewFargateEcrImage(stack, jsii.String("Writer"), &sprops.FargateEcrImageProps)

	fgrole := fg.ApplicationLoadBalancedFargateService().TaskDefinition().TaskRole()

	sv := bucket.NewSaveObject(stack, jsii.String("Storage"), &sprops.SaveObjectProps)

	bk := sv.Bucket()

	bk.GrantWrite(fgrole, jsii.String("*"), jsii.Strings("*"))

	defaultcontainer := fg.ApplicationLoadBalancedFargateService().TaskDefinition().DefaultContainer()
	defaultcontainer.AddEnvironment(jsii.String("STORAGE_SERVICE"), jsii.String("S3"))
	defaultcontainer.AddEnvironment(jsii.String("DESTINATION"), bk.BucketName())

	var component FargateWriteToSaveObject = &fargateWriteToSaveObject{
		Stack:   stack,
		writer:  fg,
		storage: sv,
	}

	return component
}

// IMPLEMENTATION
func (mo fargateWriteToSaveObject) Fargate() fargate.FargateEcrImage {
	return mo.writer
}

func (mo fargateWriteToSaveObject) SaveObject() bucket.SaveObject {
	return mo.storage
}

// SETTINGS
// DEVELOPMENT
var FargateWriteToSaveObjectProps_DEV FargateWriteToSaveObjectProps = FargateWriteToSaveObjectProps{
	StackProps:           environment.StackProps_DEV,
	FargateEcrImageProps: fargate.FargateEcrImageProps_DEV,
	SaveObjectProps:      bucket.SaveObjectProps_DEV,
}

// PRODUCTION
var FargateWriteToSaveObjectProps_PROD FargateWriteToSaveObjectProps = FargateWriteToSaveObjectProps{
	StackProps:           environment.StackProps_PROD,
	FargateEcrImageProps: fargate.FargateEcrImageProps_PROD,
	SaveObjectProps:      bucket.SaveObjectProps_PROD,
}

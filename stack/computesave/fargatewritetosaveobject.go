package computesave

import (
	"castor/construct/highlevel/bucket"
	"castor/construct/highlevel/repository"
	fargate "castor/construct/pattern/applicationloadbalancedfargateservice"
	"castor/stack/environment"
	"log"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type FargateWriteToSaveObjectProps struct {
	awscdk.StackProps
	fargate.FargateProps
	bucket.SaveObjectProps
	repository.EcrRepoProps
}

type fargateWriteToSaveObject struct {
	awscdk.Stack
	writer     fargate.Fargate
	storage    bucket.SaveObject
	repository repository.EcrRepo
}

type FargateWriteToSaveObject interface {
	awscdk.Stack
	Fargate() fargate.Fargate
	SaveObject() bucket.SaveObject
	EcrRepo() repository.EcrRepo
}

func NewFargateWriteToSaveObject(scope constructs.Construct, id *string, props *FargateWriteToSaveObjectProps) FargateWriteToSaveObject {
	var sprops *FargateWriteToSaveObjectProps = &FargateWriteToSaveObjectProps{}

	if id == nil {
		log.Panicln("parameter id is required, but nil was provided")
	}

	if props != nil {
		sprops = props
	}

	this := constructs.NewConstruct(scope, id)

	stackful := awscdk.NewStack(this, stateful, &sprops.StackProps)

	repo := repository.NewEcrRepo(stackful, fargaterepo, &sprops.EcrRepoProps)

	stackless := awscdk.NewStack(this, stateless, &sprops.StackProps)

	sprops.AddContainerImageToApplicationLoadBalancedFargate(repo.Image())

	fg := fargate.NewFargate(stackless, server, &sprops.FargateProps)

	fgrole := fg.ApplicationLoadBalancedFargateService().TaskDefinition().TaskRole()

	sv := bucket.NewSaveObject(stackless, object, &sprops.SaveObjectProps)

	bk := sv.Bucket()

	bk.GrantWrite(fgrole, jsii.String("*"), jsii.Strings("*"))

	defaultcontainer := fg.ApplicationLoadBalancedFargateService().TaskDefinition().DefaultContainer()
	defaultcontainer.AddEnvironment(jsii.String("STORAGE_SERVICE"), jsii.String("S3"))
	defaultcontainer.AddEnvironment(jsii.String("DESTINATION"), bk.BucketName())

	var component FargateWriteToSaveObject = &fargateWriteToSaveObject{
		Stack:      stackless,
		writer:     fg,
		storage:    sv,
		repository: repo,
	}

	return component
}

// IMPLEMENTATION
func (mo *fargateWriteToSaveObject) Fargate() fargate.Fargate {
	return mo.writer
}

func (mo *fargateWriteToSaveObject) SaveObject() bucket.SaveObject {
	return mo.storage
}

func (mo *fargateWriteToSaveObject) EcrRepo() repository.EcrRepo {
	return mo.repository
}

// SETTINGS
// DEVELOPMENT
var FargateWriteToSaveObjectProps_DEV FargateWriteToSaveObjectProps = FargateWriteToSaveObjectProps{
	StackProps:      environment.StackProps_DEV,
	FargateProps:    fargate.FargateProps_DEV,
	SaveObjectProps: bucket.SaveObjectProps_DEV,
	EcrRepoProps:    repository.EcrRepoProps_DEV,
}

// PRODUCTION
var FargateWriteToSaveObjectProps_PROD FargateWriteToSaveObjectProps = FargateWriteToSaveObjectProps{
	StackProps:      environment.StackProps_PROD,
	FargateProps:    fargate.FargateProps_PROD,
	SaveObjectProps: bucket.SaveObjectProps_PROD,
	EcrRepoProps:    repository.EcrRepoProps_PROD,
}

package computesave

import (
	"castor/construct/highlevel/repository"
	"castor/construct/highlevel/table"
	fargate "castor/construct/pattern/applicationloadbalancedfargateservice"
	"castor/stack/environment"
	"log"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type FargateWriteToSaveBlockDataProps struct {
	awscdk.StackProps
	fargate.FargateProps
	table.SaveBlockDataProps
	repository.EcrRepoProps
}

type fargateWriteToSaveBlockData struct {
	awscdk.Stack
	writer     fargate.Fargate
	storage    table.SaveBlockData
	repository repository.EcrRepo
}

type FargateWriteToSaveBlockData interface {
	awscdk.Stack
	Fargate() fargate.Fargate
	SaveBlockData() table.SaveBlockData
	EcrRepo() repository.EcrRepo
}

func NewFargateWriteToSaveBlockData(scope constructs.Construct, id *string, props *FargateWriteToSaveBlockDataProps) FargateWriteToSaveBlockData {
	var sprops *FargateWriteToSaveBlockDataProps = &FargateWriteToSaveBlockDataProps{}

	if id == nil {
		log.Panicln("parameter id is required, but nil was provided")
	}

	if props != nil {
		sprops = props
	}

	this := constructs.NewConstruct(scope, id)

	stackful := awscdk.NewStack(this, jsii.String("stateful"), &sprops.StackProps)

	repo := repository.NewEcrRepo(stackful, jsii.String("FargateRepository"), &sprops.EcrRepoProps)

	stackless := awscdk.NewStack(this, jsii.String("stateless"), &sprops.StackProps)

	sprops.AddContainerImageToApplicationLoadBalancedFargate(repo.Image())

	sv := table.NewSaveBlockData(stackless, jsii.String("Storage"), &sprops.SaveBlockDataProps)

	fg := fargate.NewFargate(stackless, jsii.String("Writer"), &sprops.FargateProps)

	fgrole := fg.ApplicationLoadBalancedFargateService().TaskDefinition().TaskRole()

	tb := sv.Table()

	tb.GrantWriteData(fgrole)

	defaultcontainer := fg.ApplicationLoadBalancedFargateService().TaskDefinition().DefaultContainer()
	defaultcontainer.AddEnvironment(jsii.String("STORAGE_SERVICE"), jsii.String("DYNAMODB"))
	defaultcontainer.AddEnvironment(jsii.String("DESTINATION"), tb.TableName())

	var component FargateWriteToSaveBlockData = &fargateWriteToSaveBlockData{
		Stack:   stackless,
		writer:  fg,
		storage: sv,
	}

	return component
}

// IMPLEMENTATION
func (mo *fargateWriteToSaveBlockData) Fargate() fargate.Fargate {
	return mo.writer
}

func (mo *fargateWriteToSaveBlockData) SaveBlockData() table.SaveBlockData {
	return mo.storage
}

func (mo *fargateWriteToSaveBlockData) EcrRepo() repository.EcrRepo {
	return mo.repository
}

// SETTINGS
// DEVELOPMENT
var FargateWriteToSaveBlockDataProps_DEV FargateWriteToSaveBlockDataProps = FargateWriteToSaveBlockDataProps{
	StackProps:         environment.StackProps_DEV,
	FargateProps:       fargate.FargateProps_DEV,
	SaveBlockDataProps: table.SaveBlockDataProps_DEV,
	EcrRepoProps:       repository.EcrRepoProps_DEV,
}

// PRODUCTION
var FargateWriteToSaveBlockDataProps_PROD FargateWriteToSaveBlockDataProps = FargateWriteToSaveBlockDataProps{
	StackProps:         environment.StackProps_PROD,
	FargateProps:       fargate.FargateProps_PROD,
	SaveBlockDataProps: table.SaveBlockDataProps_PROD,
	EcrRepoProps:       repository.EcrRepoProps_PROD,
}

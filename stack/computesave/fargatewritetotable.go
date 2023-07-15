package computesave

import (
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
	fargate.FargateEcrImageProps
	table.SaveBlockDataProps
}

type fargateWriteToTable struct {
	awscdk.Stack
	writer  fargate.FargateEcrImage
	storage table.SaveBlockData
}

type FargateWriteToSaveBlockData interface {
	awscdk.Stack
	Fargate() fargate.FargateEcrImage
	SaveBlockData() table.SaveBlockData
}

func NewFargateWriteToSaveBlockData(scope constructs.Construct, id *string, props *FargateWriteToSaveBlockDataProps) FargateWriteToSaveBlockData {
	var sprops *FargateWriteToSaveBlockDataProps = &FargateWriteToSaveBlockDataProps{}

	if id == nil {
		log.Panicln("parameter id is required, but nil was provided")
	}

	if props != nil {
		sprops = props
	}

	stack := awscdk.NewStack(scope, id, &sprops.StackProps)

	fg := fargate.NewFargateEcrImage(stack, jsii.String("Writer"), &sprops.FargateEcrImageProps)

	fgrole := fg.ApplicationLoadBalancedFargateService().TaskDefinition().TaskRole()

	sv := table.NewSaveBlockData(stack, jsii.String("Storage"), &sprops.SaveBlockDataProps)

	tb := sv.Table()

	tb.GrantWriteData(fgrole)

	defaultcontainer := fg.ApplicationLoadBalancedFargateService().TaskDefinition().DefaultContainer()
	defaultcontainer.AddEnvironment(jsii.String("STORAGE_SERVICE"), jsii.String("DYNAMODB"))
	defaultcontainer.AddEnvironment(jsii.String("DESTINATION"), tb.TableName())

	var component FargateWriteToSaveBlockData = &fargateWriteToTable{
		Stack:   stack,
		writer:  fg,
		storage: sv,
	}

	return component
}

// IMPLEMENTATION
func (mo fargateWriteToTable) Fargate() fargate.FargateEcrImage {
	return mo.writer
}

func (mo fargateWriteToTable) SaveBlockData() table.SaveBlockData {
	return mo.storage
}

// SETTINGS
// DEVELOPMENT
var FargateWriteToSaveBlockDataProps_DEV FargateWriteToSaveBlockDataProps = FargateWriteToSaveBlockDataProps{
	StackProps:           environment.StackProps_DEV,
	FargateEcrImageProps: fargate.FargateEcrImageProps_DEV,
	SaveBlockDataProps:   table.SaveBlockDataProps_DEV,
}

// PRODUCTION
var FargateWriteToSaveBlockDataProps_PROD FargateWriteToSaveBlockDataProps = FargateWriteToSaveBlockDataProps{
	StackProps:           environment.StackProps_PROD,
	FargateEcrImageProps: fargate.FargateEcrImageProps_PROD,
	SaveBlockDataProps:   table.SaveBlockDataProps_PROD,
}

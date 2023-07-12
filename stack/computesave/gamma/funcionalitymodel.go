package delta

import (
	fargate "castor/construct/pattern/applicationloadbalancedfargateservice"
	"castor/construct/pattern/table"
	"castor/stack/environment"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

type FuncionalityModelIds struct {
	StackId *string
	fargate.ReceiveRequestDoActionIds
	table.SaveBlockDataIds
}

func (id *FuncionalityModelIds) Stack() *string {
	return id.StackId
}

func (id *FuncionalityModelIds) ReceiveRequestDoAction() fargate.ReceiveRequestDoActionIds {
	return id.ReceiveRequestDoActionIds
}

func (id *FuncionalityModelIds) SaveBlockData() table.SaveBlockDataIds {
	return id.SaveBlockDataIds
}

type FuncionalityModelProps struct {
	awscdk.StackProps
	fargate.ReceiveRequestDoActionProps
	table.SaveBlockDataProps
}

func (props *FuncionalityModelProps) Stack() *awscdk.StackProps {
	return &props.StackProps
}

func (props *FuncionalityModelProps) ReceiveRequestDoAction() fargate.ReceiveRequestDoActionProps {
	return props.ReceiveRequestDoActionProps
}

func (props *FuncionalityModelProps) SaveBlockData() table.SaveBlockDataProps {
	return props.SaveBlockDataProps
}

type FuncionalityModel struct {
	awscdk.Stack
	receiverequestdoaction fargate.ReceiveRequestDoAction
	savefile               table.SaveBlockData
}

func (mo FuncionalityModel) ReceiveRequestDoAction() fargate.ReceiveRequestDoAction {
	return mo.receiverequestdoaction
}

func (mo FuncionalityModel) SaveBlockData() table.SaveBlockData {
	return mo.savefile
}

// SETTINGS
var FuncionalityModelIds_DEFAULT FuncionalityModelIds = FuncionalityModelIds{
	StackId:                   jsii.String("FuncionalityModel-default"),
	ReceiveRequestDoActionIds: &fargate.ReceiveRequestDoActionModelIds_DEFAULT,
	SaveBlockDataIds:          &table.SaveBlockModelIds_DEFAULT,
}

var FuncionalityModelProps_DEFAULT FuncionalityModelProps = FuncionalityModelProps{
	StackProps:                  environment.StackProps_DEFAULT,
	ReceiveRequestDoActionProps: &fargate.ReceiveRequestDoActionModelProps_DEFAULT,
	SaveBlockDataProps:          &table.SaveBlockModelProps_DEFAULT,
}

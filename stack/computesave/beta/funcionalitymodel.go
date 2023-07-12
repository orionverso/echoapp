package alfa

import (
	"castor/construct/pattern/lambdarestapi"
	"castor/construct/pattern/table"
	"castor/stack/environment"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

type FuncionalityModelIds struct {
	StackId *string
	lambdarestapi.ReceiveRequestDoActionIds
	table.SaveBlockDataIds
}

func (id *FuncionalityModelIds) Stack() *string {
	return id.StackId
}

func (id *FuncionalityModelIds) ReceiveRequestDoAction() lambdarestapi.ReceiveRequestDoActionIds {
	return id.ReceiveRequestDoActionIds
}

func (id *FuncionalityModelIds) SaveBlockData() table.SaveBlockDataIds {
	return id.SaveBlockDataIds
}

type FuncionalityModelProps struct {
	awscdk.StackProps
	lambdarestapi.ReceiveRequestDoActionProps
	table.SaveBlockDataProps
}

func (props *FuncionalityModelProps) Stack() *awscdk.StackProps {
	return &props.StackProps
}

func (props *FuncionalityModelProps) ReceiveRequestDoAction() lambdarestapi.ReceiveRequestDoActionProps {
	return props.ReceiveRequestDoActionProps
}

func (props *FuncionalityModelProps) SaveBlockData() table.SaveBlockDataProps {
	return props.SaveBlockDataProps
}

type FuncionalityModel struct {
	awscdk.Stack
	receiverequestdoaction lambdarestapi.ReceiveRequestDoAction
	saveblockdata          table.SaveBlockData
}

func (mo FuncionalityModel) ReceiveRequestDoAction() lambdarestapi.ReceiveRequestDoAction {
	return mo.receiverequestdoaction
}

func (mo FuncionalityModel) SaveBlockData() table.SaveBlockData {
	return mo.saveblockdata
}

// SETTINGS
var FuncionalityModelIds_DEFAULT FuncionalityModelIds = FuncionalityModelIds{
	StackId:                   jsii.String("FuncionalityModel-default"),
	ReceiveRequestDoActionIds: &lambdarestapi.ReceiveRequestDoActionModelIds_DEFAULT,
	SaveBlockDataIds:          &table.SaveBlockModelIds_DEFAULT,
}

var FuncionalityModelProps_DEFAULT FuncionalityModelProps = FuncionalityModelProps{
	StackProps:                  environment.StackProps_DEFAULT,
	ReceiveRequestDoActionProps: &lambdarestapi.ReceiveRequestDoActionModelProps_DEFAULT,
	SaveBlockDataProps:          &table.SaveBlockModelProps_DEFAULT,
}

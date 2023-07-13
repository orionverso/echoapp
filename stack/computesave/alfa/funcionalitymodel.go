package alfa

import (
	"castor/construct/highlevel/bucket"
	"castor/construct/pattern/lambdarestapi"
	"castor/stack/environment"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

type FuncionalityModelIds struct {
	StackId *string
	lambdarestapi.ReceiveRequestDoActionIds
	bucket.SaveFileIds
}

func (id *FuncionalityModelIds) Stack() *string {
	return id.StackId
}

func (id *FuncionalityModelIds) ReceiveRequestDoAction() lambdarestapi.ReceiveRequestDoActionIds {
	return id.ReceiveRequestDoActionIds
}

func (id *FuncionalityModelIds) SaveFile() bucket.SaveFileIds {
	return id.SaveFileIds
}

type FuncionalityModelProps struct {
	awscdk.StackProps
	lambdarestapi.ReceiveRequestDoActionProps
	bucket.SaveFileProps
}

func (props *FuncionalityModelProps) Stack() *awscdk.StackProps {
	return &props.StackProps
}

func (props *FuncionalityModelProps) ReceiveRequestDoAction() lambdarestapi.ReceiveRequestDoActionProps {
	return props.ReceiveRequestDoActionProps
}

func (props *FuncionalityModelProps) SaveFile() bucket.SaveFileProps {
	return props.SaveFileProps
}

type FuncionalityModel struct {
	awscdk.Stack
	receiverequestdoaction lambdarestapi.ReceiveRequestDoAction
	savefile               bucket.SaveFile
}

func (mo FuncionalityModel) ReceiveRequestDoAction() lambdarestapi.ReceiveRequestDoAction {
	return mo.receiverequestdoaction
}

func (mo FuncionalityModel) SaveFile() bucket.SaveFile {
	return mo.savefile
}

// SETTINGS
var FuncionalityModelIds_DEFAULT FuncionalityModelIds = FuncionalityModelIds{
	StackId:                   jsii.String("FuncionalityModel-default"),
	ReceiveRequestDoActionIds: &lambdarestapi.ReceiveRequestDoActionModelIds_DEFAULT,
	SaveFileIds:               &bucket.SaveModelIds_DEFAULT,
}

var FuncionalityModelProps_DEFAULT FuncionalityModelProps = FuncionalityModelProps{
	StackProps:                  environment.StackProps_DEFAULT,
	ReceiveRequestDoActionProps: &lambdarestapi.ReceiveRequestDoActionModelProps_DEFAULT,
	SaveFileProps:               &bucket.SaveModelProps_DEFAULT,
}

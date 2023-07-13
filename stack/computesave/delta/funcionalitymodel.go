package delta

import (
	"castor/construct/highlevel/bucket"
	fargate "castor/construct/pattern/applicationloadbalancedfargateservice"
	"castor/stack/environment"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

type FuncionalityModelIds struct {
	StackId *string
	fargate.ReceiveRequestDoActionIds
	bucket.SaveFileIds
}

func (id *FuncionalityModelIds) Stack() *string {
	return id.StackId
}

func (id *FuncionalityModelIds) ReceiveRequestDoAction() fargate.ReceiveRequestDoActionIds {
	return id.ReceiveRequestDoActionIds
}

func (id *FuncionalityModelIds) SaveFile() bucket.SaveFileIds {
	return id.SaveFileIds
}

type FuncionalityModelProps struct {
	awscdk.StackProps
	fargate.ReceiveRequestDoActionProps
	bucket.SaveFileProps
}

func (props *FuncionalityModelProps) Stack() *awscdk.StackProps {
	return &props.StackProps
}

func (props *FuncionalityModelProps) ReceiveRequestDoAction() fargate.ReceiveRequestDoActionProps {
	return props.ReceiveRequestDoActionProps
}

func (props *FuncionalityModelProps) SaveFile() bucket.SaveFileProps {
	return props.SaveFileProps
}

type FuncionalityModel struct {
	awscdk.Stack
	receiverequestdoaction fargate.ReceiveRequestDoAction
	savefile               bucket.SaveFile
}

func (mo FuncionalityModel) ReceiveRequestDoAction() fargate.ReceiveRequestDoAction {
	return mo.receiverequestdoaction
}

func (mo FuncionalityModel) SaveFile() bucket.SaveFile {
	return mo.savefile
}

// SETTINGS
var FuncionalityModelIds_DEFAULT FuncionalityModelIds = FuncionalityModelIds{
	StackId:                   jsii.String("FuncionalityModel-default"),
	ReceiveRequestDoActionIds: &fargate.ReceiveRequestDoActionModelIds_DEFAULT,
	SaveFileIds:               &bucket.SaveModelIds_DEFAULT,
}

var FuncionalityModelProps_DEFAULT FuncionalityModelProps = FuncionalityModelProps{
	StackProps:                  environment.StackProps_DEFAULT,
	ReceiveRequestDoActionProps: &fargate.ReceiveRequestDoActionModelProps_DEFAULT,
	SaveFileProps:               &bucket.SaveModelProps_DEFAULT,
}

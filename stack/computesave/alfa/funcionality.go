package alfa

import (
	"castor/construct/pattern/bucket"
	"castor/construct/pattern/lambdarestapi"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type FuncionalityIds interface {
	Stack() *string
	ReceiveRequestDoAction() lambdarestapi.ReceiveRequestDoActionIds
	SaveFile() bucket.SaveFileIds
}

type FuncionalityProps interface {
	Stack() *awscdk.StackProps
	ReceiveRequestDoAction() lambdarestapi.ReceiveRequestDoActionProps
	SaveFile() bucket.SaveFileProps
	// Connections
}

type Funcionality interface{}

func NewFuncionality(scope constructs.Construct, id FuncionalityIds, props FuncionalityProps) Funcionality {
	var sprops FuncionalityProps = &FuncionalityModelProps_DEFAULT
	var sid FuncionalityIds = &FuncionalityModelIds_DEFAULT

	if props != nil {
		sprops = props
	}

	if id != nil {
		sid = id
	}

	stack := awscdk.NewStack(scope, sid.Stack(), sprops.Stack())

	doer := lambdarestapi.NewReceiveRequestDoAction(stack, sid.ReceiveRequestDoAction(), sprops.ReceiveRequestDoAction())

	saver := bucket.NewSaveFile(stack, sid.SaveFile(), sprops.SaveFile())

	saver.Bucket().GrantWrite(doer.DoAction().Function().Role(), jsii.String("*"), jsii.Strings("*"))
	saver.Choice().Service().GrantRead(doer.DoAction().Function().Role())
	saver.Choice().Service().GrantRead(doer.DoAction().Function().Role())

	var component Funcionality = &FuncionalityModel{
		Stack:                  stack,
		receiverequestdoaction: doer,
		savefile:               saver,
	}

	return component
}

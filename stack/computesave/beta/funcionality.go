package alfa

import (
	"castor/construct/pattern/lambdarestapi"
	"castor/construct/pattern/table"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

type FuncionalityIds interface {
	Stack() *string
	ReceiveRequestDoAction() lambdarestapi.ReceiveRequestDoActionIds
	SaveBlockData() table.SaveBlockDataIds
}

type FuncionalityProps interface {
	Stack() *awscdk.StackProps
	ReceiveRequestDoAction() lambdarestapi.ReceiveRequestDoActionProps
	SaveBlockData() table.SaveBlockDataProps
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

	saver := table.NewSaveBlockData(stack, sid.SaveBlockData(), sprops.SaveBlockData())

	saver.Table().GrantWriteData(doer.DoAction().Function())
	saver.Choice().Service().GrantRead(doer.DoAction().Function())
	saver.Choice().Service().GrantRead(doer.DoAction().Function())

	var component Funcionality = &FuncionalityModel{
		Stack:                  stack,
		receiverequestdoaction: doer,
		saveblockdata:          saver,
	}

	return component
}

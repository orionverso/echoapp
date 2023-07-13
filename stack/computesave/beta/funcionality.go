package alfa

import (
	"castor/construct/highlevel/table"
	"castor/construct/pattern/lambdarestapi"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
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

	save := table.NewSaveBlockData(stack, sid.SaveBlockData(), sprops.SaveBlockData())

	save.Table().GrantWriteData(doer.DoAction().Function())

	doer.DoAction().Function().AddEnvironment(
		jsii.String("STORAGE_SERVICE"),
		jsii.String("DYNAMODB"),
		&awslambda.EnvironmentOptions{})

	doer.DoAction().Function().AddEnvironment(
		jsii.String("DESTINATION"),
		save.Table().TableName(),
		&awslambda.EnvironmentOptions{})

	var component Funcionality = &FuncionalityModel{
		Stack:                  stack,
		receiverequestdoaction: doer,
		saveblockdata:          save,
	}

	return component
}

package delta

import (
	fargate "castor/construct/pattern/applicationloadbalancedfargateservice"
	"castor/construct/pattern/table"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

type FuncionalityIds interface {
	Stack() *string
	ReceiveRequestDoAction() fargate.ReceiveRequestDoActionIds
	SaveBlockData() table.SaveBlockDataIds
}

type FuncionalityProps interface {
	Stack() *awscdk.StackProps
	ReceiveRequestDoAction() fargate.ReceiveRequestDoActionProps
	SaveBlockData() table.SaveBlockDataProps
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

	doer := fargate.NewReceiveRequestDoAction(stack, sid.ReceiveRequestDoAction(), sprops.ReceiveRequestDoAction())
	doerrole := doer.ApplicationLoadBalancedFargateService().TaskDefinition().TaskRole()

	save := table.NewSaveBlockData(stack, sid.SaveBlockData(), sprops.SaveBlockData())

	save.Table().Grant(doerrole)
	save.Choice().Service().GrantRead(doerrole)
	save.Choice().Destination().GrantRead(doerrole)

	var component Funcionality = &FuncionalityModel{
		Stack:                  stack,
		receiverequestdoaction: doer,
		savefile:               save,
	}

	return component
}

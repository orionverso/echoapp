package gamma

import (
	"castor/construct/highlevel/table"
	fargate "castor/construct/pattern/applicationloadbalancedfargateservice"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
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

	save.Table().GrantWriteData(doerrole)

	defaultcontainer := doer.ApplicationLoadBalancedFargateService().TaskDefinition().DefaultContainer()
	defaultcontainer.AddEnvironment(jsii.String("STORAGE_SERVICE"), jsii.String("DYNAMODB"))
	defaultcontainer.AddEnvironment(jsii.String("DESTINATION"), save.Table().TableName())

	var component Funcionality = &FuncionalityModel{
		Stack:                  stack,
		receiverequestdoaction: doer,
		savefile:               save,
	}

	return component
}

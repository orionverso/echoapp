package delta

import (
	"castor/construct/highlevel/bucket"
	fargate "castor/construct/pattern/applicationloadbalancedfargateservice"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type FuncionalityIds interface {
	Stack() *string
	ReceiveRequestDoAction() fargate.ReceiveRequestDoActionIds
	SaveFile() bucket.SaveFileIds
}

type FuncionalityProps interface {
	Stack() *awscdk.StackProps
	ReceiveRequestDoAction() fargate.ReceiveRequestDoActionProps
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

	doer := fargate.NewReceiveRequestDoAction(stack, sid.ReceiveRequestDoAction(), sprops.ReceiveRequestDoAction())
	doerrole := doer.ApplicationLoadBalancedFargateService().TaskDefinition().TaskRole()

	save := bucket.NewSaveFile(stack, sid.SaveFile(), sprops.SaveFile())

	save.Bucket().GrantWrite(doerrole, jsii.String("*"), jsii.Strings("*"))

	defaultcontainer := doer.ApplicationLoadBalancedFargateService().TaskDefinition().DefaultContainer()
	defaultcontainer.AddEnvironment(jsii.String("STORAGE_SERVICE"), jsii.String("S3"))
	defaultcontainer.AddEnvironment(jsii.String("DESTINATION"), save.Bucket().BucketName())

	var component Funcionality = &FuncionalityModel{
		Stack:                  stack,
		receiverequestdoaction: doer,
		savefile:               save,
	}

	return component
}

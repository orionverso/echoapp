package alfa

import (
	"castor/construct/pattern/bucket"
	"castor/construct/pattern/lambdarestapi"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
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

	save := bucket.NewSaveFile(stack, sid.SaveFile(), sprops.SaveFile())

	save.Bucket().GrantWrite(doer.DoAction().Function(), jsii.String("*"), jsii.Strings("*"))

	doer.DoAction().Function().AddEnvironment(
		jsii.String("STORAGE_SERVICE"),
		jsii.String("S3"),
		&awslambda.EnvironmentOptions{})

	doer.DoAction().Function().AddEnvironment(
		jsii.String("DESTINATION"),
		save.Bucket().BucketName(),
		&awslambda.EnvironmentOptions{})

	var component Funcionality = &FuncionalityModel{
		Stack:                  stack,
		receiverequestdoaction: doer,
		savefile:               save,
	}

	return component
}

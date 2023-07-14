package computesave

import (
	fargate "castor/construct/pattern/applicationloadbalancedfargateservice"
	"castor/stack/environment"
	"log"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type FargateWriteToTableProps struct {
	awscdk.StackProps
	fargate.FargateEcrImageProps
	awsdynamodb.TableProps
}

type fargateWriteToTable struct {
	awscdk.Stack
	writer  fargate.FargateEcrImage
	storage awsdynamodb.Table
}

type FargateWriteToTable interface {
	awscdk.Stack
	Fargate() fargate.FargateEcrImage
	Table() awsdynamodb.Table
}

func NewFargateWriteToTable(scope constructs.Construct, id *string, props *FargateWriteToTableProps) FargateWriteToTable {
	var sprops *FargateWriteToTableProps = &FargateWriteToTableProps{}

	if id == nil {
		log.Panicln("parameter id is required, but nil was provided")
	}

	if props != nil {
		sprops = props
	}

	stack := awscdk.NewStack(scope, id, &sprops.StackProps)

	fg := fargate.NewFargateEcrImage(stack, jsii.String("Writer"), &sprops.FargateEcrImageProps)

	fgrole := fg.ApplicationLoadBalancedFargateService().TaskDefinition().TaskRole()

	tb := awsdynamodb.NewTable(stack, jsii.String("Storage"), &sprops.TableProps)

	tb.GrantWriteData(fgrole)

	defaultcontainer := fg.ApplicationLoadBalancedFargateService().TaskDefinition().DefaultContainer()
	defaultcontainer.AddEnvironment(jsii.String("STORAGE_SERVICE"), jsii.String("DYNAMODB"))
	defaultcontainer.AddEnvironment(jsii.String("DESTINATION"), tb.TableName())

	var component FargateWriteToTable = &fargateWriteToTable{
		Stack:   stack,
		writer:  fg,
		storage: tb,
	}

	return component
}

// IMPLEMENTATION
func (mo fargateWriteToTable) Fargate() fargate.FargateEcrImage {
	return mo.writer
}

func (mo fargateWriteToTable) Table() awsdynamodb.Table {
	return mo.storage
}

// SETTINGS
// DEVELOPMENT
var FargateWriteToTableProps_DEV FargateWriteToTableProps = FargateWriteToTableProps{
	StackProps:           environment.StackProps_DEV,
	FargateEcrImageProps: fargate.FargateEcrImageProps_DEV,
	TableProps: awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("id"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	},
}

// PRODUCTION
var FargateWriteToTableProps_PROD FargateWriteToTableProps = FargateWriteToTableProps{
	StackProps:           environment.StackProps_PROD,
	FargateEcrImageProps: fargate.FargateEcrImageProps_PROD,
	TableProps: awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("id"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	},
}

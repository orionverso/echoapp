package table

import (
	"log"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type SaveBlockDataProps struct {
	awsdynamodb.TableProps
}

type saveBlockData struct {
	table awsdynamodb.Table
}

type SaveBlockData interface {
	Table() awsdynamodb.Table
}

func NewSaveBlockData(scope constructs.Construct, id *string, props *SaveBlockDataProps) SaveBlockData {
	var sprops *SaveBlockDataProps = &SaveBlockDataProps{}

	if id == nil {
		log.Panicln("parameter id is required, but nil was provided")
	}

	if props != nil {
		sprops = props
	}

	this := constructs.NewConstruct(scope, id)

	table := awsdynamodb.NewTable(this, jsii.String("Table"), &sprops.TableProps)

	var component SaveBlockData = &saveBlockData{
		table: table,
	}

	return component
}

// IMPLEMENTATION
func (mo *saveBlockData) Table() awsdynamodb.Table {
	return mo.table
}

// SETTINGS
// DEVELOPMENT
var SaveBlockDataProps_DEV SaveBlockDataProps = SaveBlockDataProps{
	TableProps: awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("id"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	},
}

// PRODUCTION
var SaveBlockDataProps_PROD SaveBlockDataProps = SaveBlockDataProps{
	TableProps: awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("id"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	},
}

package table

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/jsii-runtime-go"
)

type SaveBlockModelIds struct {
	ConstructId *string
	TableId     *string
}

func (id *SaveBlockModelIds) Construct() *string {
	return id.ConstructId
}

func (id *SaveBlockModelIds) Table() *string {
	return id.TableId
}

type SaveBlockModelProps struct {
	*awsdynamodb.TableProps
}

func (props *SaveBlockModelProps) Table() *awsdynamodb.TableProps {
	return props.TableProps
}

type SaveBlockModel struct {
	table awsdynamodb.Table
}

func (mo *SaveBlockModel) Table() awsdynamodb.Table {
	return mo.table
}

// SETTINGS
var SaveBlockModelIds_DEFAULT SaveBlockModelIds = SaveBlockModelIds{
	ConstructId: jsii.String("SaveModel-default"),
	TableId:     jsii.String("table-default"),
}

var SaveBlockModelProps_DEFAULT SaveBlockModelProps = SaveBlockModelProps{
	TableProps: &awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("id"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	},
}

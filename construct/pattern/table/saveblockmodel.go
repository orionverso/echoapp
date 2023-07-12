package table

import (
	"castor/construct/pattern/choice"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/jsii-runtime-go"
)

type SaveBlockModelIds struct {
	ConstructId *string
	TableId     *string
	choice.DiscoverStorageIds
}

func (id *SaveBlockModelIds) Construct() *string {
	return id.ConstructId
}

func (id *SaveBlockModelIds) Table() *string {
	return id.TableId
}

func (id *SaveBlockModelIds) Choice() choice.DiscoverStorageIds {
	return id.DiscoverStorageIds
}

type SaveBlockModelProps struct {
	*awsdynamodb.TableProps
	choice.DiscoverStorageProps
}

func (props *SaveBlockModelProps) Table() *awsdynamodb.TableProps {
	return props.TableProps
}

func (props *SaveBlockModelProps) Choice() choice.DiscoverStorageProps {
	return props.DiscoverStorageProps
}

func (props *SaveBlockModelProps) AddDestinationToChoice(arn *string) {
	props.Destination().StringValue = arn
}

type SaveBlockModel struct {
	table  awsdynamodb.Table
	choice choice.DiscoverStorage
}

func (mo *SaveBlockModel) Table() awsdynamodb.Table {
	return mo.table
}

func (mo *SaveBlockModel) Choice() choice.DiscoverStorage {
	return mo.choice
}

// SETTINGS
var SaveBlockModelIds_DEFAULT SaveBlockModelIds = SaveBlockModelIds{
	ConstructId:        jsii.String("SaveModel-default"),
	TableId:            jsii.String("table-default"),
	DiscoverStorageIds: &choice.DiscoverModelIds_DEFAULT,
}

var SaveBlockModelProps_DEFAULT SaveBlockModelProps = SaveBlockModelProps{
	TableProps: &awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("id"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	},
	DiscoverStorageProps: &choice.DiscoverModelProps_DEFAULT,
}

package table

import (
	"castor/construct/pattern/choice"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/constructs-go/constructs/v10"
)

type SaveBlockDataIds interface {
	Construct() *string
	Table() *string
	Choice() choice.DiscoverStorageIds
}

type SaveBlockDataProps interface {
	Table() *awsdynamodb.TableProps
	Choice() choice.DiscoverStorageProps
	// connections
	AddDestinationToChoice(*string)
}

type SaveBlockData interface {
	Table() awsdynamodb.Table
	Choice() choice.DiscoverStorage
}

func NewSaveBlockData(scope constructs.Construct, id SaveBlockDataIds, props SaveBlockDataProps) SaveBlockData {
	var sprops SaveBlockDataProps = &SaveBlockModelProps_DEFAULT
	var sid SaveBlockDataIds = &SaveBlockModelIds_DEFAULT

	if props != nil {
		sprops = props
	}

	if id != nil {
		sid = id
	}

	this := constructs.NewConstruct(scope, sid.Construct())

	table := awsdynamodb.NewTable(this, sid.Table(), sprops.Table())

	sprops.AddDestinationToChoice(table.TableArn())

	choice := choice.NewDiscoverStorage(this, sid.Choice(), sprops.Choice())

	var component SaveBlockData = &SaveBlockModel{
		table:  table,
		choice: choice,
	}

	return component
}

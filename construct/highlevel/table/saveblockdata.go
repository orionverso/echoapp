package table

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/constructs-go/constructs/v10"
)

type SaveBlockDataIds interface {
	Construct() *string
	Table() *string
}

type SaveBlockDataProps interface {
	Table() *awsdynamodb.TableProps
	// connections
}

type SaveBlockData interface {
	Table() awsdynamodb.Table
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

	// this := constructs.NewConstruct(scope, sid.Construct())
	this := constructs.NewConstruct(scope, sid.Construct())

	table := awsdynamodb.NewTable(this, sid.Table(), sprops.Table())

	var component SaveBlockData = &SaveBlockModel{
		table: table,
	}

	return component
}

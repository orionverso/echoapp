package table

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
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
	Writer awsiam.IRole
}

func (props *SaveBlockModelProps) Table() *awsdynamodb.TableProps {
	return props.TableProps
}

func (props *SaveBlockModelProps) HasWriterRole() bool {
	if props.Writer == nil {
		return false
	} else {
		return true
	}
}

func (props *SaveBlockModelProps) WriterRole() awsiam.IRole {
	return props.Writer
}

type SaveBlockModel struct {
	table awsdynamodb.Table
}

func (mo SaveBlockModel) Table() awsdynamodb.Table {
	return mo.table
}

// SETTINGS
var SaveBlockModelIds_DEFAULT SaveBlockModelIds = SaveBlockModelIds{
	ConstructId: jsii.String("MODEL-table-construct-default"),
	TableId:     jsii.String("MODEL-table-resource-default"),
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

package table

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/jsii-runtime-go"
)

type SaveEchoIds struct {
	SaveBlockModelIds
}

type SaveEchoProps struct {
	SaveBlockModelProps
}

// SETTINGS
// DEVELOPMENT
var SaveEchoIds_DEV SaveEchoIds = SaveEchoIds{
	SaveBlockModelIds: SaveBlockModelIds{
		TableId:     jsii.String("SaveEcho-table-resource-dev"),
		ConstructId: jsii.String("SaveEcho-construct-dev"),
	},
}

var SaveEchoProps_DEV SaveEchoProps = SaveEchoProps{
	SaveBlockModelProps: SaveBlockModelProps{
		TableProps: &awsdynamodb.TableProps{
			PartitionKey: &awsdynamodb.Attribute{
				Name: jsii.String("id"),
				Type: awsdynamodb.AttributeType_STRING,
			},
			RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
		},
	},
}

// PRODUCTION
var SaveEchoIds_PROD SaveEchoIds = SaveEchoIds{
	SaveBlockModelIds: SaveBlockModelIds{
		TableId:     jsii.String("SaveEcho-table-resource-prod"),
		ConstructId: jsii.String("SaveEcho-construct-prod"),
	},
}

var SaveEchoProps_PROD SaveEchoProps = SaveEchoProps{
	SaveBlockModelProps: SaveBlockModelProps{
		TableProps: &awsdynamodb.TableProps{
			PartitionKey: &awsdynamodb.Attribute{
				Name: jsii.String("id"),
				Type: awsdynamodb.AttributeType_STRING,
			},
			RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
		},
	},
}

package table

import (
	"castor/construct/pattern/choice"

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
		ConstructId:        jsii.String("SaveEcho-dev"),
		TableId:            jsii.String("table-dev"),
		DiscoverStorageIds: &choice.DiscoverDynamoDbIds_DEV,
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
		DiscoverStorageProps: &choice.DiscoverDynamoDbProps_DEV,
	},
}

// PRODUCTION
var SaveEchoIds_PROD SaveEchoIds = SaveEchoIds{
	SaveBlockModelIds: SaveBlockModelIds{
		TableId:            jsii.String("SaveEcho-prod"),
		ConstructId:        jsii.String("SaveEcho-prod"),
		DiscoverStorageIds: &choice.DiscoverDynamoDbIds_PROD,
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
		DiscoverStorageProps: &choice.DiscoverDynamoDbProps_PROD,
	},
}

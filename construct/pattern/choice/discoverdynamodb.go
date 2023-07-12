package choice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
	"github.com/aws/jsii-runtime-go"
)

type DiscoverDynamoDbIds struct {
	DiscoverModelIds
}

type DiscoverDynamoDbProps struct {
	DiscoverModelProps
}

// SETTINGS
// DEVELOPMENT
var DiscoverDynamoDbIds_DEV DiscoverDynamoDbIds = DiscoverDynamoDbIds{
	DiscoverModelIds: DiscoverModelIds{
		ConstructId:   jsii.String("DiscoverStorage-dev"),
		ServiceId:     jsii.String("serviceparameter-dev"),
		DestinationId: jsii.String("destinationparameter-dev"),
	},
}

var DiscoverDynamoDbProps_DEV DiscoverDynamoDbProps = DiscoverDynamoDbProps{
	DiscoverModelProps: DiscoverModelProps{
		ServiceProps: &awsssm.StringParameterProps{
			ParameterName: jsii.String("STORAGE_SOLUTION"),
			StringValue:   jsii.String("DYNAMODB"),
		}, DestinationProps: &awsssm.StringParameterProps{
			ParameterName: jsii.String("DESTINATION"),
			StringValue:   jsii.String("UNKNOW-STORAGE-SERVICE"), // overwrite at runtime
		},
	},
}

// PRODUCTION
var DiscoverDynamoDbIds_PROD DiscoverDynamoDbIds = DiscoverDynamoDbIds{
	DiscoverModelIds: DiscoverModelIds{
		ConstructId:   jsii.String("DiscoverStorage-prod"),
		ServiceId:     jsii.String("serviceparameter-prod"),
		DestinationId: jsii.String("destinationparameter-prod"),
	},
}

var DiscoverDynamoDbProps_PROD DiscoverDynamoDbProps = DiscoverDynamoDbProps{
	DiscoverModelProps: DiscoverModelProps{
		ServiceProps: &awsssm.StringParameterProps{
			ParameterName: jsii.String("STORAGE_SOLUTION"),
			StringValue:   jsii.String("DYNAMODB"),
		},
		DestinationProps: &awsssm.StringParameterProps{
			ParameterName: jsii.String("DESTINATION"),
			StringValue:   jsii.String("UNKNOW-STORAGE-SERVICE"), // overwrite at runtime
		},
	},
}

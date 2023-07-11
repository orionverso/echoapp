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
		ServiceId:     jsii.String("Discover-serviceparameter-dev"),
		DestinationId: jsii.String("Discover-destinationparameter-dev"),
	},
}

var DiscoverDynamoDbProps_DEV DiscoverDynamoDbProps = DiscoverDynamoDbProps{
	DiscoverModelProps: DiscoverModelProps{
		ServiceProps: &awsssm.StringParameterProps{
			ParameterName: jsii.String("STORAGE_SOLUTION"),
			StringValue:   jsii.String("DYNAMODB"),
		}, DestinationProps: &awsssm.StringParameterProps{
			ParameterName: jsii.String("DESTINATION"),
			// StringValue:   At runtime,
		},
	},
}

// PRODUCTION
var DiscoverDynamoDbIds_PROD DiscoverDynamoDbIds = DiscoverDynamoDbIds{
	DiscoverModelIds: DiscoverModelIds{
		ServiceId:     jsii.String("Discover-serviceparameter-prod"),
		DestinationId: jsii.String("Discover-destinationparameter-prod"),
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
			// StringValue:   At runtime,
		},
	},
}

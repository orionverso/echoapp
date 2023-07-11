package choice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
	"github.com/aws/jsii-runtime-go"
)

type DiscoverS3Ids struct {
	DiscoverModelIds
}

type DiscoverS3Props struct {
	DiscoverModelProps
}

// SETTINGS
// DEVELOPMENT
var DiscoverS3Ids_DEV DiscoverS3Ids = DiscoverS3Ids{
	DiscoverModelIds: DiscoverModelIds{
		ServiceId:     jsii.String("DiscoverS3-serviceparameter-dev"),
		DestinationId: jsii.String("DiscoverS3-destinationparameter-dev"),
	},
}

var DiscoverS3Props_DEV DiscoverS3Props = DiscoverS3Props{
	DiscoverModelProps: DiscoverModelProps{
		ServiceProps: &awsssm.StringParameterProps{
			ParameterName: jsii.String("STORAGE_SOLUTION"),
			StringValue:   jsii.String("S3"),
		},
		DestinationProps: &awsssm.StringParameterProps{
			ParameterName: jsii.String("DESTINATION"),
			// StringValue:   At runtime,
		},
	},
}

// PRODUCTION
var DiscoverS3Ids_PROD DiscoverS3Ids = DiscoverS3Ids{
	DiscoverModelIds: DiscoverModelIds{
		ServiceId:     jsii.String("DiscoverS3-serviceparameter-prod"),
		DestinationId: jsii.String("DiscoverS3-destinationparameter-prod"),
	},
}

var DiscoverS3Props_PROD DiscoverS3Props = DiscoverS3Props{
	DiscoverModelProps: DiscoverModelProps{
		ServiceProps: &awsssm.StringParameterProps{
			ParameterName: jsii.String("STORAGE_SOLUTION"),
			StringValue:   jsii.String("S3"),
		},
		DestinationProps: &awsssm.StringParameterProps{
			ParameterName: jsii.String("DESTINATION"),
			// StringValue:   At runtime,
		},
	},
}

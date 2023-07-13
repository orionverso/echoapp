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
		ConstructId:   jsii.String("DiscoverStorage-dev"),
		ServiceId:     jsii.String("serviceparameter-dev"),
		DestinationId: jsii.String("destinationparameter-dev"),
	},
}

var DiscoverS3Props_DEV DiscoverS3Props = DiscoverS3Props{
	DiscoverModelProps: DiscoverModelProps{
		ServiceProps: &awsssm.StringParameterProps{
			ParameterName: jsii.String("STORAGE_SERVICE"),
			StringValue:   jsii.String("S3"),
		},
		DestinationProps: &awsssm.StringParameterProps{
			ParameterName: jsii.String("DESTINATION"),
			StringValue:   jsii.String("UNKNOW-STORAGE-SERVICE"), // overwrite at runtime
		},
	},
}

// PRODUCTION
var DiscoverS3Ids_PROD DiscoverS3Ids = DiscoverS3Ids{
	DiscoverModelIds: DiscoverModelIds{
		ConstructId:   jsii.String("DiscoverStorage-prod"),
		ServiceId:     jsii.String("serviceparameter-prod"),
		DestinationId: jsii.String("destinationparameter-prod"),
	},
}

var DiscoverS3Props_PROD DiscoverS3Props = DiscoverS3Props{
	DiscoverModelProps: DiscoverModelProps{
		ServiceProps: &awsssm.StringParameterProps{
			ParameterName: jsii.String("STORAGE_SERVICE"),
			StringValue:   jsii.String("S3"),
		},
		DestinationProps: &awsssm.StringParameterProps{
			ParameterName: jsii.String("DESTINATION"),
			StringValue:   jsii.String("UNKNOW-STORAGE-SERVICE"), // overwrite at runtime
		},
	},
}

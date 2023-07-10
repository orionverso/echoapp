package bucket

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/jsii-runtime-go"
)

type SaveEchoIds struct {
	SaveModelIds
}

type SaveEchoProps struct {
	SaveModelProps
}

// SETTINGS
// DEVELOPMENT
var SaveEchoIds_DEV SaveEchoIds = SaveEchoIds{
	SaveModelIds: SaveModelIds{
		ConstructId: jsii.String("SaveEcho-resource-construct-default"),
		BucketId:    jsii.String("SaveEcho-resource-bucket-dev"),
	},
}

var SaveEchoProps_DEV SaveEchoProps = SaveEchoProps{
	SaveModelProps: SaveModelProps{
		BucketProps: &awss3.BucketProps{
			AutoDeleteObjects: jsii.Bool(true),
			RemovalPolicy:     awscdk.RemovalPolicy_DESTROY,
		},
	},
}

// PRODUCTION
var SaveEchoIds_PROD SaveEchoIds = SaveEchoIds{
	SaveModelIds: SaveModelIds{
		ConstructId: jsii.String("SaveEcho-resource-construct-default"),
		BucketId:    jsii.String("SaveEcho-resource-bucket-prod"),
	},
}

var SaveEchoProps_PROD SaveEchoProps = SaveEchoProps{
	SaveModelProps: SaveModelProps{
		BucketProps: &awss3.BucketProps{
			AutoDeleteObjects: jsii.Bool(true),
			RemovalPolicy:     awscdk.RemovalPolicy_DESTROY,
		},
	},
}

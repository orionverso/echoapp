package repository

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/jsii-runtime-go"
)

type DistributeGoServerImageIds struct {
	DistributeModelIds
}

type DistributeGoServerImageProps struct {
	DistributeModelProps
}

// SETTINGS
// DEVELOPMENT
var DistributeGoServerImageIds_DEV DistributeGoServerImageIds = DistributeGoServerImageIds{
	DistributeModelIds: DistributeModelIds{
		ConstructId:  jsii.String("DistributeGoServerImage-dev"),
		RepositoryId: jsii.String("repository-dev"),
	},
}

var DistributeGoServerImageProps_DEV DistributeGoServerImageProps = DistributeGoServerImageProps{
	DistributeModelProps: DistributeModelProps{
		RepositoryProps: &awsecr.RepositoryProps{
			AutoDeleteImages: jsii.Bool(true),
			Encryption:       awsecr.RepositoryEncryption_KMS(),
			RemovalPolicy:    awscdk.RemovalPolicy_DESTROY,
			RepositoryName:   jsii.String("go-server-images-dev"),
		},
	},
}

// PRODUCTION
var DistributeGoServerImageIds_PROD DistributeGoServerImageIds = DistributeGoServerImageIds{
	DistributeModelIds: DistributeModelIds{
		ConstructId:  jsii.String("DistributeGoServerImage-prod"),
		RepositoryId: jsii.String("repository-prod"),
	},
}

var DistributeGoServerImageProps_PROD DistributeGoServerImageProps = DistributeGoServerImageProps{
	DistributeModelProps: DistributeModelProps{
		RepositoryProps: &awsecr.RepositoryProps{
			AutoDeleteImages: jsii.Bool(true),
			Encryption:       awsecr.RepositoryEncryption_KMS(),
			RemovalPolicy:    awscdk.RemovalPolicy_DESTROY,
			RepositoryName:   jsii.String("go-server-images-prod"),
		},
	},
}

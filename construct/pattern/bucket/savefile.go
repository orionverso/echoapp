package bucket

import (
	"castor/construct/pattern/choice"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

type SaveFileIds interface {
	Construct() *string
	Bucket() *string
	Choice() choice.DiscoverStorageIds
}

type SaveFileProps interface {
	Bucket() *awss3.BucketProps
	Choice() choice.DiscoverStorageProps
	// connections
	AddDestinationToChoice(*string)
}

type SaveFile interface {
	Bucket() awss3.Bucket
	Choice() choice.DiscoverStorage
}

func NewSaveFile(scope constructs.Construct, id SaveFileIds, props SaveFileProps) SaveFile {
	var sprops SaveFileProps = &SaveModelProps_DEFAULT
	var sid SaveFileIds = &SaveModelIds_DEFAULT

	if props != nil {
		sprops = props
	}

	if id != nil {
		sid = id
	}

	this := constructs.NewConstruct(scope, sid.Construct())

	resource := awss3.NewBucket(this, sid.Bucket(), sprops.Bucket())

	sprops.AddDestinationToChoice(resource.BucketName())

	ch := choice.NewDiscoverStorage(this, sid.Choice(), sprops.Choice())

	var component SaveFile = &SaveModel{
		bucket: resource,
		choice: ch,
	}

	return component
}

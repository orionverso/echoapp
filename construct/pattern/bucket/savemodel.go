package bucket

import (
	"castor/construct/pattern/choice"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/jsii-runtime-go"
)

type SaveModelIds struct {
	ConstructId *string
	BucketId    *string
	choice.DiscoverStorageIds
}

func (id *SaveModelIds) Bucket() *string {
	return id.BucketId
}

func (id *SaveModelIds) Construct() *string {
	return id.BucketId
}

func (id *SaveModelIds) Choice() choice.DiscoverStorageIds {
	return id.DiscoverStorageIds
}

type SaveModelProps struct {
	*awss3.BucketProps
	choice.DiscoverStorageProps
}

func (props *SaveModelProps) Bucket() *awss3.BucketProps {
	return props.BucketProps
}

func (props *SaveModelProps) Choice() choice.DiscoverStorageProps {
	return props.DiscoverStorageProps
}

func (props *SaveModelProps) AddDestinationToChoice(arn *string) {
	props.DiscoverStorageProps.Destination().StringValue = arn
}

type SaveModel struct {
	bucket awss3.Bucket
	choice choice.DiscoverStorage
}

func (mo *SaveModel) Bucket() awss3.Bucket {
	return mo.bucket
}

func (mo *SaveModel) Choice() choice.DiscoverStorage {
	return mo.choice
}

// SETTINGS
var SaveModelIds_DEFAULT SaveModelIds = SaveModelIds{
	ConstructId:        jsii.String("MODEL-resource-construct-default"),
	BucketId:           jsii.String("MODEL-resource-bucket-default"),
	DiscoverStorageIds: &choice.DiscoverModelIds_DEFAULT,
}

var SaveModelProps_DEFAULT SaveModelProps = SaveModelProps{
	BucketProps: &awss3.BucketProps{
		AutoDeleteObjects: jsii.Bool(true),
		RemovalPolicy:     awscdk.RemovalPolicy_DESTROY,
	},
	DiscoverStorageProps: &choice.DiscoverModelProps_DEFAULT,
}

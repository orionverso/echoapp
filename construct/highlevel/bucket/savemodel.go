package bucket

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/jsii-runtime-go"
)

type SaveModelIds struct {
	ConstructId *string
	BucketId    *string
}

func (id *SaveModelIds) Bucket() *string {
	return id.BucketId
}

func (id *SaveModelIds) Construct() *string {
	return id.ConstructId
}

type SaveModelProps struct {
	*awss3.BucketProps
}

func (props *SaveModelProps) Bucket() *awss3.BucketProps {
	return props.BucketProps
}

type SaveModel struct {
	bucket awss3.Bucket
}

func (mo *SaveModel) Bucket() awss3.Bucket {
	return mo.bucket
}

// SETTINGS
var SaveModelIds_DEFAULT SaveModelIds = SaveModelIds{
	ConstructId: jsii.String("SaveModel-default"),
	BucketId:    jsii.String("bucket-default"),
}

var SaveModelProps_DEFAULT SaveModelProps = SaveModelProps{
	BucketProps: &awss3.BucketProps{
		AutoDeleteObjects: jsii.Bool(true),
		RemovalPolicy:     awscdk.RemovalPolicy_DESTROY,
	},
}

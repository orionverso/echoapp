package bucket

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type SaveFileIds interface {
	Construct() *string
	Bucket() *string
}

type SaveFileProps interface {
	Bucket() *awss3.BucketProps
	WriterRole() awsiam.IRole // At runtime
	HasWriterRole() bool
}

type SaveFile interface {
	Bucket() awss3.Bucket
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

	if sprops.HasWriterRole() {
		resource.GrantWrite(sprops.WriterRole(), jsii.String("*"), jsii.Strings("*"))
	}

	var component SaveFile = &SaveModel{
		bucket: resource,
	}

	return component
}

package bucket

import (
	"log"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type SaveObjectProps struct {
	awss3.BucketProps
}

type saveFile struct {
	constructs.Construct
	bucket awss3.Bucket
}

type SaveObject interface {
	Bucket() awss3.Bucket
}

func NewSaveObject(scope constructs.Construct, id *string, props *SaveObjectProps) SaveObject {
	var sprops *SaveObjectProps = &SaveObjectProps{}

	if id == nil {
		log.Panicln("parameter id is required, but nil was provided")
	}

	if props != nil {
		sprops = props
	}

	this := constructs.NewConstruct(scope, id)

	// RESOURCE DECLARATIONS...
	bk := awss3.NewBucket(this, jsii.String("Bucket"), &sprops.BucketProps)

	var component SaveObject = &saveFile{
		Construct: this,
		bucket:    bk,
	}

	return component
}

// IMPLEMENTATION
func (sa *saveFile) Bucket() awss3.Bucket {
	return sa.bucket
}

// SETTINGS
// DEVELOPMENT
var SaveObjectProps_DEV SaveObjectProps = SaveObjectProps{
	BucketProps: awss3.BucketProps{
		AutoDeleteObjects: jsii.Bool(true),
		RemovalPolicy:     awscdk.RemovalPolicy_DESTROY,
	},
}

// PRODUCTION
var SaveObjectProps_PROD SaveObjectProps = SaveObjectProps{
	BucketProps: awss3.BucketProps{
		AutoDeleteObjects: jsii.Bool(true),
		RemovalPolicy:     awscdk.RemovalPolicy_DESTROY,
	},
}

package repository

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/constructs-go/constructs/v10"
)

type DistributeImageIds interface {
	Construct() *string
	Repository() *string
}

type DistributeImageProps interface {
	Repository() *awsecr.RepositoryProps
}

type DistributeImage interface {
	Repository() awsecr.Repository
}

func NewDistributeImage(scope constructs.Construct, id DistributeImageIds, props DistributeImageProps) DistributeImage {
	var sprops DistributeImageProps = &DistributeModelProps_DEFAULT
	var sid DistributeImageIds = &DistributeModelIds_DEFAULT

	if props != nil {
		sprops = props
	}

	if id != nil {
		sid = id
	}

	this := constructs.NewConstruct(scope, sid.Construct())

	resource := awsecr.NewRepository(this, sid.Repository(), sprops.Repository())

	var component DistributeImage = &DistributeModel{
		repository: resource,
	}

	return component
}

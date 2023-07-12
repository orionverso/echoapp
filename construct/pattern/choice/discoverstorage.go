package choice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
	"github.com/aws/constructs-go/constructs/v10"
)

type DiscoverStorageIds interface {
	Construct() *string
	Service() *string
	Destination() *string
}

type DiscoverStorageProps interface {
	Service() *awsssm.StringParameterProps
	Destination() *awsssm.StringParameterProps
}

type DiscoverStorage interface {
	constructs.Construct
	Service() awsssm.StringParameter
	Destination() awsssm.StringParameter
}

func NewDiscoverStorage(scope constructs.Construct, id DiscoverStorageIds, props DiscoverStorageProps) DiscoverStorage {
	var sprops DiscoverStorageProps = &DiscoverModelProps_DEFAULT
	var sid DiscoverStorageIds = &DiscoverModelIds_DEFAULT

	if props != nil {
		sprops = props
	}

	if id != nil {
		sid = id
	}

	this := constructs.NewConstruct(scope, sid.Construct())

	service := awsssm.NewStringParameter(this, sid.Service(), sprops.Service())
	destination := awsssm.NewStringParameter(this, sid.Destination(), sprops.Destination())

	var component DiscoverStorage = &DiscoverModel{
		Construct:   this,
		service:     service,
		destination: destination,
	}

	return component
}

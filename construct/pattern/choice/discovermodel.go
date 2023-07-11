package choice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
	"github.com/aws/jsii-runtime-go"
)

type DiscoverModelIds struct {
	ConstructId   *string
	ServiceId     *string
	DestinationId *string
}

func (id *DiscoverModelIds) Construct() *string {
	return id.ConstructId
}

func (id *DiscoverModelIds) Service() *string {
	return id.ServiceId
}

func (id *DiscoverModelIds) Destination() *string {
	return id.DestinationId
}

type DiscoverModelProps struct {
	ServiceProps     *awsssm.StringParameterProps
	DestinationProps *awsssm.StringParameterProps
}

func (props *DiscoverModelProps) Service() *awsssm.StringParameterProps {
	return props.ServiceProps
}

func (props *DiscoverModelProps) Destination() *awsssm.StringParameterProps {
	return props.DestinationProps
}

type DiscoverModel struct {
	service     awsssm.StringParameter
	destination awsssm.StringParameter
}

func (mo *DiscoverModel) Service() awsssm.StringParameter {
	return mo.service
}

func (mo *DiscoverModel) Destination() awsssm.StringParameter {
	return mo.destination
}

// SETTINGS
var DiscoverModelIds_DEFAULT DiscoverModelIds = DiscoverModelIds{
	ServiceId:     jsii.String("DiscoverModel-serviceparameter-default"),
	DestinationId: jsii.String("DiscoverModel-destinationparameter-default"),
}

var DiscoverModelProps_DEFAULT DiscoverModelProps = DiscoverModelProps{
	ServiceProps: &awsssm.StringParameterProps{
		ParameterName: jsii.String("STORAGE_SOLUTION"),
		StringValue:   jsii.String("UNKNOW-STORAGE-SERVICE"),
	},
	DestinationProps: &awsssm.StringParameterProps{
		ParameterName: jsii.String("DESTINATION"),
		StringValue:   jsii.String("UNKNOW-DESTINATION-ARN"),
	},
}

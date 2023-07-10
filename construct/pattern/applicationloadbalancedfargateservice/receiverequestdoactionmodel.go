package applicationloadbalancedfargateservice

import (
	"castor/construct/pattern/repository"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecspatterns"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type ReceiveRequestDoActionModelIds struct {
	ConstructId                             *string
	ApplicationLoadBalancedFargateServiceId *string
	repository.DistributeImageIds
}

func (id *ReceiveRequestDoActionModelIds) Construct() *string {
	return id.ConstructId
}

func (id *ReceiveRequestDoActionModelIds) ApplicationLoadBalancedFargateService() *string {
	return id.ApplicationLoadBalancedFargateServiceId
}

func (id *ReceiveRequestDoActionModelIds) DistributeImage() repository.DistributeImageIds {
	return id.DistributeImageIds
}

type ReceiveRequestDoActionModelProps struct {
	*awsecspatterns.ApplicationLoadBalancedFargateServiceProps
	repository.DistributeImageProps
	TagImage *string
}

func (props *ReceiveRequestDoActionModelProps) ApplicationLoadBalancedFargateService() *awsecspatterns.ApplicationLoadBalancedFargateServiceProps {
	return props.ApplicationLoadBalancedFargateServiceProps
}

func (props *ReceiveRequestDoActionModelProps) DistributeImage() repository.DistributeImageProps {
	return props.DistributeImageProps
}

func (props *ReceiveRequestDoActionModelProps) Tag() *string {
	return props.TagImage
}

func (props *ReceiveRequestDoActionModelProps) AddContainerImageToApplicationLoadBalancedFargate(image awsecs.ContainerImage) {
	var taskImageOptions awsecspatterns.ApplicationLoadBalancedTaskImageOptions
	taskImageOptions.Image = image
	props.ApplicationLoadBalancedFargateServiceProps.TaskImageOptions = &taskImageOptions
}

type ReceiveRequestDoActionModel struct {
	constructs.Construct
	applicationloadbalancedfargateservice awsecspatterns.ApplicationLoadBalancedFargateService
}

func (mo *ReceiveRequestDoActionModel) ApplicationLoadBalancedFargateService() awsecspatterns.ApplicationLoadBalancedFargateService {
	return mo.applicationloadbalancedfargateservice
}

// SETTINGS
var ReceiveRequestDoActionModelIds_DEFAULT ReceiveRequestDoActionModelIds = ReceiveRequestDoActionModelIds{
	ConstructId:                             jsii.String("MODEL-resource-construct-default"),
	ApplicationLoadBalancedFargateServiceId: jsii.String("MODEL-resource-applicationloadbalancedfargateservice-default"),
	DistributeImageIds:                      &repository.DistributeModelIds_DEFAULT,
}

var ReceiveRequestDoActionModelProps_DEFAULT ReceiveRequestDoActionModelProps = ReceiveRequestDoActionModelProps{
	ApplicationLoadBalancedFargateServiceProps: &awsecspatterns.ApplicationLoadBalancedFargateServiceProps{
		MemoryLimitMiB:   jsii.Number(1024),
		DesiredCount:     jsii.Number(1),
		Cpu:              jsii.Number(512),
		LoadBalancerName: jsii.String("echoapp-alb-default"),
	},
	DistributeImageProps: &repository.DistributeModelProps_DEFAULT,
	TagImage:             jsii.String("latest"),
}

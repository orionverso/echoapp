package applicationloadbalancedfargateservice

import (
	"log"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecspatterns"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type FargateProps struct {
	awsecspatterns.ApplicationLoadBalancedFargateServiceProps
}

type fargateEcrImage struct {
	constructs.Construct
	applicationloadbalancedfargateservice awsecspatterns.ApplicationLoadBalancedFargateService
}

type Fargate interface {
	constructs.Construct
	ApplicationLoadBalancedFargateService() awsecspatterns.ApplicationLoadBalancedFargateService
}

func NewFargate(scope constructs.Construct, id *string, props *FargateProps) Fargate {
	var sprops *FargateProps = &FargateProps{}

	if id == nil {
		log.Panicln("parameter id is required, but nil was provided")
	}

	if props != nil {
		sprops = props
	}

	this := constructs.NewConstruct(scope, id)

	fargate := awsecspatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("ApplicationLoadBalancedFargateService"),
		&sprops.ApplicationLoadBalancedFargateServiceProps)

	var component Fargate = &fargateEcrImage{
		Construct:                             this,
		applicationloadbalancedfargateservice: fargate,
	}
	return component
}

// PROPS
func (props *FargateProps) AddContainerImageToApplicationLoadBalancedFargate(image awsecs.ContainerImage) {
	var taskImageOptions awsecspatterns.ApplicationLoadBalancedTaskImageOptions
	taskImageOptions.Image = image
	props.ApplicationLoadBalancedFargateServiceProps.TaskImageOptions = &taskImageOptions
}

// IMPLEMENTATION
func (fa *fargateEcrImage) ApplicationLoadBalancedFargateService() awsecspatterns.ApplicationLoadBalancedFargateService {
	return fa.applicationloadbalancedfargateservice
}

// SETTINGS
// DEVELOPMENT
var FargateProps_DEV FargateProps = FargateProps{
	ApplicationLoadBalancedFargateServiceProps: awsecspatterns.ApplicationLoadBalancedFargateServiceProps{
		MemoryLimitMiB:   jsii.Number(1024),
		DesiredCount:     jsii.Number(1),
		Cpu:              jsii.Number(512),
		LoadBalancerName: jsii.String("echoapp-alb-dev"),
	},
}

// PRODUCTION
var FargateProps_PROD FargateProps = FargateProps{
	ApplicationLoadBalancedFargateServiceProps: awsecspatterns.ApplicationLoadBalancedFargateServiceProps{
		MemoryLimitMiB:   jsii.Number(1024),
		DesiredCount:     jsii.Number(1),
		Cpu:              jsii.Number(512),
		LoadBalancerName: jsii.String("echoapp-alb-prod"),
	},
}

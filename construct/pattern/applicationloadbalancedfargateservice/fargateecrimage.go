package applicationloadbalancedfargateservice

import (
	"log"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecspatterns"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type FargateEcrImageProps struct {
	awsecspatterns.ApplicationLoadBalancedFargateServiceProps
	awsecr.RepositoryProps
	TagImage string
}

type fargateEcrImage struct {
	constructs.Construct
	applicationloadbalancedfargateservice awsecspatterns.ApplicationLoadBalancedFargateService
	repository                            awsecr.Repository
}

type FargateEcrImage interface {
	constructs.Construct
	ApplicationLoadBalancedFargateService() awsecspatterns.ApplicationLoadBalancedFargateService
	Repository() awsecr.Repository
}

func NewFargateEcrImage(scope constructs.Construct, id *string, props *FargateEcrImageProps) FargateEcrImage {
	var sprops *FargateEcrImageProps = &FargateEcrImageProps{}

	if id == nil {
		log.Panicln("parameter id is required, but nil was provided")
	}

	if props != nil {
		sprops = props
	}

	this := constructs.NewConstruct(scope, id)

	repo := awsecr.NewRepository(this, jsii.String("Repository"), &sprops.RepositoryProps)

	image := awsecs.AssetImage_FromEcrRepository(repo, jsii.String(sprops.TagImage))

	sprops.AddContainerImageToApplicationLoadBalancedFargate(image)

	fargate := awsecspatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("ApplicationLoadBalancedFargateService"),
		&sprops.ApplicationLoadBalancedFargateServiceProps)

	var component FargateEcrImage = &fargateEcrImage{
		Construct:                             this,
		applicationloadbalancedfargateservice: fargate,
		repository:                            repo,
	}
	return component
}

// PROPS
func (props *FargateEcrImageProps) AddContainerImageToApplicationLoadBalancedFargate(image awsecs.ContainerImage) {
	var taskImageOptions awsecspatterns.ApplicationLoadBalancedTaskImageOptions
	taskImageOptions.Image = image
	props.ApplicationLoadBalancedFargateServiceProps.TaskImageOptions = &taskImageOptions
}

// IMPLEMENTATION
func (fa *fargateEcrImage) ApplicationLoadBalancedFargateService() awsecspatterns.ApplicationLoadBalancedFargateService {
	return fa.applicationloadbalancedfargateservice
}

func (fa *fargateEcrImage) Repository() awsecr.Repository {
	return fa.repository
}

// SETTINGS
// DEVELOPMENT
var FargateEcrImageProps_DEV FargateEcrImageProps = FargateEcrImageProps{
	ApplicationLoadBalancedFargateServiceProps: awsecspatterns.ApplicationLoadBalancedFargateServiceProps{
		MemoryLimitMiB:   jsii.Number(1024),
		DesiredCount:     jsii.Number(1),
		Cpu:              jsii.Number(512),
		LoadBalancerName: jsii.String("echoapp-alb-dev"),
	},
	RepositoryProps: awsecr.RepositoryProps{},
	TagImage:        "latest",
}

// PRODUCTION
var FargateEcrImageProps_PROD FargateEcrImageProps = FargateEcrImageProps{
	ApplicationLoadBalancedFargateServiceProps: awsecspatterns.ApplicationLoadBalancedFargateServiceProps{
		MemoryLimitMiB:   jsii.Number(1024),
		DesiredCount:     jsii.Number(1),
		Cpu:              jsii.Number(512),
		LoadBalancerName: jsii.String("echoapp-alb-prod"),
	},
	RepositoryProps: awsecr.RepositoryProps{},
	TagImage:        "latest",
}

package applicationloadbalancedfargateservice

import (
	"castor/construct/pattern/repository"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecspatterns"
	"github.com/aws/constructs-go/constructs/v10"
)

type ReceiveRequestDoActionIds interface {
	Construct() *string
	ApplicationLoadBalancedFargateService() *string
	DistributeImage() repository.DistributeImageIds
}

type ReceiveRequestDoActionProps interface {
	ApplicationLoadBalancedFargateService() *awsecspatterns.ApplicationLoadBalancedFargateServiceProps
	DistributeImage() repository.DistributeImageProps
	Tag() *string
	// connections
	AddContainerImageToApplicationLoadBalancedFargate(awsecs.ContainerImage)
}

type ReceiveRequestDoAction interface {
	constructs.Construct
	ApplicationLoadBalancedFargateService() awsecspatterns.ApplicationLoadBalancedFargateService
	DistributeImage() repository.DistributeImage
}

func NewReceiveRequestDoAction(scope constructs.Construct, id ReceiveRequestDoActionIds, props ReceiveRequestDoActionProps) ReceiveRequestDoAction {
	var sprops ReceiveRequestDoActionProps = &ReceiveRequestDoActionModelProps_DEFAULT
	var sid ReceiveRequestDoActionIds = &ReceiveRequestDoActionModelIds_DEFAULT

	if props != nil {
		sprops = props
	}

	if id != nil {
		sid = id
	}

	repo := repository.NewDistributeImage(scope, sid.DistributeImage(), sprops.DistributeImage())

	this := constructs.NewConstruct(scope, sid.Construct())

	image := awsecs.AssetImage_FromEcrRepository(repo.Repository(), sprops.Tag())

	sprops.AddContainerImageToApplicationLoadBalancedFargate(image)

	resource := awsecspatterns.NewApplicationLoadBalancedFargateService(this, sid.ApplicationLoadBalancedFargateService(), sprops.ApplicationLoadBalancedFargateService())

	var component ReceiveRequestDoAction = &ReceiveRequestDoActionModel{
		Construct:                             this,
		applicationloadbalancedfargateservice: resource,
		distributeimage:                       repo,
	}

	return component
}

package applicationloadbalancedfargateservice

import (
	"castor/construct/highlevel/repository"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecspatterns"
	"github.com/aws/jsii-runtime-go"
)

type ReceiveEchoWriteEchoIds struct {
	ReceiveRequestDoActionModelIds
}

type ReceiveEchoWriteEchoProps struct {
	ReceiveRequestDoActionModelProps
}

// SETTINGS
// DEVELOPMENT
var ReceiveEchoWriteEchoIds_DEV ReceiveEchoWriteEchoIds = ReceiveEchoWriteEchoIds{
	ReceiveRequestDoActionModelIds: ReceiveRequestDoActionModelIds{
		ConstructId:                             jsii.String("ReceiveEchoWriteEcho-dev"),
		ApplicationLoadBalancedFargateServiceId: jsii.String("applicationloadbalancedfargateservice-dev"),
		DistributeImageIds:                      &repository.DistributeGoServerImageIds_DEV,
	},
}

var ReceiveEchoWriteEchoProps_DEV ReceiveEchoWriteEchoProps = ReceiveEchoWriteEchoProps{
	ReceiveRequestDoActionModelProps: ReceiveRequestDoActionModelProps{
		ApplicationLoadBalancedFargateServiceProps: &awsecspatterns.ApplicationLoadBalancedFargateServiceProps{
			MemoryLimitMiB:   jsii.Number(1024),
			DesiredCount:     jsii.Number(1),
			Cpu:              jsii.Number(512),
			LoadBalancerName: jsii.String("echoapp-alb-dev"),
		},
		DistributeImageProps: &repository.DistributeGoServerImageProps_DEV,
		TagImage:             jsii.String("latest"),
	},
}

// PRODUCTION
var ReceiveEchoWriteEchoIds_PROD ReceiveEchoWriteEchoIds = ReceiveEchoWriteEchoIds{
	ReceiveRequestDoActionModelIds: ReceiveRequestDoActionModelIds{
		ConstructId:                             jsii.String("ReceiveEchoWriteEcho-prod"),
		ApplicationLoadBalancedFargateServiceId: jsii.String("applicationloadbalancedfargateservice-prod"),
		DistributeImageIds:                      &repository.DistributeGoServerImageIds_PROD,
	},
}

var ReceiveEchoWriteEchoProps_PROD ReceiveEchoWriteEchoProps = ReceiveEchoWriteEchoProps{
	ReceiveRequestDoActionModelProps: ReceiveRequestDoActionModelProps{
		ApplicationLoadBalancedFargateServiceProps: &awsecspatterns.ApplicationLoadBalancedFargateServiceProps{
			MemoryLimitMiB:   jsii.Number(1024),
			DesiredCount:     jsii.Number(1),
			Cpu:              jsii.Number(512),
			LoadBalancerName: jsii.String("echoapp-alb-prod"),
		},
		DistributeImageProps: &repository.DistributeGoServerImageProps_PROD,
		TagImage:             jsii.String("latest"),
	},
}

package environment

import "github.com/aws/aws-cdk-go/awscdk/v2"

var StageProps_DEV awscdk.StageProps = awscdk.StageProps{
	Env: StackProps_DEV.Env,
}

var StageProps_PROD awscdk.StageProps = awscdk.StageProps{
	Env: StackProps_PROD.Env,
}

package alfa

import (
	computesave "castor/stack/computesave/alfa"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

type StackCollectionIds interface {
	Stage() *string
	Funcionality() computesave.FuncionalityIds
}

type StackCollectionProps interface {
	Stage() *awscdk.StageProps
	Funcionality() computesave.FuncionalityProps
}

type StackCollection interface {
	Stage() awscdk.Stage
	Funcionality() computesave.Funcionality
}

func NewStackCollection(scope constructs.Construct, id StackCollectionIds, props StackCollectionProps) StackCollection {
	var sprops StackCollectionProps = &StackCollectionModelProps_DEFAULT
	var sid StackCollectionIds = &StackCollectionModelIds_DEFAULT

	if props != nil {
		sprops = props
	}

	if id != nil {
		sid = id
	}

	stage := awscdk.NewStage(scope, sid.Stage(), sprops.Stage())

	funcionality := computesave.NewFuncionality(stage, sid.Funcionality(), sprops.Funcionality())

	var component StackCollection = &StackCollectionModel{
		stage:       stage,
		computesave: funcionality,
	}

	return component
}

package beta

import (
	computesave "castor/stack/computesave/beta"
	"castor/stack/environment"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

type StackCollectionModelIds struct {
	StageId        *string
	FuncionalityId computesave.FuncionalityIds
}

func (id *StackCollectionModelIds) Stage() *string {
	return id.StageId
}

func (id *StackCollectionModelIds) Funcionality() computesave.FuncionalityIds {
	return id.FuncionalityId
}

type StackCollectionModelProps struct {
	*awscdk.StageProps
	computesave.FuncionalityProps
}

func (props *StackCollectionModelProps) Stage() *awscdk.StageProps {
	return props.StageProps
}

func (props StackCollectionModelProps) Funcionality() computesave.FuncionalityProps {
	return props.FuncionalityProps
}

type StackCollectionModel struct {
	stage       awscdk.Stage
	computesave computesave.Funcionality
}

func (mo StackCollectionModel) Stage() awscdk.Stage {
	return mo.stage
}

func (mo StackCollectionModel) Funcionality() computesave.Funcionality {
	return mo.computesave
}

// SETTINGS
var StackCollectionModelIds_DEFAULT StackCollectionModelIds = StackCollectionModelIds{
	StageId:        jsii.String("StackColletionModelStage-default"),
	FuncionalityId: &computesave.FuncionalityModelIds_DEFAULT,
}

var StackCollectionModelProps_DEFAULT StackCollectionModelProps = StackCollectionModelProps{
	StageProps:        &environment.StageProps_DEFAULT,
	FuncionalityProps: &computesave.FuncionalityModelProps_DEFAULT,
}

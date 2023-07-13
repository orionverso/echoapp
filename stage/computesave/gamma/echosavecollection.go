package gamma

import (
	computesave "castor/stack/computesave/gamma"

	"github.com/aws/jsii-runtime-go"
)

type EchoSaveCollectionIds struct {
	StackCollectionModelIds
}

type EchoSaveCollectionProps struct {
	StackCollectionModelProps
}

// SETTINGS
// DEVELOPMENT
var EchoSaveCollectionIds_DEV EchoSaveCollectionIds = EchoSaveCollectionIds{
	StackCollectionModelIds: StackCollectionModelIds{
		StageId:        jsii.String("EchoSaveCollection-dev"),
		FuncionalityId: &computesave.EchoSaveIds_DEV,
	},
}

var EchoSaveCollectionProps_DEV EchoSaveCollectionProps = EchoSaveCollectionProps{
	StackCollectionModelProps: StackCollectionModelProps{
		FuncionalityProps: &computesave.EchoSaveProps_DEV,
	},
}

// PRODUCTION
var EchoSaveCollectionIds_PROD EchoSaveCollectionIds = EchoSaveCollectionIds{
	StackCollectionModelIds: StackCollectionModelIds{
		StageId:        jsii.String("EchoSaveCollection-prod"),
		FuncionalityId: &computesave.EchoSaveIds_PROD,
	},
}

var EchoSaveCollectionProps_PROD EchoSaveCollectionProps = EchoSaveCollectionProps{
	StackCollectionModelProps: StackCollectionModelProps{
		FuncionalityProps: &computesave.EchoSaveProps_PROD,
	},
}

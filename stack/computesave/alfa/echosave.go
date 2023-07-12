package alfa

import (
	"castor/construct/pattern/bucket"
	"castor/construct/pattern/lambdarestapi"
	"castor/stack/environment"

	"github.com/aws/jsii-runtime-go"
)

type EchoSaveIds struct {
	FuncionalityModelIds
}

type EchoSaveProps struct {
	FuncionalityModelProps
}

// SETTINGS
// DEVELOPMENT
var EchoSaveIds_DEV EchoSaveIds = EchoSaveIds{
	FuncionalityModelIds: FuncionalityModelIds{
		StackId:                   jsii.String("EchoSave-DEV"),
		ReceiveRequestDoActionIds: &lambdarestapi.ReceiveEchoWriteEchoIds_DEV,
		SaveFileIds:               &bucket.SaveEchoIds_DEV,
	},
}

var EchoSaveProps_DEV EchoSaveProps = EchoSaveProps{
	FuncionalityModelProps: FuncionalityModelProps{
		StackProps:                  environment.StackProps_DEV,
		ReceiveRequestDoActionProps: &lambdarestapi.ReceiveEchoWriteEchoProps_DEV,
		SaveFileProps:               &bucket.SaveEchoProps_DEV,
	},
}

// PRODUCTION
var EchoSaveIds_PROD EchoSaveIds = EchoSaveIds{
	FuncionalityModelIds: FuncionalityModelIds{
		StackId:                   jsii.String("EchoSave-PROD"),
		ReceiveRequestDoActionIds: &lambdarestapi.ReceiveEchoWriteEchoIds_PROD,
		SaveFileIds:               &bucket.SaveEchoIds_PROD,
	},
}

var EchoSaveProps_PROD EchoSaveProps = EchoSaveProps{
	FuncionalityModelProps: FuncionalityModelProps{
		StackProps:                  environment.StackProps_PROD,
		ReceiveRequestDoActionProps: &lambdarestapi.ReceiveEchoWriteEchoProps_PROD,
		SaveFileProps:               &bucket.SaveEchoProps_PROD,
	},
}

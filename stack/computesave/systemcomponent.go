package computesave

import "github.com/aws/jsii-runtime-go"

// We are going to state what should be considered the same piece of the system.
// To force them to change one for another.
type (
	compute  *string
	save     *string
	auxiliar *string
	state    *string
)

var (
	server     compute = jsii.String("Compute")
	serverless compute = jsii.String("Compute")
)

var (
	object   save = jsii.String("Save")
	database save = jsii.String("Save")
)

var (
	stateless state = jsii.String("Stateless")
	stateful  state = jsii.String("Stateful")
)

var fargaterepo auxiliar = jsii.String("FargateRepo")

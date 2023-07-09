package repository

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type DistributeModelIds struct {
	ConstructId  *string
	RepositoryId *string
}

func (id *DistributeModelIds) Construct() *string {
	return id.ConstructId
}

func (id *DistributeModelIds) Repository() *string {
	return id.RepositoryId
}

type DistributeModelProps struct {
	*awsecr.RepositoryProps
}

func (props *DistributeModelProps) Repository() *awsecr.RepositoryProps {
	return props.RepositoryProps
}

type DistributeModel struct {
	constructs.Construct
	repository awsecr.Repository
}

func (mo *DistributeModel) Repository() awsecr.Repository {
	return mo.repository
}

// SETTINGS
var DistributeModelIds_DEFAULT DistributeModelIds = DistributeModelIds{
	ConstructId:  jsii.String("MODEL-resource-distributemodel-default"),
	RepositoryId: jsii.String("MODEL-resource-repository-default"),
}

var DistributeModelProps_DEFAULT DistributeModelProps = DistributeModelProps{
	RepositoryProps: &awsecr.RepositoryProps{
		AutoDeleteImages: jsii.Bool(true),
		Encryption:       awsecr.RepositoryEncryption_KMS(),
		RemovalPolicy:    awscdk.RemovalPolicy_DESTROY,
		RepositoryName:   jsii.String("go-server-images-default"),
	},
}

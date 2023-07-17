package repository

import (
	"log"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type EcrRepoProps struct {
	awsecr.RepositoryProps
	Tag string
}

type ecrRepo struct {
	constructs.Construct
	repository awsecr.Repository
	image      awsecs.EcrImage
}

type EcrRepo interface {
	Repository() awsecr.Repository
	Image() awsecs.EcrImage
}

func NewEcrRepo(scope constructs.Construct, id *string, props *EcrRepoProps) EcrRepo {
	var sprops *EcrRepoProps = &EcrRepoProps{}

	if id == nil {
		log.Panicln("parameter id is required, but nil was provided")
	}

	if props != nil {
		sprops = props
	}

	this := constructs.NewConstruct(scope, id)

	repo := awsecr.NewRepository(this, jsii.String("Repository"), &sprops.RepositoryProps)

	image := awsecs.EcrImage_FromEcrRepository(repo, &sprops.Tag)

	var component EcrRepo = &ecrRepo{
		Construct:  this,
		repository: repo,
		image:      image,
	}

	return component
}

// IMPLEMENTATION
func (mo *ecrRepo) Repository() awsecr.Repository {
	return mo.repository
}

func (mo *ecrRepo) Image() awsecs.EcrImage {
	return mo.image
}

// SETTINGS
// DEVELOPMENT
var EcrRepoProps_DEV EcrRepoProps = EcrRepoProps{
	RepositoryProps: awsecr.RepositoryProps{
		RepositoryName: jsii.String("echoapp-dev"),
	},
	Tag: "latest",
}

// PRODUCTION
var EcrRepoProps_PROD EcrRepoProps = EcrRepoProps{
	RepositoryProps: awsecr.RepositoryProps{
		RepositoryName: jsii.String("echoapp-prod"),
	},
	Tag: "latest",
}

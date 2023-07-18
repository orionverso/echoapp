package repository

import (
	"log"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type EcrRepoProps struct {
	awsecr.RepositoryProps
	Tag string
	awsiam.PolicyStatementProps
}

type ecrRepo struct {
	constructs.Construct
	repository awsecr.Repository
	image      awsecs.EcrImage
	sts        awsiam.PolicyStatement
}

type EcrRepo interface {
	Repository() awsecr.Repository
	Image() awsecs.EcrImage
	PushImagePolicyStatement() awsiam.PolicyStatement
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

	sprops.AddResourceToPolicyStatement(repo.RepositoryArn())

	sts := awsiam.NewPolicyStatement(&sprops.PolicyStatementProps)

	var component EcrRepo = &ecrRepo{
		Construct:  this,
		repository: repo,
		image:      image,
		sts:        sts,
	}

	return component
}

// PROPS
func (props *EcrRepoProps) AddResourceToPolicyStatement(arn *string) {
	var resource *[]*string = &[]*string{}
	*resource = append(*resource, arn)
	props.Resources = resource
}

// IMPLEMENTATION
func (mo *ecrRepo) Repository() awsecr.Repository {
	return mo.repository
}

func (mo *ecrRepo) Image() awsecs.EcrImage {
	return mo.image
}

func (mo *ecrRepo) PushImagePolicyStatement() awsiam.PolicyStatement {
	return mo.sts
}

// SETTINGS
// DEVELOPMENT
var EcrRepoProps_DEV EcrRepoProps = EcrRepoProps{
	RepositoryProps: awsecr.RepositoryProps{
		RepositoryName:   jsii.String("echoapp-dev"),
		RemovalPolicy:    awscdk.RemovalPolicy_DESTROY,
		AutoDeleteImages: jsii.Bool(true),
	},
	Tag: "latest",

	PolicyStatementProps: awsiam.PolicyStatementProps{
		Actions: jsii.Strings("ecr:PutImageTagMutability",
			"ecr:PutReplicationConfiguration",
			"ecr:PutImage"),
		Principals: &[]awsiam.IPrincipal{
			awsiam.NewServicePrincipal(jsii.String("codebuild.amazonaws.com"), &awsiam.ServicePrincipalOpts{}),
		},
	},
}

// PRODUCTION
var EcrRepoProps_PROD EcrRepoProps = EcrRepoProps{
	RepositoryProps: awsecr.RepositoryProps{
		RepositoryName:   jsii.String("echoapp-prod"),
		RemovalPolicy:    awscdk.RemovalPolicy_DESTROY,
		AutoDeleteImages: jsii.Bool(true),
	},
	Tag: "latest",
}

package receiver

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

type getReceiverProps struct {
	stgGetParameterInput ssm.GetParameterInput
	dstGetParameterInput ssm.GetParameterInput
}

type Receiver interface {
	Write(ctx context.Context, st string) error
}

func GetReceiverFromParameter(ctx context.Context, cfg aws.Config) Receiver {
	var sprops getReceiverProps = getReceiverProps_DEFAULT

	ssmclient := ssm.NewFromConfig(cfg)
	stg, err := ssmclient.GetParameter(ctx, &sprops.stgGetParameterInput)
	if err != nil {
		log.Panicln("We do not know the service solution")
	}

	storage := aws.ToString(stg.Parameter.Value)
	log.Println("Storage:", storage)

	dst, err := ssmclient.GetParameter(ctx, &sprops.dstGetParameterInput)
	if err != nil {
		log.Panicln("We do not know the destination to write")
	}

	destination := aws.ToString(dst.Parameter.Value)
	log.Println("Destination:", destination)

	if storage == "DYNAMODB" {
		return dynamoDbReceiver{
			*dynamodb.NewFromConfig(cfg),
			destination,
		}
	}

	if storage == "S3" {
		return s3Receiver{
			*s3.NewFromConfig(cfg),
			destination,
		}
	}

	return nil
}

// CONFIGURATIONS

var getReceiverProps_DEFAULT getReceiverProps = getReceiverProps{
	stgGetParameterInput: ssm.GetParameterInput{
		Name: aws.String("STORAGE_SERVICE"),
	},
	dstGetParameterInput: ssm.GetParameterInput{
		Name: aws.String("DESTINATION"),
	},
}

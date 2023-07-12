package main

import (
	"castor/asset/receiver"
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
)

var (
	rec receiver.Receiver
	ctx context.Context
)

func handler(ctx context.Context, ev events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	err := rec.Write(ctx, ev.Body)
	if err != nil {
		log.Println(err)
		return events.APIGatewayProxyResponse{}, err
	}
	return events.APIGatewayProxyResponse{
		IsBase64Encoded:   false,
		StatusCode:        200,
		Headers:           map[string]string{},
		MultiValueHeaders: map[string][]string{},
		Body:              "Thank you for take a look. I am from Lambda. See you",
	}, nil
}

func main() {
	lambda.Start(handler)
}

func init() {
	ctx = context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Panicln("Configuration has not been loaded")
	}
	rec, err = receiver.GetReceiver(ctx, cfg)
	if err != nil {
		log.Panicln("Receive has not been loaded")
	}
}

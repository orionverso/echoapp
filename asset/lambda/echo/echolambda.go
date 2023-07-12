package main

import (
	"castor/asset/receiver"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
)

func handler(ctx context.Context, ev events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Panicln("Configuration has not been loaded")
	}

	rec := receiver.GetReceiver(ctx, cfg) // Change service destination if change env vars
	if err != nil {
		log.Panicln("Receive has not been loaded")
	}

	err = rec.Write(ctx, ev.Body)
	if err != nil {
		log.Println(err)
		return events.APIGatewayProxyResponse{}, err
	}
	return events.APIGatewayProxyResponse{
		IsBase64Encoded:   false,
		StatusCode:        200,
		Headers:           map[string]string{},
		MultiValueHeaders: map[string][]string{},
		Body: fmt.Sprint("Thank you for take a look. I am from Lambda writing to ",
			os.Getenv("STORAGE_SERVICE"), ". See you"),
	}, nil
}

func main() {
	lambda.Start(handler)
}

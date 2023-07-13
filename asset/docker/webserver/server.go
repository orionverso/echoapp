package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"server/receiver"

	"github.com/aws/aws-sdk-go-v2/config"
)

var (
	rec receiver.Receiver
	ctx context.Context
)

func main() {
	http.HandleFunc("/", postHandler)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Println("Server in port 80 has not been working properly")
	}

	err = http.ListenAndServe(":443", nil)
	if err != nil {
		log.Println("Server in port 443 has not been working properly")
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	ctx = context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx, LoadDefaultConfigProps_BINARY_IN_LOCAL.profile, LoadDefaultConfigProps_BINARY_IN_LOCAL.region)
	if err != nil {
		log.Panicln("Config has not been loaded properly")
	}

	rec = receiver.GetReceiverFromEnv(ctx, cfg)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error Body Read", http.StatusInternalServerError)
		return
	}

	if string(body) != "" {
		err = rec.Write(ctx, string(body))
		if err != nil {
			log.Println("Data has not been delivered properly")
		}
	}

	fmt.Fprintf(w, fmt.Sprint("Thanks you for take a look. I am from a Container writing to ", os.Getenv("STORAGE_SERVICE"), ". See you"))
}

type LoadDefaultConfigProps struct {
	profile config.LoadOptionsFunc
	region  config.LoadOptionsFunc
}

// CONFIGURATION
var LoadDefaultConfigProps_IMAGE_IN_LOCAL LoadDefaultConfigProps = LoadDefaultConfigProps{
	profile: config.WithSharedConfigProfile("cdk-role"),
	region:  config.WithRegion("us-east-1"),
}

var LoadDefaultConfigProps_BINARY_IN_LOCAL LoadDefaultConfigProps = LoadDefaultConfigProps{
	profile: config.WithSharedConfigProfile("workerdev"),
	region:  config.WithRegion("us-east-1"),
}

var LoadDefaultConfigProps_IN_CLOUD LoadDefaultConfigProps = LoadDefaultConfigProps{ // Asumme Credentials from Task Role
	profile: nil,
	region:  nil,
}

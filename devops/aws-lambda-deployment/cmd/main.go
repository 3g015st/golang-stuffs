package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
)

type EventBridgeMessage struct {
	Schedule string
}

func handler(ctx context.Context, event events.CloudWatchEvent) error {
	var msg EventBridgeMessage
	if err := json.Unmarshal(event.Detail, &msg); err != nil {
		fmt.Println("Unmarshal error", err.Error())
	}

	fmt.Printf("Received Event Bridge Message: %+v\n", msg.Schedule)

	return nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	lambdaServer := os.Getenv("_LAMBDA_SERVER_PORT")
	
	appName := os.Getenv("APP_NAME")
    fmt.Println("AWS Lambda Deployment Sample!")
	fmt.Println("App Name : ", appName)
	fmt.Println("Lamda Server : ", lambdaServer)

	lambda.Start(handler)
}
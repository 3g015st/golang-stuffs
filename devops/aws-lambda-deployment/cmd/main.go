package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	lambdaServer := os.Getenv("_LAMBDA_SERVER_PORT")
	
	appName := os.Getenv("APP_NAME")
    fmt.Println("AWS Lambda Deployment Sample")
	fmt.Println("App Name : ", appName)
	fmt.Println("Lamda Server : ", lambdaServer)
}
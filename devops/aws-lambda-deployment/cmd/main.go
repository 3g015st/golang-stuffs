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
	
	appName := os.Getenv("APP_NAME")
    fmt.Println("AWS Lambda Deployment Sample!")
	fmt.Println("App Name : ", appName)
}
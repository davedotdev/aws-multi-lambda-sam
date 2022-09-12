package main

import (
	"fmt"

	"messaging/common"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler() error {
	fmt.Println("hello world")
	return nil
}

func main() {
	clients := []common.ClientData{}
	fmt.Println(clients)
	lambda.Start(handler)
}
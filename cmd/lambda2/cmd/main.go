package main

import (
	"fmt"

	"messaging/common"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler() error {
	fmt.Println("hello world from lambda2")
	return nil
}

func main() {
	people := []common.Person{}
	fmt.Println(people)
	lambda.Start(handler)
}

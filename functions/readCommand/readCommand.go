package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(iotRequest interface{}) error {
	return nil
}

func main() {
	lambda.Start(Handler)
}

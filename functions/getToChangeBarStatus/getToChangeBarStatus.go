package main

import (
	"bartenderAsFunction/dao"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"strings"
)

var DataConnectionManager dao.CommandConnectionInterface

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	newState := "CLOSED"
	if strings.Contains(request.Body,"OPEN") {
		newState = "OPEN"
	}
	dao.UpdateBarShadow("serveuse-with-shadow",newState)

	return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(newState)}, nil
}


func main() {
	DataConnectionManager = dao.CreateCommandConnection()
	lambda.Start(Handler)
}

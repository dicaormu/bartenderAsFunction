package main

import (
	"bartenderAsFunction/dao"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var IotConnectionManager dao.IotConnectionInterface

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	idClient := request.PathParameters["idClient"]
	if idClient != "" {
		errChangeStatus := IotConnectionManager.UpdateShadow(idClient, "CLOSED")
		return events.APIGatewayProxyResponse{StatusCode: 200}, errChangeStatus
	}
	return events.APIGatewayProxyResponse{StatusCode: 400}, nil
}

func main() {
	IotConnectionManager = dao.CreateIotConnection()
	lambda.Start(Handler)
}

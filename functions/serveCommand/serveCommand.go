package main

import (
	"github.com/aws/aws-lambda-go/events"
	"encoding/json"
	"bartenderAsFunction/dao"
	"bartenderAsFunction/model"
	"github.com/aws/aws-lambda-go/lambda"
)

var DataConnectionManager dao.CommandConnectionInterface

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	idCommand := request.PathParameters["idCommand"]
	typeItem := request.PathParameters["type"]
	toServe := model.Item{}
	json.Unmarshal([]byte(request.Body), &toServe)

	// TODO 1. read the command
	command := DataConnectionManager.GetCommandById(idCommand)
	// TODO 2. Verify command exist
	if command.IdCommand == "" {
		return events.APIGatewayProxyResponse{StatusCode: 200, Body: "not available command to serve"}, nil
	}
	// TODO 3. search item
	if typeItem == "beer" {
		serveCommand(&command.Beer, toServe.Name)
	} else {
		serveCommand(&command.Food, toServe.Name)
	}
	// TODO 4. save command
	DataConnectionManager.SaveCommand(command)
	body, _ := json.Marshal(command)
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(body)}, nil
}

// TODO 3
func serveCommand(items *[]model.Item, name string) {
	for i, item := range *items {
		if item.Name == name {
			(*items)[i].Served = true
		}
	}
}

func main() {
	DataConnectionManager = dao.CreateCommandConnection()
	lambda.Start(Handler)
}

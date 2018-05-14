package main

import (
	"github.com/aws/aws-lambda-go/events"
	"encoding/json"
	"bartenderAsFunction/dao"
	"bartenderAsFunction/model"
	"github.com/aws/aws-lambda-go/lambda"
	"fmt"
)

var DataConnectionManager dao.CommandConnectionInterface

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	idCommand := request.PathParameters["idCommand"]
	typeItem := request.PathParameters["type"]
	toServe := model.Item{}
	json.Unmarshal([]byte(request.Body), &toServe)
	fmt.Println(idCommand)

	// TODO 1. read the command by idCommand. Hint, there are a dao package with all you need. Use the DataConnectionManager var
	var command model.Command

	// TODO 2. Verify command exist. If not, return 200 but with body "not available command to serve"
	// TODO 3. search item (just implement method in TODO 3
	if typeItem == "beer" {
		serveCommand(&command.Beer, toServe.Name)
	} else {
		serveCommand(&command.Food, toServe.Name)
	}
	// TODO 4. save command. User dao package

	// And return 200 with command. Use Json marshall to transform the command in []byte
	return events.APIGatewayProxyResponse{}, nil
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

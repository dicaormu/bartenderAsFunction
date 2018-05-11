package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"bartenderAsFunction/dao"
	"bartenderAsFunction/model"
	"encoding/json"
)

var DataConnectionManager dao.CommandConnectionInterface

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// TODO 1. read all unserved commands
	commands, err := DataConnectionManager.GetCommands()
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 404, Body: err.Error()}, nil
	}
	items := model.Command{}
	for _, command := range commands {
		items.Beer = getNoServedItemsForCommand(command.Beer)
		items.Food = getNoServedItemsForCommand(command.Food)
		if len(items.Beer) > 0 || len(items.Food) > 0 {
			items.IdCommand = command.IdCommand
		}
	}
	// TODO 3. return unserved commands
	body, _ := json.Marshal(items)
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(body)}, nil
}

// TODO 2. complete function to return no served items
func getNoServedItemsForCommand(items []model.Item) (noServedItems []model.Item) {
	for _, item := range items {
		if item.Served == false {
			noServedItems = append(noServedItems, item)
		}
	}
	return noServedItems
}

func main() {
	DataConnectionManager = dao.CreateCommandConnection()
	lambda.Start(Handler)
}

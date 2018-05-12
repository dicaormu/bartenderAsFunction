package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"bartenderAsFunction/dao"
	"bartenderAsFunction/model"
	"encoding/json"
	"fmt"
)

var DataConnectionManager dao.CommandConnectionInterface

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// TODO 1. read all unserved commands
	commands, err := DataConnectionManager.GetCommands()
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 404, Body: err.Error()}, nil
	}
	var commandsReturn []model.Command
	// TODO 2 iterate over items to get non served commands
	for _, command := range commands {
		items := model.Command{}
		items.Beer = getNoServedItemsForCommand(command.Beer)
		items.Food = getNoServedItemsForCommand(command.Food)
		if len(items.Beer) > 0 || len(items.Food) > 0 {
			items.IdCommand = command.IdCommand
			commandsReturn = append(commandsReturn, items)
		}
	}
	// TODO 3. return unserved commands
	body, _ := json.Marshal(commandsReturn)
	fmt.Println("unserved commands return ", string(body))
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

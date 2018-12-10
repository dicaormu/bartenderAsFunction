package main

import (
	"bartenderAsFunction/dao"
	"bartenderAsFunction/model"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var DataConnectionManager dao.CommandConnectionInterface

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// TODO 1. read all unserved commands
	commands, err := DataConnectionManager.GetCommands()
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 404, Body: err.Error()}, nil
	}
	// TODO 2 iterate over items to get non served commands
	commandsReturn := getNoServedCommands(commands)
	// TODO 3. return unserved commands
	body, _ := json.Marshal(commandsReturn)
	fmt.Println("unserved commands return ", string(body))
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(body)}, nil
}

// TODO 2. complete function to return no served items
func getNoServedCommands(items []model.Command) ([]model.Command) {
	noServedItems := []model.Command{}
	for _, item := range items {
		if (!item.Beer.Served && item.Beer.Amount > 0) || (!item.Food.Served && item.Food.Amount > 0) {
			noServedItems = append(noServedItems, item)
		}
	}
	return noServedItems
}

func main() {
	DataConnectionManager = dao.CreateCommandConnection()
	lambda.Start(Handler)
}

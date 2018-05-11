package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"encoding/json"
	"bartenderAsFunction/dao"
	"bartenderAsFunction/model"
)

var DataConnectionManager dao.CommandConnectionInterface

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// TODO 1. read all served commands
	commands, err := DataConnectionManager.GetCommands()
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 404, Body: err.Error()}, nil
	}
	// TODO 2. get items by type
	beer := make(map[string]int)
	food := make(map[string]int)
	for _, command := range commands {
		getItemsForType(command.Beer, &beer)
		getItemsForType(command.Food, &food)
	}
	// TODO 3. return unserved commands
	items := model.CommandRequest{}
	items.Beer = getItemsFromMap(beer)
	items.Food = getItemsFromMap(food)

	body, _ := json.Marshal(items)
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(body)}, nil
}


// TODO 3
func getItemsFromMap(itemMap map[string]int) (items []model.Item) {
	for key, value := range itemMap {
		items = append(items, model.Item{Name: key, Amount: value, Served:true})
	}
	return items
}

// TODO 2
func getItemsForType(items []model.Item, classifiedItems *map[string]int) {
	for _, item := range items {
		if item.Served == true {
			amt := (*classifiedItems)[item.Name]
			(*classifiedItems)[item.Name] = amt + item.Amount
		}
	}
}

func main() {
	DataConnectionManager = dao.CreateCommandConnection()
	lambda.Start(Handler)
}

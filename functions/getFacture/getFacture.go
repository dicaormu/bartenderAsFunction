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
	// TODO 1. read all served commands. Hint, there are a dao package with all you need. Use the DataConnectionManager var
	var commands []model.Command

	// TODO and validate error, return 404 and the error message in the body

	// TODO 2. get items by type (just implement the associated methods- see TODO 2)
	beer := make(map[string]int)
	food := make(map[string]int)
	for _, command := range commands {
		getItemsForType(command.Beer, &beer)
		getItemsForType(command.Food, &food)
	}
	// TODO 3. return unserved commands (just implement the associated methods- see TODO 3)
	items := model.CommandRequest{}
	items.Beer = getItemsFromMap(beer)
	items.Food = getItemsFromMap(food)

	// TODO Hint, use json marshall function to transform struct to []byte

	//TODO return a status 200 with body
	return
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

// TODO 3
func getItemsFromMap(itemMap map[string]int) (items []model.Item) {
	for key, value := range itemMap {
		items = append(items, model.Item{Name: key, Amount: value, Served:true})
	}
	return items
}

func main() {
	DataConnectionManager = dao.CreateCommandConnection()
	lambda.Start(Handler)
}

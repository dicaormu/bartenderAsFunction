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
	// TODO 1. read all unserved commands: Hint, there are a dao package with all you need. Use the DataConnectionManager var
	var commands []model.Command
	// TODO And validate error. If error return a code 404 and the error messagge in the body

	//
	var commandsReturn []model.Command
	// TODO 2 iterate over items to get non served commands
	for _, command := range commands {
		items := model.Command{}
		items.Beer = getNoServedItemsForCommand(command.Beer)
		items.Food = getNoServedItemsForCommand(command.Food)
		//TODO append command to result (commandsReturn). Dont forget the idcommand in the variable items

	}
	// TODO 3. return unserved commands. Hint, use json marshall function to transform struct to []byte


	//TODO return a status 200 with body
	return
}

// TODO 2. complete function to return no served items
func getNoServedItemsForCommand(items []model.Item) (noServedItems []model.Item) {
	for _, item := range items {
		// TODO append no served item
	}
	return noServedItems
}

func main() {
	DataConnectionManager = dao.CreateCommandConnection()
	lambda.Start(Handler)
}

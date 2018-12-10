package main

import (
	"bartenderAsFunction/dao"
	"bartenderAsFunction/model"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

var DataConnectionManager dao.CommandConnectionInterface

func Handler(request model.IotEvent) (string, error)  {
	fmt.Println("reported status " + request.Current.State.Reported.BarStatus)
	fmt.Println("desired status " + request.Current.State.Desired.BarStatus)
	if request.Current.State.Reported.BarStatus == request.Current.State.Desired.BarStatus && request.Current.State.Desired.BarStatus =="CLOSED" {
		// TODO 1. read all served commands
		commands, err := DataConnectionManager.GetCommands()
		commandsResponse := []model.Command{}
		if err != nil {
			return "", err
		}
		// TODO 2. get items and serve items
		for _, command := range commands {
			command.Food.Served=true
			command.Beer.Served=true
			DataConnectionManager.SaveCommand(command)
			commandsResponse = append(commandsResponse,command)
		}
		// TODO 3. return unserved commands
		body, _ := json.Marshal(commandsResponse)
		return string(body), nil
	}
	return "",nil
}

func main() {
	DataConnectionManager = dao.CreateCommandConnection()
	lambda.Start(Handler)
}

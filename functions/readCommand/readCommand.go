package main

import (
	"bartenderAsFunction/dao"
	"bartenderAsFunction/model"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/satori/go.uuid"
	"strings"
	"time"
)

var DataConnectionManager dao.CommandConnectionInterface

func Handler(request model.IotBeerCommandFromIot) error {
	predictions, _ := request.UnmarshalAssociatedData()
	for _, v := range predictions {
		// TODO 2. generate command (model.command) with date in utc format
		command := generateCommand(v)
		if command.IdCommand != "" {
			// TODO 3. save command in dynamo
			lastCommand, _ := DataConnectionManager.GetLastCommand()

			return saveCommandIfNew(command, lastCommand)
		}
	}
	return nil
}

func saveCommandIfNew(newCommand model.Command,lastCommand *model.Command) error {
	saveCommand := shouldSaveCommand(lastCommand, newCommand)
	if saveCommand {
		saveCommandError := DataConnectionManager.SaveCommand(newCommand)
		if saveCommandError != nil {
			return saveCommandError
		}
	}
	return nil
}

func shouldSaveCommand(lastCommand *model.Command, newCommand model.Command) bool {
	saveCommand := true
	if lastCommand != nil && (len(lastCommand.Beer) > 0 || len(lastCommand.Food) > 0) &&
		len(lastCommand.Beer) == len(newCommand.Beer) && len(lastCommand.Food) == len(newCommand.Food) {
		lastDate, _ := time.Parse(time.RFC3339, lastCommand.DateCommand)
		fmt.Println("lastDate",lastDate)
		newDate, _ := time.Parse(time.RFC3339, newCommand.DateCommand)
		fmt.Println("newDate",newDate)

		if lastDate.Add(time.Second * 60).After(newDate) {
			saveCommand = false
		}
	}
	return saveCommand
}

func generateCommand(prediction model.BeerPrediction) model.Command {
	var command model.Command
	// TODO 1. generate id to the command (uuid)
	uid, _ := uuid.NewV4()
	if strings.Contains(strings.ToLower(prediction.Label), "beer") {
		// TODO 2. generate command (model.command) with date in utc format
		command = model.Command{IdCommand: uid.String(), DateCommand: time.Now().UTC().Format(time.RFC3339), Beer: []model.Item{{prediction.Label, 1, false}}}
	}
	if strings.Contains(strings.ToLower(prediction.Label), "water") {
		// TODO 2. generate command (model.command) with date in utc format
		command = model.Command{IdCommand: uid.String(), DateCommand: time.Now().UTC().Format(time.RFC3339), Food: []model.Item{{prediction.Label, 1, false}}}
	}
	return command
}

func main() {
	DataConnectionManager = dao.CreateCommandConnection()
	lambda.Start(Handler)
}

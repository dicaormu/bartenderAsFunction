package main

import (
	"bartenderAsFunction/dao"
	"bartenderAsFunction/model"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/satori/go.uuid"
	"strings"
	"time"
)

var DataConnectionManager dao.CommandConnectionInterface

func Handler(iotRequest model.IotBeerCommandFromIot) error {
	predictions, _ := iotRequest.UnmarshalAssociatedData()
	for _, v := range predictions {
		// TODO 2. generate command (model.command) with date in utc format
		command := generateCommand(v)
		if command.IdCommand != "" {
			// TODO 3. save command in dynamo
			saveCommandError := DataConnectionManager.SaveCommand(command)
			if saveCommandError != nil {
				return saveCommandError
			}
		}
	}
	return nil
}

func generateCommand(prediction model.BeerPrediction) model.Command {
	var command model.Command
	// TODO 1. generate id to the command (uuid)
	uid, _ := uuid.NewV4()
	if strings.Contains(strings.ToLower(prediction.Label), "beer") {
		// TODO 2. generate command (model.command) with date in utc format
		command = model.Command{IdCommand: uid.String(), DateCommand: time.Now().UTC().Format(time.RFC3339), Beer: []model.Item{{prediction.Label, 1, false}}}
	}
	if strings.Contains(strings.ToLower(prediction.Label), "coke") {
		// TODO 2. generate command (model.command) with date in utc format
		command = model.Command{IdCommand: uid.String(), DateCommand: time.Now().UTC().Format(time.RFC3339), Food: []model.Item{{prediction.Label, 1, false}}}
	}

	return command
}

func main() {
	DataConnectionManager = dao.CreateCommandConnection()
	lambda.Start(Handler)
}

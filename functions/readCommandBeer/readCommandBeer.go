package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"bartenderAsFunction/model"
	"time"
	"bartenderAsFunction/dao"
	"github.com/satori/go.uuid"
	"fmt"
)

var DataConnectionManager dao.CommandConnectionInterface

func Handler(iotRequest model.CommandRequest) error {
	// TODO 1. generate id to the command (uuid)
	uid, _ := uuid.NewV4()
	fmt.Println("beer:", iotRequest.Beer)
	// TODO 2. generate command (model.command) with date in utc format
	command := model.Command{IdCommand: uid.String(), DateCommand: time.Now().UTC().Format(time.RFC3339), Beer: iotRequest.Beer}
	// TODO 3. save command in dynamo
	//  -> verify if there is not command in the last 2 minutes
	commands, err := DataConnectionManager.GetCommandsByClient(command.Client)
	if err != nil {
		return err
	}
	if shouldSaveCommand(commands, time.Now()) {
		saveCommandError := DataConnectionManager.SaveCommand(command)
		if saveCommandError != nil {
			return saveCommandError
		}
	}
	return nil
}

func shouldSaveCommand(commands []model.Command, actualDate time.Time) bool {
	for _, val := range commands {
		if val.Beer.Amount >0  {
				dateCommand, _ := time.Parse(time.RFC3339, val.DateCommand)
				fmt.Println("dateCommand", dateCommand)
				fmt.Println("actualDate", actualDate)
				// 2 - minutes to hold
				if dateCommand.Add(time.Second * 60 * 2).After(actualDate) {
					return false
			}
		}
	}
	return true
}

func main() {
	DataConnectionManager = dao.CreateCommandConnection()
	lambda.Start(Handler)
}

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
	fmt.Println("food:",iotRequest.Food)
	// TODO 2. generate command (model.command) with date in utc format
	command := model.Command{IdCommand: uid.String(), DateCommand: time.Now().UTC().Format(time.RFC3339), Food: iotRequest.Food}
	// TODO 3. save command in dynamo
	saveCommandError := DataConnectionManager.SaveCommand(command)
	if saveCommandError != nil {
		return saveCommandError
	}
	return nil
}

func main() {
	DataConnectionManager = dao.CreateCommandConnection()
	lambda.Start(Handler)
}

package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"bartenderAsFunction/model"
	"time"
	"bartenderAsFunction/dao"
	"github.com/satori/go.uuid"
)

var DataConnectionManager dao.CommandConnectionInterface

func Handler(iotRequest model.CommandRequest) error {
	// TODO 1. generate id to the command (uuid)
	uid, _ := "",""
	// TODO 2. generate command (model.command) with date in utc format
	command := model.Command{}
	// TODO 3. save command in dynamo. Hint, there are a dao package with all you need. Use the DataConnectionManager var

	// TODO return the error if it exist

	//or return nil if everything is ok
	return nil
}

func main() {
	DataConnectionManager = dao.CreateCommandConnection()
	lambda.Start(Handler)
}

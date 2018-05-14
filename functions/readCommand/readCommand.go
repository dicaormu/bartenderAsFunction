package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"bartenderAsFunction/model"
	"bartenderAsFunction/dao"
	"fmt"
)

var DataConnectionManager dao.CommandConnectionInterface

func Handler(iotRequest model.CommandRequest) error {
	// TODO 1. generate id to the command (uuid) see github.com/satori/go.uuid
	uid, _ := "", ""
	fmt.Println(uid)
	// TODO 2. generate command (model.command) with date in utc format
	command := model.Command{}
	fmt.Println(command)
	// TODO 3. save command in dynamo. Hint, there are a dao package with all you need. Use the DataConnectionManager var

	// TODO return the error if it exist

	//or return nil if everything is ok
	return nil
}

func main() {
	DataConnectionManager = dao.CreateCommandConnection()
	lambda.Start(Handler)
}

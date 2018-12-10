package dao

import (
	"github.com/aws/aws-sdk-go/aws"
	"os"
	"strings"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"bartenderAsFunction/model"
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type CommandConnection struct {
	DynamoConnection *dynamodb.DynamoDB
}

type CommandConnectionInterface interface {
	SaveCommand(command model.Command) error
	GetCommands() ([]model.Command, error)
	GetCommandsByClient(idClient string) ([]model.Command, error)
	GetCommandById(id string) model.Command
}

func (con *CommandConnection) GetCommandsByClient(idClient string) ([]model.Command, error) {
	tableName := os.Getenv("TABLE_COMMANDS")
	params := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}
	result, err := con.DynamoConnection.Scan(params)
	if err != nil {
		fmt.Println("error getting dynamo", err)
		return nil, err
	}
	var commands []model.Command
	for _, v := range result.Items {
		var item model.Command
		err = dynamodbattribute.UnmarshalMap(v, &item)
		if item.Client == idClient {
			commands = append(commands, item)
			fmt.Printf("getting %d commands", len(commands))
		}
	}
	return commands, nil
}

func (con *CommandConnection) GetCommandById(id string) model.Command {
	tableName := os.Getenv("TABLE_COMMANDS")
	result, err := con.DynamoConnection.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})
	if err != nil {
		fmt.Println(err)
		return model.Command{}
	}
	command := model.Command{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &command)

	if err != nil {
		fmt.Println(err)
		return model.Command{}
	}

	if command.IdCommand == "" {
		return model.Command{}
	}
	return command
}

func (con *CommandConnection) SaveCommand(command model.Command) error {
	tableName := os.Getenv("TABLE_COMMANDS")
	av, err := dynamodbattribute.MarshalMap(command)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// Create item in table
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}
	_, err = con.DynamoConnection.PutItem(input)
	if err != nil {
		fmt.Println("error saving", err)
		return err
	}
	return nil
}

func (con *CommandConnection) GetCommands() ([]model.Command, error) {
	tableName := os.Getenv("TABLE_COMMANDS")
	params := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}
	result, err := con.DynamoConnection.Scan(params)
	if err != nil {
		fmt.Println("error getting dynamo", err)
		return nil, err
	}
	var commands []model.Command
	for _, v := range result.Items {
		var item model.Command
		err = dynamodbattribute.UnmarshalMap(v, &item)
		commands = append(commands, item)
		fmt.Printf("getting %d commands", len(commands))
	}
	return commands, nil
}

func CreateCommandConnection() CommandConnectionInterface {
	return &CommandConnection{initializeDynamoDBClient()}
}

func initializeDynamoDBClient() *dynamodb.DynamoDB {
	localEnv := os.Getenv("AWS_SAM_LOCAL")
	dynamoUrl := os.Getenv("dynamoUrl")
	awsConfig := &aws.Config{Region: aws.String(os.Getenv("AWS_DEFAULT_REGION"))}
	if len(localEnv) > 0 && strings.ToLower(localEnv) == "true" {
		if dynamoUrl == "" {
			dynamoUrl = "http://docker.for.mac.localhost:8000"
		}
		awsConfig.Endpoint = aws.String(dynamoUrl)
	}
	sessionVar, err := session.NewSession(awsConfig)
	if err != nil {
		fmt.Println("error connecting", err)
		os.Exit(1)
	}
	return dynamodb.New(sessionVar, aws.NewConfig().WithLogLevel(aws.LogDebugWithHTTPBody))
}

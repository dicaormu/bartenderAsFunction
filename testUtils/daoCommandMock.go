package testUtils

import (
	"github.com/stretchr/testify/mock"
	"bartenderAsFunction/model"
	"reflect"
	"errors"
	"time"
	"fmt"
	"encoding/json"
)

type CommandConnectionMock struct {
	mock.Mock
	Command       model.Command
	ExpectedError error
}

func isDate(date string) bool {
	_, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return false
	}
	return true
}

func (service *CommandConnectionMock) SaveCommand(command model.Command) error {
	bytes, _ := json.Marshal(command)
	fmt.Println(string(bytes))

	if service.ExpectedError != nil {
		return service.ExpectedError
	}
	if command.DateCommand == "" || ! isDate(command.DateCommand) {
		return errors.New("error, no date in format RFC3339")
	}
	if command.IdCommand == "" {
		return errors.New("error, no id")
	}
	if reflect.DeepEqual(service.Command.Food, command.Food) && reflect.DeepEqual(service.Command.Beer, command.Beer) {
		return nil
	}
	return errors.New("no matching expected and saved command")
}

func (service *CommandConnectionMock) GetCommands() ([]model.Command, error) {
	var aaa []model.Command
	aaa = append(aaa, service.Command)
	return aaa, service.ExpectedError
}

func (service *CommandConnectionMock) GetCommandsByClient(idClient string) ([]model.Command, error) {
	var aaa []model.Command
	aaa = append(aaa, service.Command)
	return aaa, service.ExpectedError
}

func (service *CommandConnectionMock) GetCommandById(id string) model.Command {
	fmt.Println("id:::::",id)
	fmt.Println("service id:::::",service.Command.IdCommand )
	if service.Command.IdCommand == id {
		return service.Command
	}
	return model.Command{}
}

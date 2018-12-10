package main

import (
	"bartenderAsFunction/model"
	"bartenderAsFunction/testUtils"
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlerShouldReturn404(t *testing.T) {
	mock := testUtils.CommandConnectionMock{ExpectedError: errors.New("error")}
	DataConnectionManager = &mock
	response, err := Handler(model.IotEvent{})
	assert.Equal(t, response, "")
	assert.Equal(t, err, nil)
}

func TestHandlerShouldReturnBillOneItemBeer(t *testing.T) {
	beerItems := model.Item{Name: "leffe", Served: false, Amount: 1}
	mock := testUtils.CommandConnectionMock{Command: model.Command{IdCommand: "111", Beer: beerItems}}
	DataConnectionManager = &mock

	response, err := Handler(model.IotEvent{Current:model.IotShadowDoc{State:model.IotShadowState{Desired:model.ClientObjectState{"CLOSED"}, Reported:model.ClientObjectState{"CLOSED"}}}})
	resultCommand := []model.Command{}

	fmt.Println(response)
	assert.Nil(t, err)
	json.Unmarshal([]byte(response), &resultCommand)
	assert.Len(t, resultCommand, 1)
	assert.Equal(t, resultCommand[0].Food.Served, true)
	assert.Equal(t, resultCommand[0].Food.Amount, 0)
	assert.Equal(t, resultCommand[0].Beer.Served, true)
	assert.Equal(t, resultCommand[0].Beer.Name, beerItems.Name)
	assert.Equal(t, resultCommand[0].Beer.Amount, beerItems.Amount)
}

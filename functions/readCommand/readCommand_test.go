package main

import (
	"bartenderAsFunction/model"
	"testing"
	"errors"
	"bartenderAsFunction/testUtils"
	"github.com/stretchr/testify/assert"
	"encoding/json"
)

func TestHandlerShouldReturnError(t *testing.T) {
	mock := testUtils.CommandConnectionMock{ExpectedError: errors.New("error")}
	DataConnectionManager = &mock
	iotRequest := model.CommandRequest{}
	err := Handler(iotRequest)
	assert.Equal(t, err.Error(), "error")
}

func TestHandlerShouldReturnErrorNotValidDate(t *testing.T) {
	commandIot := `{"food": [{"item":"pizza","amount":2}], "beer": [{"item":"1664","amount":3}]}`
	iotRequest := model.CommandRequest{}

	json.Unmarshal([]byte(commandIot),&iotRequest)

	beerItems := []model.Item{{Served:false,Name:"1664",Amount:3}}
	foodItems := []model.Item{{Served:false,Name:"pizza",Amount:2}}//
	mock := testUtils.CommandConnectionMock{Command:model.Command{Beer: beerItems,Food:foodItems}}
	DataConnectionManager = &mock
	err := Handler(iotRequest)
	assert.Equal(t, err, nil)
}

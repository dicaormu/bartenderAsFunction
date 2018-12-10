package main

import (
	"bartenderAsFunction/model"
	"bartenderAsFunction/testUtils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandlerShouldReturnError(t *testing.T) {
	mock := testUtils.CommandConnectionMock{ExpectedError: errors.New("error")}
	DataConnectionManager = &mock
	iotRequest := model.CommandRequest{}
	err := Handler(iotRequest)
	assert.Equal(t, err.Error(), "error")
}

func TestHandlerShouldReturnErrorNotValidDate(t *testing.T) {
	commandIot := `{"food": {"item":"pizza","amount":2}}`
	iotRequest := model.CommandRequest{}

	json.Unmarshal([]byte(commandIot), &iotRequest)

	fmt.Println("iotrequest =======",iotRequest)

	foodItems := model.Item{Served: false, Name: "pizza", Amount: 2} //
	mock := testUtils.CommandConnectionMock{Command: model.Command{Food: foodItems}}
	DataConnectionManager = &mock
	err := Handler(iotRequest)
	assert.Equal(t, err, nil)
}

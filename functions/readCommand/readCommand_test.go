package main

import (
	"bartenderAsFunction/model"
	"testing"
	"errors"
	"bartenderAsFunction/testUtils"
	"github.com/stretchr/testify/assert"
)

func TestHandlerShouldReturnError(t *testing.T) {
	mock := testUtils.CommandConnectionMock{ExpectedError: errors.New("error")}
	DataConnectionManager = &mock
	iotRequest := model.CommandRequest{}
	err := Handler(iotRequest)
	assert.Equal(t, err.Error(), "error")
}

func TestHandlerShouldReturnErrorNotValidDate(t *testing.T) {
	beerItems := []model.Item{{Served:false,Name:"aaa",Amount:3}}
	foodItems := []model.Item{{Served:false,Name:"bbb",Amount:2}}
	mock := testUtils.CommandConnectionMock{Command:model.Command{Beer: beerItems,Food:foodItems}}
	DataConnectionManager = &mock
	iotRequest := model.CommandRequest{Food:foodItems,Beer:beerItems}
	err := Handler(iotRequest)
	assert.Equal(t, err, nil)
}

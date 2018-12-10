package main

import (
	"bartenderAsFunction/model"
	"bartenderAsFunction/testUtils"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandlerShouldReturnNoCommand(t *testing.T) {
	mock := testUtils.IotConnectionMock{DrunkClient: model.DrunkClient{IdClient: "111"}}
	IotConnectionManager = &mock
	event := events.APIGatewayProxyRequest{PathParameters: map[string]string{"idClient": "1111"}}
	response, _ := Handler(event)
	assert.Equal(t, response.Body, "")
	assert.Equal(t, response.StatusCode, 200)
}

func TestHandlerShouldReturnCommand(t *testing.T) {
	item := model.Item{Name: "1664", Amount: 1, Served: false}
	mock := testUtils.IotConnectionMock{DrunkClient: model.DrunkClient{IdClient: "111"}}
	IotConnectionManager = &mock
	event := events.APIGatewayProxyRequest{}
	response, _ := Handler(event)
	item.Served = true

	assert.Equal(t, response.Body, "")
	assert.Equal(t, response.StatusCode, 400)
}

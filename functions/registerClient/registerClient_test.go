package main

import (
	"bartenderAsFunction/testUtils"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandlerShouldReturnError(t *testing.T) {
	mock := testUtils.IotConnectionMock{ExpectedError: errors.New("error")}
	IotConnectionManager = &mock
	iotRequest :=events.APIGatewayProxyRequest{Body:`{"id":"123"}`}
	res,err := Handler(iotRequest)
	assert.Nil(t, err)
	assert.Equal(t, res.StatusCode, 200)
}



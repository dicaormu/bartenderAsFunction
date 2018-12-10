package main

import (
	"bartenderAsFunction/model"
	"bartenderAsFunction/testUtils"
	"errors"
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestHandlerShouldReturn404(t *testing.T) {
	mock := testUtils.CommandConnectionMock{ExpectedError: errors.New("error")}
	DataConnectionManager = &mock
	response, _ := Handler(events.APIGatewayProxyRequest{})
	assert.Equal(t, response.StatusCode, 404)
	assert.Equal(t, response.Body, "error")
}

func Test_getNoServedItemsForCommand(t *testing.T) {
	type args struct {
		items []model.Command
	}
	tests := []struct {
		name              string
		args              args
		wantNoServedItems []model.Command
	}{
		{"no items", args{[]model.Command{}}, []model.Command{}},
		{"1 item served", args{[]model.Command{{Beer:model.Item{Amount: 1, Served: true}}}},[]model.Command{}},
		{"1 item no served", args{[]model.Command{{Beer:model.Item{Amount: 1, Served: true, Name: "item"}}, {Food:model.Item{Amount: 1, Served: true, Name: "item"}}}},[]model.Command{}},
		{"2 items, served and no served", args{[]model.Command{{Beer:model.Item{Amount: 2, Served: true, Name: "served item"}}, {Food:model.Item{Amount: 1, Served: false, Name: "item"}}}}, []model.Command{{Food:model.Item{Amount: 1, Served: false, Name: "item"}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNoServedItems := getNoServedCommands(tt.args.items); !reflect.DeepEqual(gotNoServedItems, tt.wantNoServedItems) {
				t.Errorf("%s getNoServedItemsForCommand() = %v, want %v", tt.name, gotNoServedItems, tt.wantNoServedItems)
			}
		})
	}
}
/*
func TestHandlerShouldReturnOneBeer(t *testing.T) {
	beerItems := []model.Item{{Name: "leffe", Served: false, Amount: 1}, {Name: "1664", Served: true, Amount: 1}}
	mock := testUtils.CommandConnectionMock{Command: model.Command{IdCommand: "111", Beer: beerItems}}
	DataConnectionManager = &mock
	response, _ := Handler(events.APIGatewayProxyRequest{})
	assert.Equal(t, response.StatusCode, 200)
	var resultCommand []model.Command
	json.Unmarshal([]byte(response.Body), &resultCommand)
	assert.Len(t, resultCommand, 1)
	assert.Equal(t, resultCommand[0].IdCommand, "111")
	assert.Len(t, resultCommand[0].Food, 0)
	assert.Len(t, resultCommand[0].Beer, 1)
	assert.Equal(t, resultCommand[0].Beer[0].Name, "leffe")
}

func TestHandlerShouldReturnOneFood(t *testing.T) {
	foodItems := []model.Item{{Name: "pizza", Served: false, Amount: 1}, {Name: "Burger", Served: true, Amount: 1}}
	mock := testUtils.CommandConnectionMock{Command: model.Command{IdCommand: "222", Food: foodItems}}
	DataConnectionManager = &mock
	response, _ := Handler(events.APIGatewayProxyRequest{})
	assert.Equal(t, response.StatusCode, 200)
	var resultCommand []model.Command
	json.Unmarshal([]byte(response.Body), &resultCommand)
	assert.Len(t, resultCommand, 1)
	assert.Equal(t, resultCommand[0].IdCommand, "222")
	assert.Len(t, resultCommand[0].Beer, 0)
	assert.Len(t, resultCommand[0].Food, 1)
	assert.Equal(t, resultCommand[0].Food[0].Name, "pizza")
}
*/
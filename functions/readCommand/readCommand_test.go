package main

import (
	"bartenderAsFunction/model"
	"bartenderAsFunction/testUtils"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHandlerShouldReturnError(t *testing.T) {
	mock := testUtils.CommandConnectionMock{ExpectedError: errors.New("error")}
	DataConnectionManager = &mock
	iotRequest := model.IotBeerCommandFromIot{}
	err := Handler(iotRequest)
	assert.Nil(t, err, "error")
}

func TestHandlerShouldReturnErrorNotValidDate(t *testing.T) {
	commandIot := `{"food": [{"item":"pizza","amount":2}], "beer": [{"item":"1664","amount":3}]}`
	iotRequest := model.IotBeerCommandFromIot{}

	json.Unmarshal([]byte(commandIot), &iotRequest)

	beerItems := []model.Item{{Served: false, Name: "1664", Amount: 3}}
	foodItems := []model.Item{{Served: false, Name: "pizza", Amount: 2}} //
	mock := testUtils.CommandConnectionMock{Command: model.Command{Beer: beerItems, Food: foodItems}}
	DataConnectionManager = &mock
	err := Handler(iotRequest)
	assert.Equal(t, err, nil)
}


func createCommandTestBeer(hour, min,sec int)model.Command {
	location := time.UTC
	return model.Command{IdCommand:"1",Beer:[]model.Item{{Amount:1}},DateCommand: time.Date(2018, 11, 15, hour, min, sec, 0, location).UTC().Format(time.RFC3339)}
}

func createCommandTestFood(hour, min,sec int)model.Command {
	location := time.UTC
	return model.Command{IdCommand:"1",Food:[]model.Item{{Amount:1}},DateCommand: time.Date(2018, 11, 15, hour, min, sec, 0, location).UTC().Format(time.RFC3339)}
}

func Test_shouldSaveCommand(t *testing.T) {
	type args struct {
		lastCommand model.Command
		newCommand  model.Command
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name:"same date",args:args{createCommandTestBeer(10,10,10), createCommandTestBeer(10,10,10)},want:false},
		{name:"lower by some seconds",args:args{createCommandTestBeer(10,10,0), createCommandTestBeer(10,10,10)},want:false},
		{name:"lower by 59 seconds",args:args{createCommandTestBeer(10,10,0), createCommandTestBeer(10,10,59)},want:false},
		{name:"greater by 62 seconds",args:args{createCommandTestBeer(10,10,0), createCommandTestBeer(10,11,1)},want:true},
		{name:"greater by 120 seconds",args:args{createCommandTestBeer(10,10,0), createCommandTestBeer(10,12,0)},want:true},
		{name:"with last food and new beer",args:args{createCommandTestFood(10,10,0), createCommandTestBeer(10,12,0)},want:true},
		{name:"with last food and new beer lower by some seconds",args:args{createCommandTestFood(10,12,0), createCommandTestBeer(10,12,10)},want:true},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shouldSaveCommand(&tt.args.lastCommand, tt.args.newCommand); got != tt.want {
				t.Errorf("shouldSaveCommand()  %s =  %v, want %v",tt.name, got, tt.want)
			}
		})
	}
}

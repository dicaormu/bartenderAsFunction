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
	iotRequest := model.CommandRequest{}
	err := Handler(iotRequest)
	assert.Equal(t, err.Error(), "error")
}

func TestHandlerShouldReturnErrorNotValidDate(t *testing.T) {
	commandIot := `{"beer": {"item":"1664","amount":3}}`
	iotRequest := model.CommandRequest{}

	json.Unmarshal([]byte(commandIot), &iotRequest)

	beerItems := model.Item{Served: false, Name: "1664", Amount: 3}
	mock := testUtils.CommandConnectionMock{Command: model.Command{Beer: beerItems}}
	DataConnectionManager = &mock
	err := Handler(iotRequest)
	assert.Equal(t, err, nil)
}

func Test_shouldSaveCommand(t *testing.T) {
	type args struct {
		commands   []model.Command
		actualDate time.Time
	}
	actualDate := time.Date(2018, 12, 21, 14, 0, 0, 0, time.UTC)
	actualDatePlus1 := time.Date(2018, 12, 21, 14, 1, 0, 0, time.UTC)
	commands := []model.Command{{IdCommand: "1", DateCommand: actualDate.Format(time.RFC3339), Beer: model.Item{Served: false, Name: "1664", Amount: 3}},
		{IdCommand: "2", DateCommand: actualDatePlus1.Format(time.RFC3339), Beer: model.Item{Served: false, Name: "1664", Amount: 3}},}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"0 minutes", args{commands, actualDatePlus1}, false},
		{"0 minutes plus 1", args{commands, actualDate}, false},
		{"i min 59 secs", args{commands, time.Date(2018, 12, 21, 14, 1, 59, 0, time.UTC)}, false},
		{"121 secs", args{commands, time.Date(2018, 12, 21, 14, 3, 1, 0, time.UTC)}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shouldSaveCommand(tt.args.commands, tt.args.actualDate); got != tt.want {
				t.Errorf("%s shouldSaveCommand() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

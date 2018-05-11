package main

import (
	"bartenderAsFunction/model"
	"reflect"
	"testing"
	"bartenderAsFunction/testUtils"
	"github.com/stretchr/testify/assert"
	"github.com/aws/aws-lambda-go/events"
	"encoding/json"
)

func Test_serveCommand(t *testing.T) {
	type args struct {
		items   *[]model.Item
		toServe string
	}
	tests := []struct {
		name      string
		args      args
		wantItems []model.Item
	}{
		{"no items", args{&[]model.Item{}, "aaa"}, []model.Item{}},
		{"1 item, not equal", args{&[]model.Item{{Amount: 1, Name: "bbb", Served: false}}, "aaa"}, []model.Item{{Amount: 1, Name: "bbb", Served: false}}},
		{"1 item, equal", args{&[]model.Item{{Amount: 1, Name: "bbb", Served: false}}, "bbb"}, []model.Item{{Amount: 1, Name: "bbb", Served: true}}},
		{"1 item equal, 1 not equal", args{&[]model.Item{{Amount: 1, Name: "bbb", Served: false}, {Amount: 1, Name: "aaa", Served: false}}, "bbb"}, []model.Item{{Amount: 1, Name: "bbb", Served: true}, {Amount: 1, Name: "aaa", Served: false}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if serveCommand(tt.args.items, tt.args.toServe); !reflect.DeepEqual(*tt.args.items, tt.wantItems) {
				t.Errorf("%s getItemsFromMap() = %v, want %v", tt.name, tt.args.items, tt.wantItems)
			}
		})
	}
}

func TestHandlerShouldReturnNoCommand(t *testing.T) {
	mock := testUtils.CommandConnectionMock{Command: model.Command{IdCommand: "111"}}
	DataConnectionManager = &mock
	event := events.APIGatewayProxyRequest{PathParameters: map[string]string{"idCommand": "1", "type": "beer"}}
	response, _ := Handler(event)
	assert.Equal(t, response.Body, "not available command to serve")
	assert.Equal(t, response.StatusCode, 200)
}

func TestHandlerShouldReturnCommand(t *testing.T) {
	item := model.Item{Name: "1664", Amount: 1, Served: false}
	mock := testUtils.CommandConnectionMock{Command: model.Command{IdCommand: "111", Beer: []model.Item{item}}}
	DataConnectionManager = &mock
	body, _ := json.Marshal(item)
	event := events.APIGatewayProxyRequest{PathParameters: map[string]string{"idCommand": "111", "type": "beer"},
		Body: string(body),
	}
	response, _ := Handler(event)
	item.Served = true
	command := model.Command{IdCommand: "111", Beer: []model.Item{item}}
	resp, _ := json.Marshal(command)

	assert.Equal(t, response.Body, string(resp))
	assert.Equal(t, response.StatusCode, 200)
}

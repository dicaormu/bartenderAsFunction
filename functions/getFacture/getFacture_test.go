package main

import (
	"bartenderAsFunction/model"
	"bartenderAsFunction/testUtils"
	"errors"
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"encoding/json"
)

func TestHandlerShouldReturn404(t *testing.T) {
	mock := testUtils.CommandConnectionMock{ExpectedError: errors.New("error")}
	DataConnectionManager = &mock
	response, _ := Handler(events.APIGatewayProxyRequest{})
	assert.Equal(t, response.StatusCode, 404)
	assert.Equal(t, response.Body, "error")
}

func Test_getItemsForType(t *testing.T) {
	type args struct {
		items           []model.Item
		classifiedItems map[string]int
	}
	tests := []struct {
		name                string
		args                args
		wantClassifiedItems map[string]int
	}{
		{"nil items", args{nil, make(map[string]int)}, map[string]int{}},
		{"no items", args{[]model.Item{}, make(map[string]int)}, map[string]int{}},
		{"1 item no served", args{[]model.Item{{Amount: 1, Served: false}}, make(map[string]int)}, map[string]int{}},
		{"1 item served", args{[]model.Item{{Amount: 1, Served: true, Name: "a name"}}, make(map[string]int)}, map[string]int{"a name": 1}},
		{"1 item served, 1 no served", args{[]model.Item{{Amount: 1, Served: true, Name: "a name"}, {Amount: 1, Served: false, Name: "another name"}}, make(map[string]int)}, map[string]int{"a name": 1}},
		{"1 item served, existing", args{[]model.Item{{Amount: 1, Served: true, Name: "a name"}}, map[string]int{"a name": 2, "no name": 1}}, map[string]int{"a name": 3, "no name": 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if getItemsForType(tt.args.items, &tt.args.classifiedItems); !reflect.DeepEqual(tt.args.classifiedItems, tt.wantClassifiedItems) {
				t.Errorf("%s getItemsForType() = %v, want %v", tt.name, tt.args.classifiedItems, tt.wantClassifiedItems)
			}
		})
	}
}

func Test_getItemsFromMap(t *testing.T) {
	type args struct {
		itemMap map[string]int
	}
	tests := []struct {
		name      string
		args      args
		wantItems []model.Item
	}{
		{"no items", args{make(map[string]int)}, nil},
		{"nil items", args{nil}, nil},
		{"1 item", args{map[string]int{"a name": 3}}, []model.Item{{Amount: 3, Name: "a name", Served: true}}},
		{"2 items", args{map[string]int{"a name": 3, "another name": 1}}, []model.Item{{Amount: 3, Name: "a name", Served: true}, {Amount: 1, Name: "another name", Served: true}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotItems := getItemsFromMap(tt.args.itemMap); !reflect.DeepEqual(gotItems, tt.wantItems) {
				t.Errorf("%s getItemsFromMap() = %v, want %v", tt.name, gotItems, tt.wantItems)
			}
		})
	}
}

func TestHandlerShouldReturnBillOneItemBeer(t *testing.T) {
	beerItems := []model.Item{{Name: "leffe", Served: true, Amount: 1}}
	mock := testUtils.CommandConnectionMock{Command: model.Command{IdCommand: "111", Beer: beerItems}}
	DataConnectionManager = &mock

	response, _ := Handler(events.APIGatewayProxyRequest{})
	resultCommand := model.CommandRequest{}

	assert.Equal(t, response.StatusCode, 200)
	json.Unmarshal([]byte(response.Body), &resultCommand)
	assert.Len(t, resultCommand.Beer, 1)
	assert.Len(t, resultCommand.Food, 0)
	assert.Equal(t, resultCommand.Beer, beerItems)
}

func TestHandlerShouldReturnBillOneItemFood(t *testing.T) {
	foodItems := []model.Item{{Name: "pizza", Served: true, Amount: 1}}
	mock := testUtils.CommandConnectionMock{Command: model.Command{IdCommand: "111", Food: foodItems}}
	DataConnectionManager = &mock

	response, _ := Handler(events.APIGatewayProxyRequest{})
	resultCommand := model.CommandRequest{}

	assert.Equal(t, response.StatusCode, 200)
	json.Unmarshal([]byte(response.Body), &resultCommand)
	assert.Len(t, resultCommand.Food, 1)
	assert.Len(t, resultCommand.Beer, 0)
	assert.Equal(t, resultCommand.Food, foodItems)
}

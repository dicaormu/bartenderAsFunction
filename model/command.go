package model

import (
	"regexp"
	"strconv"
	"strings"
)

type Command struct {
	IdCommand   string `json:"id"`
	DateCommand string `json:"date"`
	Food        []Item `json:"food"`
	Beer        []Item `json:"beer"`
}

type Item struct {
	Name   string `json:"item"`
	Amount int    `json:"amount"`
	Served bool   `json:"served"`
}

type CommandRequest struct {
	Food []Item `json:"food"`
	Beer []Item `json:"beer"`
}

type IotBeerCommandFromIot struct {
	Format  string `json:"format"`
	Payload string `json:"payload"`
	Qos     string `json:"qos"`
}

type BeerPrediction struct {
	Confidence float64
	Label      string
}

func (cmd IotBeerCommandFromIot) UnmarshalAssociatedData() ([]BeerPrediction, error) {
	var prediction []BeerPrediction
	re := regexp.MustCompile(`^\[\((\d*.\d*), ('.*')\), \((\d*.\d*), ('.*')\), \((\d*.\d*), ('.*')\), \((\d*.\d*), ('.*')\), \((\d*.\d*), ('.*')\)\]`)
	submatch := re.FindStringSubmatch(cmd.Payload)
	for i := 0; i < 5; i++ {
		var unmarshalled BeerPrediction
		unmarshalled.Confidence, _ = strconv.ParseFloat(submatch[i*2+1], 64)
		unmarshalled.Label = strings.Replace(submatch[i*2+2],"'","",-1)
		prediction = append(prediction, unmarshalled)
	}
	return prediction, nil
}

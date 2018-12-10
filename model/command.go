package model

type Command struct {
	IdCommand   string `json:"id"`
	DateCommand string `json:"date"`
	Food        Item `json:"food"`
	Beer        Item `json:"beer"`
	Client      string `json:"client"`
}

type Item struct {
	Name   string `json:"item"`
	Amount int    `json:"amount"`
	Served bool   `json:"served"`
}

type CommandRequest struct {
	Food Item `json:"food"`
	Beer Item `json:"beer"`
}

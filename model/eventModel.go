package model

type Event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type AllEvents []Event

var Events = AllEvents{
	{
		ID:          "1",
		Title:       "Introduction to Golang",
		Description: "Learn Golang",
	},
}

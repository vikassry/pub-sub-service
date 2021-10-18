package model

type Event struct {
	Type string    `json:"type"`
	Data EventData `json:"data"`
}

type EventData struct {
	Id          int    `json:"id"`
	Message     string `json:"message"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Time        string `json:"time"`
}

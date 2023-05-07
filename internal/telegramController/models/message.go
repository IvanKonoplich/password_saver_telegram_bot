package models

type Message struct {
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}
type Chat struct {
	ChatID int `json:"id"`
}

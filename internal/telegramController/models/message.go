package models

type Message struct {
	MessageID int    `json:"message_Id"`
	Chat      Chat   `json:"chat"`
	Text      string `json:"text"`
}
type Chat struct {
	ChatID int `json:"id"`
}

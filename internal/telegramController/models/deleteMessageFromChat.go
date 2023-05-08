package models

type DeleteMessageQuery struct {
	ChatId    int `json:"chat_id"`
	MessageID int `json:"message_id"`
}

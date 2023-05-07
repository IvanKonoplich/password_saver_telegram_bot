package controller

import (
	"github.com/sirupsen/logrus"
	"password_storage_telegram/internal/entities"
)

type UseCase interface {
	Set(inc entities.IncomingData) error
	Get(inc entities.IncomingData) (string, error)
	Del(inc entities.IncomingData) error
}
type TgController struct {
	uc UseCase
}

func New(uc UseCase) *TgController {
	return &TgController{
		uc: uc,
	}
}
func (tgc *TgController) RunTgController(botToken string) {
	botURL := "https://api.telegram.org/bot" + botToken
	offset := 0
	for {
		updates, err := GetUpdates(botURL, offset)
		if err != nil {
			logrus.Errorf("smt went wrong: %s", err.Error())
		}
		for _, update := range updates {
			logrus.Infof("new message: %s, from: %d", update.Message.Text, update.Message.Chat.ChatID)
			err = tgc.Respond(botURL, update)
			offset = update.UpdateID + 1
			if err != nil {
				logrus.Errorf("smt went wrong: %s", err.Error())
			}
		}
	}

}

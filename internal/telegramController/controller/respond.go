package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"password_storage_telegram/internal/entities"
	"password_storage_telegram/internal/telegramController/models"
	"strings"
)

func (tgc *TgController) Respond(botURL string, update models.Update) error {
	//структура для ответа
	var botMessage models.BotMessage
	//тут дергаем юзкейс
	incoming := strings.Split(update.Message.Text, " ")
	command := incoming[0]
	var responseText string
	switch command {
	case "/set":
		if len(incoming) != 3 {
			responseText = "Неправильная команда. Введите команду, ресурс и пароль через пробел. Пример: /set ресурс.ru пароль12345"
			break
		}
		dts := entities.DataToSave{ResourceName: incoming[1], Password: incoming[2]}
		err := tgc.uc.Set(dts)
		if err != nil {
			responseText = err.Error()
		}
		responseText = "пароль успешно сохранен"
	case "/get":
		if len(incoming) != 2 {
			responseText = "Неправильная команда. Введите команду и ресурс через пробел. Пример: /get ресурс.ru"
			break
		}
		savedData, err := tgc.uc.Get(incoming[1])
		if err != nil {
			responseText = err.Error()
		}
		responseText = savedData
	case "/del":
		if len(incoming) != 2 {
			responseText = "Неправильная команда. Введите команду и ресурс через пробел. Пример: /del ресурс.ru"
			break
		}
		err := tgc.uc.Del(incoming[1])
		if err != nil {
			responseText = err.Error()
		}
		responseText = "пароль успешно удален"
	default:
		responseText = "неизвестная команда"
	}

	botMessage.ChatId = update.Message.Chat.ChatID
	botMessage.Text = responseText
	//формируем ответ
	buf, err := json.Marshal(botMessage)
	if err != nil {
		return err
	}
	//отправляем ответ
	_, err = http.Post(botURL+"/sendMessage", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	return nil
}

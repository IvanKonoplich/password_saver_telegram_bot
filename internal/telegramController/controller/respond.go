package controller

import (
	"bytes"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"password_storage_telegram/internal/entities"
	"password_storage_telegram/internal/telegramController/models"
	"strings"
	"time"
)

const helpMessage string = "Я знаю следующие команды: " +
	"\n/set для сохранения пароля для какого-нибудь ресурса. Пример: /set ресурс.ru пароль12345 " +
	"\n/get для получения сохраненного пароля для какого-нибудь ресурса. Пример: /get ресурс.ru" +
	"\n/del для удаления сохраненного пароля от какого-нибудь ресурса. Пример: /del ресурс.ru" +
	"\n/help для того, чтобы посмотреть инструкцию"
const helloMessage string = "Привет, я бот для сохранения паролей. Я знаю следующие команды: " +
	"\n/set для сохранения пароля для какого-нибудь ресурса. Пример: /set ресурс.ru пароль12345 " +
	"\n/get для получения сохраненного пароля для какого-нибудь ресурса. Пример: /get ресурс.ru" +
	"\n/del для удаления сохраненного пароля от какого-нибудь ресурса. Пример: /del ресурс.ru" +
	"\n/help для того, чтобы посмотреть инструкцию"

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
		incData := entities.IncomingData{ResourceName: incoming[1], Password: incoming[2], ChatID: update.Message.Chat.ChatID}
		err := tgc.uc.Set(incData)
		if err != nil {
			responseText = err.Error()
		} else {
			responseText = "Пароль успешно сохранен"
		}
		go func() {
			time.Sleep(time.Second * 3)
			var deleteQuery models.DeleteMessageQuery
			deleteQuery.ChatId = update.Message.Chat.ChatID
			deleteQuery.MessageID = update.Message.MessageID
			buf, err := json.Marshal(deleteQuery)
			_, err = http.Post(botURL+"/deleteMessage", "application/json", bytes.NewBuffer(buf))
			if err != nil {
				logrus.Infof("error while deleting message from chat: %s", err.Error())
			}
			botMessage.ChatId = update.Message.Chat.ChatID
			botMessage.Text = "Сообщение с паролем удалено из этого чата"
			//формируем ответ
			buf, err = json.Marshal(botMessage)
			if err != nil {
				logrus.Errorf("error while deleting message from chat: %s", err.Error())
			}
			//отправляем ответ
			_, err = http.Post(botURL+"/sendMessage", "application/json", bytes.NewBuffer(buf))
			if err != nil {
				logrus.Errorf("error while deleting message from chat: %s", err.Error())
			}
		}()
	case "/get":
		if len(incoming) != 2 {
			responseText = "Неправильная команда. Введите команду и ресурс через пробел. Пример: /get ресурс.ru"
			break
		}
		incData := entities.IncomingData{ResourceName: incoming[1], ChatID: update.Message.Chat.ChatID}
		savedData, err := tgc.uc.Get(incData)
		if err != nil {
			responseText = err.Error()
		} else {
			responseText = savedData
		}
		go func() {
			time.Sleep(time.Second * 5)
			var deleteQuery models.DeleteMessageQuery
			deleteQuery.ChatId = update.Message.Chat.ChatID
			deleteQuery.MessageID = update.Message.MessageID + 1
			buf, err := json.Marshal(deleteQuery)
			_, err = http.Post(botURL+"/deleteMessage", "application/json", bytes.NewBuffer(buf))
			if err != nil {
				logrus.Infof("error while deleting message from chat: %s", err.Error())
			}
			botMessage.ChatId = update.Message.Chat.ChatID
			botMessage.Text = "Сообщение с паролем удалено из этого чата"
			//формируем ответ
			buf, err = json.Marshal(botMessage)
			if err != nil {
				logrus.Errorf("error while deleting message from chat: %s", err.Error())
			}
			//отправляем ответ
			_, err = http.Post(botURL+"/sendMessage", "application/json", bytes.NewBuffer(buf))
			if err != nil {
				logrus.Errorf("error while deleting message from chat: %s", err.Error())
			}
		}()
	case "/del":
		if len(incoming) != 2 {
			responseText = "Неправильная команда. Введите команду и ресурс через пробел. Пример: /del ресурс.ru"
			break
		}
		incData := entities.IncomingData{ResourceName: incoming[1], ChatID: update.Message.Chat.ChatID}
		err := tgc.uc.Del(incData)
		if err != nil {
			responseText = err.Error()
		} else {
			responseText = "Пароль успешно удален"
		}
	case "/help":
		responseText = helpMessage
	case "/start":
		responseText = helloMessage
	default:
		responseText = "Неизвестная команда"
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

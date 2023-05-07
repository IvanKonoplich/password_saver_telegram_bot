package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"password_storage_telegram/internal/telegramController/models"
	"strconv"
)

func GetUpdates(botURL string, offset int) ([]models.Update, error) {
	resp, err := http.Get(botURL + "/getUpdates" + "?offset=" + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response models.Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return response.Updates, nil
}

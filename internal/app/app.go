package app

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
	"password_storage_telegram/internal/infrastructure/storage"
	"password_storage_telegram/internal/telegramController/controller"
	"password_storage_telegram/internal/usecase"
)

func RunApp() {
	botToken := initGoDotEnv()
	store := storage.New()
	uc := usecase.New(store)
	cont := controller.New(uc)
	cont.RunTgController(botToken)
}

func initGoDotEnv() string {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}

	return os.Getenv("bot_token")
}

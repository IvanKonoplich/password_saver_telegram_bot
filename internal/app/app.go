package app

import (
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"password_storage_telegram/internal/infrastructure/storage"
	"password_storage_telegram/internal/telegramController/controller"
	"password_storage_telegram/internal/usecase"
)

func RunApp(botToken string, postgresConfig storage.ConfigDB) {
	db, err := storage.OpenDBConnection(postgresConfig)
	if err != nil {
		logrus.Fatalf("error opening postgres connection:%s", err.Error())
	}
	store := storage.New(db)
	uc := usecase.New(store)
	cont := controller.New(uc)
	cont.RunTgController(botToken)
}

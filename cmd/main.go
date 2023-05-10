package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"password_storage_telegram/internal/app"
	"password_storage_telegram/internal/infrastructure/storage"
)

//docker run --name tg_password_bot -p 5432:5432 -e POSTGRES_PASSWORD=qwerty -d postgres
func main() {
	botToken := initGoDotEnv()
	postgresPassword := "qwerty"
	if err := initConfig(); err != nil {
		logrus.Fatalf("error while reading config file:%s", err.Error())
	}
	postgresConfig := storage.ConfigDB{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.DBName"),
		Password: postgresPassword,
		SSLMode:  viper.GetString("db.SSLMode"),
	}
	app.RunApp(botToken, postgresConfig)
}

func initConfig() error {
	viper.SetDefault("port", "8000")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("config")
	return viper.ReadInConfig()
}

func initGoDotEnv() string {
	err := godotenv.Load()
	if err != nil {
		logrus.Error("Error loading .env file")
	}

	return os.Getenv("BOT_TOKEN")
}

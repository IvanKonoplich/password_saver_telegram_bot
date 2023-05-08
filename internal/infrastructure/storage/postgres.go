package storage

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type ConfigDB struct {
	Host     string
	Port     string
	Username string
	DBName   string
	Password string
	SSLMode  string
}

func OpenDBConnection(cfg ConfigDB) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	logrus.Info("postgres connection opened successfully")
	schema := `CREATE TABLE IF NOT EXISTS passwords(
		chat_id int not null ,
		resource varchar not null,
		password varchar not null
	);`
	_, err = db.Exec(schema)
	if err != nil {
		return nil, err
	}
	return db, nil
}

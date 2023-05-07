package usecase

import "password_storage_telegram/internal/entities"

func (*UseCase) Set(inc entities.IncomingData) error {
	return nil
}
func (*UseCase) Get(inc entities.IncomingData) (string, error) {
	return "", nil
}
func (*UseCase) Del(inc entities.IncomingData) error {
	return nil
}

func check() bool {
	return false
}

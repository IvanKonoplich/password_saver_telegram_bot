package usecase

import "password_storage_telegram/internal/entities"

func (*UseCase) Set(inc entities.DataToSave) error {
	return nil
}
func (*UseCase) Get(ResourceName string) (string, error) {
	return "", nil
}
func (*UseCase) Del(ResourceName string) error {
	return nil
}

func check() bool {
	return false
}

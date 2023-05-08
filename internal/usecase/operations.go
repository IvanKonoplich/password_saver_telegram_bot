package usecase

import (
	"errors"
	"password_storage_telegram/internal/entities"
)

func (uc *UseCase) Set(inc entities.IncomingData) error {
	is, err := uc.check(inc)
	if err != nil {
		return err
	}
	if is {
		return errors.New("У вас уже сохранен пароль для этого ресурса. Вы можете удалить его используя /del и назначить заново через /set")
	}
	return uc.st.Set(inc)
}
func (uc *UseCase) Get(inc entities.IncomingData) (string, error) {
	is, err := uc.check(inc)
	if err != nil {
		return "", err
	}
	if !is {
		return "", errors.New("у вас нет сохраненного пароля для этого ресурса")
	}
	return uc.st.Get(inc)
}
func (uc *UseCase) Del(inc entities.IncomingData) error {
	is, err := uc.check(inc)
	if err != nil {
		return err
	}
	if !is {
		return errors.New("у вас нет сохраненного пароля для этого ресурса")
	}
	return uc.st.Del(inc)
}

func (uc *UseCase) check(inc entities.IncomingData) (bool, error) {
	is, err := uc.st.Get(inc)
	if err != nil {
		return false, err
	}
	if is != "" {
		return true, nil
	}
	return false, nil
}

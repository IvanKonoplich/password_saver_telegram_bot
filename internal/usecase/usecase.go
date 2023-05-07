package usecase

import "password_storage_telegram/internal/entities"

type Storage interface {
	Set(inc entities.IncomingData) error
	Get(inc entities.IncomingData) (string, error)
	Del(inc entities.IncomingData) error
}

type UseCase struct {
	st Storage
}

func New(st Storage) *UseCase {
	return &UseCase{
		st: st,
	}
}

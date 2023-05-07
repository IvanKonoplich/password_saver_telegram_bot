package storage

import "password_storage_telegram/internal/entities"

func (s *Storage) Set(inc entities.IncomingData) error {
	return nil
}
func (s *Storage) Get(inc entities.IncomingData) (string, error) {
	return "", nil
}
func (s *Storage) Del(inc entities.IncomingData) error {
	return nil
}

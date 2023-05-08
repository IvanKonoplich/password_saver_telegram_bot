package storage

import "password_storage_telegram/internal/entities"

func (s *Storage) Set(inc entities.IncomingData) error {
	query := `INSERT INTO passwords (chat_id, resource, password) values ($1, $2,$3)`
	_, err := s.db.Exec(query, inc.ChatID, inc.ResourceName, inc.Password)
	return err
}
func (s *Storage) Get(inc entities.IncomingData) (string, error) {
	var password string
	query := `SELECT password FROM passwords WHERE chat_id=$1`
	row := s.db.QueryRow(query, inc.ChatID)
	if err := row.Scan(&password); err != nil {
		return "", nil
	}
	return password, nil
}
func (s *Storage) Del(inc entities.IncomingData) error {
	query := `DELETE FROM passwords WHERE chat_id=$1 AND resource=$2`
	_, err := s.db.Exec(query, inc.ChatID, inc.ResourceName)
	return err
}

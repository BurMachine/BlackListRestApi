package database

import "context"

var getTokenShema = `SELECT id FROM tokens WHERE token=$1`
var addTokenShema = `INSERT INTO tokens(token) VALUES($1)`

func (s *Storage) CheckToken(token string) bool {
	var id int
	err := s.Conn.QueryRow(context.Background(), getTokenShema, token).Scan(&id)
	if err != nil {
		return false
	}
	return true
}

func (s *Storage) AddToken(token string) error {
	_, err := s.Conn.Exec(context.Background(), addTokenShema, token)
	if err != nil {
		return err
	}
	return nil
}

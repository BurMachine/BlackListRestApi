package database

import "context"

var removeSchema = `DELETE FROM blacklist WHERE id = $1;`

func (s *Storage) Remove(id string) error {
	_, err := s.Conn.Exec(context.Background(), removeSchema, id)
	if err != nil {
		return err
	}
	return nil
}

package database

import (
	"blacklistApi/internal/models"
	"context"
)

var addSchema = `INSERT INTO blacklist (phone_number, user_name, reason_for_adding, adding_date, admin_name) 
	VALUES ($1, $2, $3, $4, $5)`

func (s *Storage) Add(addiction models.Addiction) error {
	_, err := s.Conn.Exec(context.Background(), addSchema, addiction.UserPhone, addiction.UserName, addiction.Reason,
		addiction.Date, addiction.AdminName)
	if err != nil {
		return err
	}
	return nil
}

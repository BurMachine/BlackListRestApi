package database

import (
	"blacklistApi/internal/models"
	"context"
	"github.com/jackc/pgx/v5"
)

var searchSchemaPhone = `SELECT phone_number, user_name, reason_for_adding, adding_date, admin_name
        FROM blacklist WHERE phone_number=$1`

var searchSchemaName = `SELECT phone_number, user_name, reason_for_adding, adding_date, admin_name
        FROM blacklist WHERE user_name=$1`

func (s *Storage) Search(param string, value string) ([]models.Addiction, error) {
	var rows pgx.Rows
	var err error
	if param == "phone_number" {
		rows, err = s.Conn.Query(context.Background(), searchSchemaPhone, value)
		if err != nil {
			return nil, err
		}
	} else if param == "user_name" {
		rows, err = s.Conn.Query(context.Background(), searchSchemaName, value)
		if err != nil {
			return nil, err
		}
	}

	defer rows.Close()

	var users []models.Addiction
	for rows.Next() {
		var user models.Addiction
		err = rows.Scan(&user.UserPhone, &user.UserName, &user.Reason, &user.Date, &user.AdminName)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

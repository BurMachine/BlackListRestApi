package database

import (
	"blacklistApi/internal/config"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

type Storage struct {
	Conn *pgx.Conn
}

var schema = `CREATE TABLE IF NOT EXISTS blacklist (
	id integer PRIMARY KEY  NOT NULL,
	phone_number VARCHAR(20) NOT NULL,
	user_name  VARCHAR(50) NOT NULL,
	reason_for_adding TEXT NOT NULL,
	adding_date TIMESTAMP NOT NULL,
	admin_name  VARCHAR(50) NOT NULL
);`

func InitConn(conf config.Conf) (*Storage, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("POSTGRES_URI"))
	if err != nil {
		conn, err = pgx.Connect(context.Background(), conf.DbUrl)
		if err != nil {
			return &Storage{}, err
		}
	}

	_, err = conn.Exec(context.Background(), schema)
	if err != nil {
		return nil, fmt.Errorf("table creating error: %v", err)
	}
	return &Storage{conn}, nil
}

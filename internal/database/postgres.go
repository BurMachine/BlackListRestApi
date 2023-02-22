package database

import (
	"blacklistApi/internal/config"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

var schema = `CREATE TABLE IF NOT EXISTS links (
	original_link VARCHAR(1000) NOT NULL,
	short_link VARCHAR(10) UNIQUE NOT NULL
);`

func InitConn(conf config.Conf) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("POSTGRES_URI"))
	if err != nil {
		return nil, err
	}

	_, err = conn.Exec(context.Background(), schema)
	if err != nil {
		return nil, fmt.Errorf("table creating error: %v", err)
	}
	return conn, nil
}

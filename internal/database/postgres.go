package database

import (
	"blacklistApi/internal/config"
	"context"
	"fmt"
	pgx "github.com/jackc/pgx/v5"
	"os"
)

type Storage struct {
	Conn *pgx.Conn
}

var schema = `CREATE TABLE IF NOT EXISTS blacklist (
	id SERIAL PRIMARY KEY,
	phone_number VARCHAR(20) NOT NULL,
	user_name  VARCHAR(50) NOT NULL,
	reason_for_adding TEXT NOT NULL,
	adding_date TIMESTAMP NOT NULL,
	admin_name  VARCHAR(50) NOT NULL
);`

var tokenShema = `CREATE TABLE IF NOT EXISTS tokens (
	id SERIAL PRIMARY KEY,
	token  TEXT NOT NULL
);`

func InitConn(conf config.Conf) (*Storage, error) {
	dsn := os.Getenv("POSTGRES_URI")
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		fmt.Println(conn, dsn)
		fmt.Println(err)
		conn, err = pgx.Connect(context.Background(), conf.DbUrl)
		if err != nil {
			fmt.Println(conn, dsn, err)
			return &Storage{}, err
		}
	}

	fmt.Println(conn, "addr conn2")

	_, err = conn.Exec(context.Background(), schema)
	if err != nil {
		return nil, fmt.Errorf("users table creating error: %v", err)
	}
	_, err = conn.Exec(context.Background(), tokenShema)
	if err != nil {
		return nil, fmt.Errorf("tokens table creating error: %v", err)
	}
	return &Storage{conn}, nil
}

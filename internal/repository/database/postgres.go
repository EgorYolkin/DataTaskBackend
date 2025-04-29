package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

const (
	UsersTable = "users"
)

func ConnectPostgres(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

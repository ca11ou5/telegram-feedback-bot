package postgres

import (
	"database/sql"
	"feedback-bot/configs"
)

type Database struct {
	db *sql.DB
}

func NewDatabase(cfg *configs.Config) *Database {
	return &Database{
		db: ConnectToPostgres(cfg.PostgresURL),
	}
}

func (d *Database) CreateUser() error {
	return nil
}

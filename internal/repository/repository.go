package repository

import (
	"feedback-bot/configs"
	"feedback-bot/internal/repository/postgres"
)

type Repository struct {
	pg PgRepo
}

func NewRepository(cfg *configs.Config) *Repository {
	return &Repository{
		pg: postgres.NewDatabase(cfg),
	}
}

type PgRepo interface {
	CreateUser() error
}

func (r *Repository) CreateUser() error {
	return nil
}

package repository

import (
	"errors"
	"feedback-bot/configs"
	"feedback-bot/internal/entity"
	"feedback-bot/internal/repository/mongo"
	"feedback-bot/internal/repository/postgres"
)

type Repository struct {
	pg    PgRepo
	mongo *mongo.Mongo
}

func NewRepository(cfg *configs.Config) *Repository {
	return &Repository{
		pg:    postgres.NewDatabase(cfg),
		mongo: mongo.ConnectToMongo(cfg.MongoURL),
	}
}

type PgRepo interface {
	CreateUser() error
}

func (r *Repository) CreateUser() error {
	return nil
}

func (r *Repository) GetQuestion() (*entity.Question, error) {
	ans := r.mongo.GetQuestion()
	if ans == nil {
		return nil, errors.New("no answer found")
	}

	return ans, nil
}

func (r *Repository) InsertQA(qa *entity.QA) error {
	err := r.mongo.InsertQA(qa)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) InsertQuestion(question *entity.Question) error {
	err := r.mongo.InsertQuestion(question)
	if err != nil {
		return err
	}

	return nil
}

package service

import "time"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

type Repository interface {
	CreateUser() error
}

func (s *Service) Login() {
	time.Sleep(10 * time.Second)
}

package service

import (
	"feedback-bot/internal/entity"
	"log"
	"strings"
	"time"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

type Repository interface {
	CreateUser() error
	GetQuestion() (*entity.Question, error)
	InsertQuestion(question *entity.Question) error
}

func (s *Service) Login() {
	time.Sleep(10 * time.Second)
}

var AQ = map[string]map[string]string{
	"слово": {
		"Первый вопрос": "Ответ на первый вопросОтвет на пер",
		"Второй вопрос": "Ответ на второй вопрос",
	},
	"тест": {
		"Третий вопрос": "Ответ на третий вопрос",
	},
}

func (s *Service) GetAnswers(questionText string) []map[string]string {
	keywords := strings.Split(questionText, " ")
	var answers []map[string]string

	for _, keyword := range keywords {
		val, ok := AQ[keyword]
		if !ok {
			continue
		}

		answers = append(answers, val)
	}

	return answers
}

func (s *Service) InsertQuestion(questionText string) {
	question := &entity.Question{
		Message: questionText,
	}

	err := s.repo.InsertQuestion(question)
	if err != nil {
		log.Fatal(err)
	}
}

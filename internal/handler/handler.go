package handler

import (
	"feedback-bot/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Handler struct {
	svc *service.Service
}

func NewHandler(svc *service.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CheckUpdates(bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil || !update.Message.IsCommand() {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		switch update.Message.Command() {
		case "start":
			msg.Text = "Привет =)\nСейчас я все расскажу тебе о себе\nДля начала вам нужно авторизироваться, введите /login"
		case "login":
			msg.Text = "Введите ваш логин на сайте:"
			bot.Send(msg)
		}

	}
}

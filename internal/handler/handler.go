package handler

import (
	"feedback-bot/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type Handler struct {
	svc *service.Service
}

func NewHandler(svc *service.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CheckUpdates(bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) {
	for update := range updates {

		if update.Message != nil {
			if update.Message.IsCommand() {
				var msg tgbotapi.MessageConfig
				if update.Message == nil {
					msg = tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")
				} else {
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "")
				}
				switch update.Message.Command() {
				case "start":
					msg.Text = "Заглушка старт"
				case "login":
					msg.Text = "Заглушка логин"
				default:
					msg.Text = "ГОВНО"
				}
				bot.Send(msg)
			}
		}

		if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

			answers := h.svc.GetAnswers(update.Message.Text)
			if len(answers) > 0 {
				msg.Text = "Вот возможные ответы на ваш вопрос"
				var keyboard tgbotapi.InlineKeyboardMarkup
				for _, answer := range answers {
					for i, v := range answer {
						keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, []tgbotapi.InlineKeyboardButton{tgbotapi.NewInlineKeyboardButtonData(i, v)})
					}
				}
				msg.ReplyMarkup = keyboard
				bot.Send(msg)
			}

			msg.Text = "Возможные ответы на поставленный вопрос не найдены, задать вопрос эксперту?"
			keyboard := tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Да", "Вопрос отправлен, ожидайте ответа"),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Нет", "Сожалеем, что не смогли вам помочь"),
				),
			)
			msg.ReplyMarkup = keyboard
			bot.Send(msg)
		}

		if update.CallbackQuery != nil {
			callback := update.CallbackQuery
			if callback.Data == "Вопрос отправлен, ожидайте ответа" || callback.Data == "Сожалеем, что не смогли вам помочь" {
				edit := tgbotapi.NewEditMessageReplyMarkup(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, tgbotapi.InlineKeyboardMarkup{InlineKeyboard: make([][]tgbotapi.InlineKeyboardButton, 0)})
				_, err := bot.Send(edit)
				if err != nil {
					log.Fatal(err)
				}
				msg := tgbotapi.NewMessage(callback.Message.Chat.ID, "")
				msg.Text = callback.Data
				bot.Send(msg)
			}
		}

		//if update.CallbackQuery != nil {
		//	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")
		//	msg.Text = update.CallbackQuery.Data
		//	bot.Send(msg)
		//}

	}
}

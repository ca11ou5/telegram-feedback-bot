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
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "")
				switch update.Message.Command() {
				case "start":
					msg.Text = "Привет, я бот обратной связи\n" +
						"Напиши свой вопрос и я постараюсь на него ответить\n\n" +
						"Если ты являешься сотрудником для начала авторизируйся через /login"
					bot.Send(msg)
					continue
				case "login":
					msg.Text = "Введите ваш логин и пароль в формате:\n exampleLogin examplePassword"
					bot.Send(msg)
					continue
				case "next":
					b := h.svc.CheckPermissions(update.Message.Chat.ID)
					if b {
						h.svc.PrintQuestionToSupport(bot, msg)
						continue
					}
					msg.Text = "Вы не авторизованы, пожалуйста авторизуйтесь через /login"
					bot.Send(msg)
					continue
				default:
					msg.Text = "Пока что я не знаю такой команды =("
					bot.Send(msg)
					continue
				}
			}
		}

		if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

			b := h.svc.CheckPermissions(update.Message.Chat.ID)
			if b {

			}

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

			if len(answers) == 0 {
				msg.ReplyToMessageID = update.Message.MessageID
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
		}

		if update.CallbackQuery != nil {
			callback := update.CallbackQuery
			if callback.Data == "Вопрос отправлен, ожидайте ответа" {
				h.svc.InsertQuestion(update.CallbackQuery.Message.ReplyToMessage.Text)
			}

			if callback.Data == "Вопрос отправлен, ожидайте ответа" || callback.Data == "Сожалеем, что не смогли вам помочь" {
				edit := tgbotapi.NewEditMessageReplyMarkup(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, tgbotapi.InlineKeyboardMarkup{InlineKeyboard: make([][]tgbotapi.InlineKeyboardButton, 0)})
				_, err := bot.Send(edit)
				if err != nil {
					log.Fatal(err)
				}
				msg := tgbotapi.NewMessage(callback.Message.Chat.ID, "")
				msg.Text = callback.Data
				bot.Send(msg)
				continue
			}

			msg := tgbotapi.NewMessage(callback.Message.Chat.ID, "")
			msg.Text = update.CallbackQuery.Data
			bot.Send(msg)
		}

	}
}

package service

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

var Authorized = map[int64]bool{
	1349185687: true,
}

func (s *Service) CheckPermissions(id int64) bool {
	v, ok := Authorized[id]
	if v && ok {
		return true
	}
	return false
}

func (s *Service) PrintQuestionToSupport(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {
	ans, err := s.repo.GetQuestion()
	if err != nil {
		msg.Text = "Отличная работа\nВ данный момент вопросов нету, попробуйте позже =)"
		bot.Send(msg)
		return
	}

	msg.Text = ans.Message
	bot.Send(msg)
}

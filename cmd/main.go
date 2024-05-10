package main

import (
	"feedback-bot/configs"
	"feedback-bot/internal/handler"
	"feedback-bot/internal/repository"
	"feedback-bot/internal/repository/postgres"
	"feedback-bot/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	var cfg configs.Config

	err := cleanenv.ReadConfig("../envs/local.env", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	bot, err := tgbotapi.NewBotAPI(cfg.TelegramAPIToken)
	if err != nil {
		log.Panic(err)
	}

	db := postgres.ConnectToPostgres(cfg.PostgresURL)
	if db == nil {
		log.Panic("db is nil")
	}

	repo := repository.NewRepository(&cfg)
	svc := service.NewService(repo)
	h := handler.NewHandler(svc)

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}

	h.CheckUpdates(bot, updates)

	//for update := range updates {
	//	if update.Message != nil { // If we got a message
	//		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	//
	//		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	//		msg.ReplyToMessageID = update.Message.MessageID
	//
	//		bot.Send(msg)
	//	}
	//}
}

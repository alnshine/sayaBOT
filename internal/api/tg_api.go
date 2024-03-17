package api

import (
	"time"

	"github.com/alnshine/sayaBOT/internal/models"
	"github.com/alnshine/sayaBOT/internal/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

func RunTelegramAPI(log *logrus.Logger, token string, repo *repository.Repository) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		if update.Message == nil || update.Message.Sticker != nil {
			continue
		}

		chatID := update.Message.Chat.ID

		msg := tgbotapi.NewMessage(chatID, "")

		if update.Message.Chat.Type != "group" && update.Message.Chat.Type != "supergroup" {
			msg.Text = "Firstly, add me to group"
			bot.Send(msg)
			continue
		}

		var messageTime time.Time
		if update.Message.Date != 0 {
			messageTime = time.Unix(int64(update.Message.Date), 0)
		}

		message := models.Message{
			Content:  update.Message.Text,
			Username: update.Message.From.UserName,
			Time:     messageTime,
			ChatId:   update.Message.Chat.ID,
		}

		if err := repo.Message.CreateMessage(message); err != nil {
			log.Errorf("Failed to create message: %s", err.Error())
		}

		switch update.Message.Command() {
		case "start":
			msg.Text = "Начинаем запуск!"
			bot.Send(msg)
			continue
		case "shortHour":

		}
	}
}

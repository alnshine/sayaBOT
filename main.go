package main

import (
	"fmt"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()

	tgbotapi.SetLogger(log)

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error with loading env files: %s", err.Error())
	}

	token := os.Getenv("TOKEN")

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
		if update.Message == nil {
			continue
		}
		chatID := update.Message.Chat.ID

		msg := tgbotapi.NewMessage(chatID, "")

		if update.Message.Chat.Type != "group" && update.Message.Chat.Type != "supergroup" {
			msg.Text = "Я не могу нихуя делать. Даун, добавь сначала меня в группу!"
			bot.Send(msg)
			continue
		}

		switch update.Message.Command() {
		case "start":
			msg.Text = "Начинаем запуск!"
			bot.Send(msg)
			continue
		case "shortHour":
			messageID := update.Message.MessageID

			updates, err := bot.GetUpdates(tgbotapi.UpdateConfig{
				Offset:  messageID + 1,
				Limit:   1000,
				Timeout: 0,
			})
			if err != nil {
				log.Println(err)
				continue
			}

			for _, update := range updates {
				if update.Message != nil {
					fmt.Printf("[%s] %s\n", update.Message.From.UserName, update.Message.Text)
				}
			}
		}
	}
}

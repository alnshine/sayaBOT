package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()

	tgbotapi.SetLogger(log)

	bot, err := tgbotapi.NewBotAPI("6791189994:AAHhBbR36kIx4iRzi51kyWZZJBaR0em5FX4")
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

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		if update.Message.Chat.Type != "group" && update.Message.Chat.Type != "supergroup" {
			msg.Text = "Я не могу нихуя делать. Даун, добавь сначала меня в группу!"
			bot.Send(msg)
			continue
		}

		switch update.Message.Command() {
		case "start":
			msg.Text = "Начинаем запуск!"
		default:
			msg.Text = "Дебил, я не знаю что это за команда!!!"
		}

		bot.Send(msg)
	}
}

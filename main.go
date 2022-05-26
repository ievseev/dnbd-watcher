package main

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	token := os.Getenv("DNBD_TOKEN")
	bot, _ := tgbotapi.NewBotAPI(token)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hochew eshe poboltat, "+update.Message.Chat.FirstName+"?")

		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
	}
}

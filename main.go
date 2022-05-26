package main

import (
	"os"
	"math/rand"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	phrasesDictionary:=[]string{"раз","два","три",")"}

	token := os.Getenv("DNBD_TOKEN")
	bot, _ := tgbotapi.NewBotAPI(token)

	updateConfig := tgbotapi.NewUpdate(0)
    updateConfig.Timeout = 30
    updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
        // Telegram can send many types of updates depending on what your Bot
        // is up to. We only want to look at messages for now, so we can
        // discard any other updates.
        if update.Message == nil {
            continue
        }

        // Now that we know we've gotten a new message, we can construct a
        // reply! We'll take the Chat ID and Text from the incoming message
        // and use it to create a new message.
        randomIndex := rand.Intn(len(phrasesDictionary))
        msg := tgbotapi.NewMessage(update.Message.Chat.ID, phrasesDictionary[randomIndex])
        // We'll also say that this message is a reply to the previous message.
        // For any other specifications than Chat ID or Text, you'll need to
        // set fields on the `MessageConfig`.
        // msg.ReplyToMessageID = update.Message.MessageID

        // Okay, we're sending our message off! We don't care about the message
        // we just sent, so we'll discard it.
        if _, err := bot.Send(msg); err != nil {
            // Note that panics are a bad way to handle errors. Telegram can
            // have service outages or network errors, you should retry sending
            // messages or more gracefully handle failures.
            panic(err)
        }
	}
}

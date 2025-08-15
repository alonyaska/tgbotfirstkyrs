package main

import (
	"log"

	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	token := "7996936625:AAEJnIFa_JH0jBcLibhjk2-YqGUWUGtZmu4"
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("bot init %v", err)
	}
	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 5
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() && update.Message.Command() == "start" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "I alive")
			if _, err := bot.Send(msg); err != nil {
				log.Println("send /start again", err)

			}
			continue
		}
		if len(update.Message.NewChatMembers) > 0 {
			for _, newMember := range update.Message.NewChatMembers {
				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("tgbot/tgbot.jpg"))

				if newMember.UserName != "" {
					photo.Caption = "Добро пожаловать, @" + newMember.UserName + "!"
				} else {
					photo.Caption = fmt.Sprintf("Добро пожаловать, %s %s!", newMember.FirstName, newMember.LastName)
				}

				bot.Send(photo)
			}
		}
	}

}

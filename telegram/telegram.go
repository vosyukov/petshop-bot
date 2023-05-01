package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var botApi *tgbotapi.BotAPI

func Test() {
	fmt.Println("dsf")
}

func init() {
	bot, err := tgbotapi.NewBotAPI("5292301851:AAFLzeZ3c4-tTauZMhcxHHtZ9__UHFgk-zo")
	if err != nil {
		log.Panic(err)
	}

	botApi = bot

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	//updates := bot.GetUpdatesChan(u)
	//
	//for update := range updates {
	//	if update.Message != nil { // If we got a message
	//		log.Printf("[%d] %s", update.Message.Chat.ID, update.Message.Text)
	//		//263537201
	//		SendMessage(263537201, "kjdfnskdjfn")
	//		msg := tgbotapi.NewMessage(263537201, update.Message.Text)
	//		msg.ReplyToMessageID = update.Message.MessageID
	//
	//		bot.Send(msg)
	//	}
	//}
}

func SendMessage(chatId int64, message string) {
	msg := tgbotapi.NewMessage(chatId, message)

	botApi.Send(msg)
}

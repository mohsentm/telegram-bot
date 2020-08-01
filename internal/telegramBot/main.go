package telegrambot

import (
	"log"

	"github.com/elastic/go-elasticsearch/v7"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/mohsentm/telegram-bot/config"
)

func WakeUp() {
	conf := config.Get()

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	bot, err := tgbotapi.NewBotAPI(conf.ApiToken)
	if err != nil {
		log.Panic(err)
	}
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	log.Println(res)

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		go parseUpdate(bot, update)
	}
}

func parseUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	switch {
	case update.Message.Audio != nil:
		shareAudio(bot, update)
	case update.Message.Text != "":
		replyMessage(bot, update)
	}
}

func shareAudio(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	msg := tgbotapi.NewAudioShare(update.Message.Chat.ID, update.Message.Audio.FileID)
	msg.ReplyToMessageID = update.Message.MessageID
	msg.Caption = update.Message.Caption
	bot.Send(msg)
}

func replyMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	msg.ReplyToMessageID = update.Message.MessageID

	bot.Send(msg)
}

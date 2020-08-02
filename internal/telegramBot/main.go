package telegrambot

import (
	"fmt"
	"log"
	"reflect"

	"github.com/elastic/go-elasticsearch/v7"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/mohsentm/telegram-bot/config"
	elasticSerive "github.com/mohsentm/telegram-bot/internal/elasticservice"
	"github.com/olivere/elastic"
)

var indexName = "telegram"

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

	service := elasticSerive.InitClient(indexName)
	service.CheckOrCreateIndex()
	audioData := elasticSerive.AudioData{
		Performer: update.Message.Audio.Performer,
		FileID:    update.Message.Audio.FileID,
		Title:     update.Message.Audio.Title,
		Caption:   update.Message.Caption,
	}
	_, err := service.IndexMessage(audioData)
	if err != nil {
		// Handle error
		log.Fatalf("Error getting response: %s", err)
	}

	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	log.Printf("file name %s", update.Message.Caption)
	// msg := tgbotapi.NewAudioShare(update.Message.Chat.ID, update.Message.Audio.FileID)
	// // msg.ReplyToMessageID = update.Message.MessageID
	// msg.Caption = update.Message.Caption
	// bot.Send(msg)
}

func replyMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) {

	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	service := elasticSerive.InitClient(indexName)

	termQuery := elastic.NewMultiMatchQuery(update.Message.Text, "title", "caption", "performer")

	searchResult, err := service.Search(termQuery)
	if err != nil {
		// Handle error
		log.Fatalf("Error getting result: %s", err)
	}

	// fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)

	// TotalHits is another convenience function that works even when something goes wrong.
	fmt.Printf("Found a total of %d tweets\n", searchResult.TotalHits())

	// 	// Each is a convenience function that iterates over hits in a search result.
	// 	// It makes sure you don't need to check for nil values in the response.
	// 	// However, it ignores errors in serialization. If you want full control
	// 	// over iterating the hits, see below.
	var ttyp elasticSerive.AudioData
	for _, files := range searchResult.Each(reflect.TypeOf(ttyp)) {
		file := files.(elasticSerive.AudioData)

		fmt.Printf("Tweet by %s: %s\n", file.FileID, file.Caption)

		msg := tgbotapi.NewAudioShare(update.Message.Chat.ID, file.FileID)
		// msg.ReplyToMessageID = update.Message.MessageID
		msg.Caption = file.Caption
		bot.Send(msg)

	}

	// msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	// 	msg.ReplyToMessageID = update.Message.MessageID
	// 	bot.Send(msg)

}

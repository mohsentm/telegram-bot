// package elastic-service
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/mohsentm/telegram-bot/config"
	"github.com/olivere/elastic"
)

func main() {

	client := InitClient()
	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := client.Ping("http://elasticsearch:9200").Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

}

/*
 * init new client
 */
func InitClient() *elastic.Client {
	conf := config.Get()

	errorlog := log.New(os.Stdout, "APP ", log.LstdFlags)
	// Obtain a client. You can also provide your own HTTP client here.
	client, err := elastic.NewClient(
		elastic.SetURL(conf.ElasticSearchURL),
		elastic.SetHealthcheck(true),
		elastic.SetSniff(false),
		elastic.SetErrorLog(errorlog),
	)
	// Trace request and response details like this
	// client, err := elastic.NewClient(elastic.SetTraceLog(log.New(os.Stdout, "", 0)))
	if err != nil {
		// Handle error
		panic(err)
	}

	return client
}

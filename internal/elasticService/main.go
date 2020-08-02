package elasticservice

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/mohsentm/telegram-bot/config"
	"github.com/olivere/elastic"
)

type Service struct {
	Client *elastic.Client
}

type Data struct {
	User     string `json:"user"`
	Message  string `json:"message"`
	Retweets string `json:"retweets"`
}

/*
 * InitClient efwef
 * @return *elastic.Client
 */
func InitClient() *Service {
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
	return &Service{Client: client}
}

func (service *Service) Ping() {
	info, code, err := service.Client.Ping("http://elasticsearch:9200").Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
}

func (service *Service) IndexMessage(indexName string, bodyData interface{}) (*elastic.IndexResponse, error) {
	return service.Client.Index().
		Type("_doc").
		Index(indexName).
		BodyJson(bodyData).
		Do(context.Background())
}

func (service *Service) CheckOrCreateIndex(indexName string) {
	exists, err := service.Client.IndexExists(indexName).Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}

	if !exists {
		createIndex, err := service.Client.CreateIndex(indexName).Body(GetMapping()).Do(context.Background())
		if err != nil {
			// Handle error
			panic(err)
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
		}
	}
}

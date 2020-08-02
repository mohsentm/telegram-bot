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
	Client    *elastic.Client
	IndexName string
}

type Data struct {
	User     string `json:"user"`
	Message  string `json:"message"`
	Retweets string `json:"retweets"`
}

type AudioData struct {
	FileID  string `json:"file_id"`
	Title   string `json:"title"`
	Caption string `json:"caption"`
}

/*
 * InitClient efwef
 * @return *elastic.Client
 */
func InitClient(indexName string) *Service {
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
	return &Service{Client: client, IndexName: indexName}
}

func (service *Service) Ping() {
	info, code, err := service.Client.Ping("http://elasticsearch:9200").Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
}

func (service *Service) IndexMessage(bodyData interface{}) (*elastic.IndexResponse, error) {
	return service.Client.Index().
		Type("_doc").
		Index(service.IndexName).
		BodyJson(bodyData).
		Do(context.Background())
}

func (service *Service) CheckOrCreateIndex() {
	exists, err := service.Client.IndexExists(service.IndexName).Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}

	if !exists {
		createIndex, err := service.Client.CreateIndex(service.IndexName).Body(GetMapping(service.IndexName)).Do(context.Background())
		if err != nil {
			// Handle error
			panic(err)
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
		}
	}
}

func (service *Service) Search(termQuery elastic.Query) (*elastic.SearchResult, error) {
	return service.Client.Search().
		RestTotalHitsAsInt(true).
		Index(service.IndexName). // search in index "twitter"
		Query(termQuery).         // specify the query
		// Sort("user", true).       // sort by "user" field, ascending
		// From(0).Size(10).        // take documents 0-9
		// Pretty(true).            // pretty print request and response JSON
		Do(context.Background()) // execute
}

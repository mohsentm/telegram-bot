package config

type Config struct {
	ApiToken         string
	ElasticSearchURL string
}

var defaultConfig = Config{
	ApiToken:         "BotAPIToken",
	ElasticSearchURL: "http://elasticsearch:9200",
}

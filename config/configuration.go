package config

type Config struct {
	ApiToken string
}

var defaultConfig = Config{
	ApiToken: "BotAPIToken",
}

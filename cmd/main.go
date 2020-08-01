package main

import (
	"github.com/mohsentm/telegram-bot/config"
	telegramBot "github.com/mohsentm/telegram-bot/internal/telegrambot"
)

func init() {
	config.Init()
}

func main() {
	telegramBot.WakeUp()
}

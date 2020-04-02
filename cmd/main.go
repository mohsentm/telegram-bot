package main

import (
	"github.com/mohsentm/telegram-bot/config"
	"github.com/mohsentm/telegram-bot/internal/telegramBot"
)

func init() {
	config.Init()
}

func main() {
	telegramBot.WakeUp()
}

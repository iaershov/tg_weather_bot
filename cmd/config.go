package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func getUpdateConfig() tgbotapi.UpdateConfig {
	cfg := tgbotapi.NewUpdate(0)
	cfg.Timeout = 5

	return cfg
}

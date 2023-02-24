package main

import (
	"context"
	"fmt"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"tg_weather_bot/internal/gateway/weather"
	"tg_weather_bot/internal/handler/updates"
	"tg_weather_bot/internal/models"
)

func getUpdateConfig() tgbotapi.UpdateConfig {
	cfg := tgbotapi.NewUpdate(0)
	cfg.Timeout = 5

	return cfg
}

func main() {
	fmt.Println("Starting bot")

	ctx := context.Background()

	bot, err := tgbotapi.NewBotAPI(models.TelegramBotToken)
	if err != nil {
		panic(err)
	}

	updateConfig := getUpdateConfig()

	httpClient := &http.Client{}
	weatherGateway := weather.NewGateway(models.WeatherApiToken, httpClient)

	handler := updates.NewHandler(bot, weatherGateway)

	handler.HandleMessages(ctx, updateConfig)
}

package updates

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"tg_weather_bot/internal/models"
)

type bot interface {
	GetUpdatesChan(config tgbotapi.UpdateConfig) tgbotapi.UpdatesChannel
	Send(c tgbotapi.Chattable) (tgbotapi.Message, error)
}

type weatherGateway interface {
	GetWeatherByCity(ctx context.Context, city string) (models.Weather, error)
}

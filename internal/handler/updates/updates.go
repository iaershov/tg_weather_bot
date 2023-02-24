package updates

import (
	"context"
	"errors"
	"fmt"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"tg_weather_bot/internal/models"
	presenter "tg_weather_bot/internal/presenter/weather"
)

type Handler struct {
	bot     bot
	weather weatherGateway
}

func NewHandler(b bot, wg weatherGateway) *Handler {
	return &Handler{
		bot:     b,
		weather: wg,
	}
}

func (h *Handler) HandleMessages(ctx context.Context, updateCfg tgbotapi.UpdateConfig) {
	updates := h.bot.GetUpdatesChan(updateCfg)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func(updates tgbotapi.UpdatesChannel) {
		defer wg.Done()

		for update := range updates {
			if update.Message == nil {
				continue
			}

			command := update.Message.Command()
			if len(command) != 0 {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Напишите название города")
				_, _ = h.bot.Send(msg)
				continue
			}

			// Получаем текущую погоду по названию города
			weather, err := h.weather.GetWeatherByCity(ctx, update.Message.Text)
			if err != nil {
				if errors.Is(err, models.ErrInvalidCityName) {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Некорректное название города")
					_, _ = h.bot.Send(msg)
				}
				continue
			}

			answer := presenter.PresentCurrentWeather(weather)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, answer)

			_, err = h.bot.Send(msg)
			if err != nil {
				fmt.Println("failed to send message", err)
				continue
			}
		}
	}(updates)

	wg.Wait()
}

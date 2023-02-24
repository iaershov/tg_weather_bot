package weather

import (
	"errors"

	"tg_weather_bot/internal/models"
)

func toWeather(weather weather) (models.Weather, error) {
	if len(weather.Description) == 0 {
		return models.Weather{}, errors.New("received empty description")
	}

	description := weather.Description[0]

	return models.Weather{
		City: weather.City,
		Temperature: models.Temperature{
			Current:          weather.Temperature.Current,
			CurrentFeelsLike: weather.Temperature.CurrentFeelsLike,
			Min:              weather.Temperature.Min,
			Max:              weather.Temperature.Max,
		},
		Description: models.WeatherDescription{
			ID:   description.ID,
			Text: description.Text,
			Main: description.Main,
		},
	}, nil
}

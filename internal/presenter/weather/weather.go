package weather

import (
	"fmt"

	"tg_weather_bot/internal/models"
	"tg_weather_bot/internal/presenter"
)

var icons = map[string]string{
	"Thunderstorm": "🌩️",
	"Drizzle":      "🌧️",
	"Rain":         "🌧️",
	"Snow":         "🌨️",
	"Clouds":       "☁️",
	"Clear":        "☀️",
}

func PresentCurrentWeather(weather models.Weather) string {
	header := presentHeader(weather)
	description := presentDescription(weather)
	temperature := presentTemperature(weather)

	return header + description + temperature
}

func presentHeader(weather models.Weather) string {
	return fmt.Sprintf("Погода в городе %s\n\n", weather.City)
}

func presentDescription(weather models.Weather) string {
	description := presenter.ToUpper(weather.Description.Text, 0)
	icon := icons[weather.Description.Main]

	return description + " " + icon + "\n\n"
}

func presentTemperature(weather models.Weather) string {
	return fmt.Sprintf(
		"Температура: %.1f°C\n"+
			"Ощущается как: %.1f°C\n"+
			"Минимум сегодня: %.1f°C\n"+
			"Максимум сегодня: %.1f°C\n",
		weather.Temperature.Current,
		weather.Temperature.CurrentFeelsLike,
		weather.Temperature.Min,
		weather.Temperature.Max,
	)
}

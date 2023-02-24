package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"tg_weather_bot/internal/models"
)

const (
	getWeatherUrlTemplate = "https://api.openweathermap.org/data/2.5/weather?q=%s,ru&appid=%s&lang=ru&units=metric"
)

type Gateway struct {
	key        string
	httpClient httpClient
}

func NewGateway(key string, cl httpClient) *Gateway {
	return &Gateway{
		key:        key,
		httpClient: cl,
	}
}

func (g *Gateway) GetWeatherByCity(ctx context.Context, city string) (models.Weather, error) {
	url := fmt.Sprintf(getWeatherUrlTemplate, city, g.key)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return models.Weather{}, fmt.Errorf("failed to create http request: %w", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req = req.WithContext(ctx)

	resp, err := g.httpClient.Do(req)
	if err != nil {
		return models.Weather{}, fmt.Errorf("failed to perform http request: %w", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return models.Weather{}, models.ErrInvalidCityName
		}

		return models.Weather{}, fmt.Errorf("received invalid response code")
	}

	var w weather
	err = json.NewDecoder(resp.Body).Decode(&w)
	if err != nil {
		return models.Weather{}, fmt.Errorf("failed to unmarshall response: %w", err)
	}

	return toWeather(w)
}

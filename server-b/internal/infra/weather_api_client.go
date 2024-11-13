package infra

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jhonasalves/go-expert-fc-labs-otel/server-b/internal/entity"
)

type WeatherAPIClient struct {
	APIKey string
}

type WeatherAPIResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

func (c *WeatherAPIClient) FetchWeather(city string) (*entity.Weather, error) {
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", c.APIKey, city)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch weather data: status %d", resp.StatusCode)
	}

	var apiResp WeatherAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, err
	}

	weather := &entity.Weather{
		TempC: apiResp.Current.TempC,
	}

	weather.Convert()

	return weather, nil
}

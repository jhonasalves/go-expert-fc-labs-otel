package infra

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/jhonasalves/go-expert-fc-labs-otel/weather-api/internal/entity"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type WeatherAPIClient struct {
	APIKey string
}

type WeatherAPIResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

func (c *WeatherAPIClient) FetchWeather(ctx context.Context, city string) (*entity.Weather, error) {
	ctx, span := tracer.Start(ctx, "FetchWeather")
	defer span.End()

	encodedCity := url.QueryEscape(city)

	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", c.APIKey, encodedCity)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}

	resp, err := client.Do(req)
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
		City:  city,
		TempC: apiResp.Current.TempC,
	}

	weather.Convert()

	return weather, nil
}

package infra

import (
	"context"

	"github.com/jhonasalves/go-expert-fc-labs-otel/weather-api/internal/entity"
)

type WeatherRepository struct {
	WeatherAPIClient WeatherAPIClient
	ViaCepClient     ViaCepClient
}

func NewWeatherRepository(weatherAPIClient WeatherAPIClient, viaCepClient ViaCepClient) *WeatherRepository {
	return &WeatherRepository{
		WeatherAPIClient: weatherAPIClient,
		ViaCepClient:     viaCepClient,
	}
}

func (r *WeatherRepository) GetWeather(ctx context.Context, city string) (*entity.Weather, error) {
	return r.WeatherAPIClient.FetchWeather(ctx, city)
}

func (r *WeatherRepository) GetLocationByZipCode(ctx context.Context, zipcode string) (*entity.Location, error) {
	return r.ViaCepClient.FetchLocation(ctx, zipcode)
}

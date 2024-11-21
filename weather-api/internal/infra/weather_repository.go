package infra

import "github.com/jhonasalves/go-expert-fc-labs-otel/weather-api/internal/entity"

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

func (r *WeatherRepository) GetWeather(city string) (*entity.Weather, error) {
	return r.WeatherAPIClient.FetchWeather(city)
}

func (r *WeatherRepository) GetLocationByZipCode(zipcode string) (*entity.Location, error) {
	return r.ViaCepClient.FetchLocation(zipcode)
}

package usecase

import (
	"context"
	"fmt"

	"github.com/jhonasalves/go-expert-fc-labs-otel/weather-api/internal/entity"
	"github.com/jhonasalves/go-expert-fc-labs-otel/weather-api/internal/infra"
)

type WeatherUseCase struct {
	WeatherRepository *infra.WeatherRepository
}

func NewWeatherUseCase(weatherRepository *infra.WeatherRepository) *WeatherUseCase {
	return &WeatherUseCase{
		WeatherRepository: weatherRepository,
	}
}

func (uc *WeatherUseCase) GetWeather(ctx context.Context, city string) (*entity.Weather, error) {
	weatherData, err := uc.WeatherRepository.GetWeather(ctx, city)
	if err != nil {
		return nil, fmt.Errorf("could not fetch weather data: %v", err)
	}

	return weatherData, nil
}

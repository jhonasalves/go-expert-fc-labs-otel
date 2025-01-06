package usecase

import (
	"context"
	"fmt"

	"github.com/jhonasalves/go-expert-fc-labs-otel/weather-api/internal/entity"
	"github.com/jhonasalves/go-expert-fc-labs-otel/weather-api/internal/infra"
)

type LocationUseCase struct {
	WeatherRepository *infra.WeatherRepository
}

func NewLocationUseCase(weatherRepository *infra.WeatherRepository) *LocationUseCase {
	return &LocationUseCase{
		WeatherRepository: weatherRepository,
	}
}

func (uc *LocationUseCase) GetLocation(ctx context.Context, zipCode string) (*entity.Location, error) {
	locationData, err := uc.WeatherRepository.GetLocationByZipCode(ctx, zipCode)
	if err != nil {
		return nil, fmt.Errorf("could not fetch location data: %v", err)
	}

	return locationData, nil
}

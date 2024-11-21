package usecase

import (
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

func (uc *LocationUseCase) GetLocation(zipCode string) (*entity.Location, error) {
	locationData, err := uc.WeatherRepository.GetLocationByZipCode(zipCode)
	if err != nil {
		return nil, fmt.Errorf("could not fetch location data: %v", err)
	}

	return locationData, nil
}

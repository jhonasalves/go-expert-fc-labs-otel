package usecase

import (
	"github.com/jhonasalves/go-expert-fc-labs-otel/server-a/internal/entity"
	"github.com/jhonasalves/go-expert-fc-labs-otel/server-a/internal/infra"
)

type ZipCodeUseCase struct {
	ZipCodeRepository infra.ZipCodeRepository
}

func NewZipCodeUseCase(zipCodeRepository infra.ZipCodeRepository) *ZipCodeUseCase {
	return &ZipCodeUseCase{
		ZipCodeRepository: zipCodeRepository,
	}
}

func (u *ZipCodeUseCase) ProcessZipCode(zipCode string) (*entity.Weather, error) {
	zip := entity.ZipCode{Value: zipCode}

	if err := zip.IsValid(); err != nil {
		return nil, err
	}

	weather, err := u.ZipCodeRepository.GetWeather(zip)
	if err != nil {
		return nil, err
	}

	return weather, nil
}

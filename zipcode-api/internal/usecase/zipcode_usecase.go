package usecase

import (
	"context"

	"github.com/jhonasalves/go-expert-fc-labs-otel/zipcode-api/internal/entity"
	"github.com/jhonasalves/go-expert-fc-labs-otel/zipcode-api/internal/infra"
)

type ZipCodeUseCase struct {
	ZipCodeRepository infra.ZipCodeRepository
}

func NewZipCodeUseCase(zipCodeRepository infra.ZipCodeRepository) *ZipCodeUseCase {
	return &ZipCodeUseCase{
		ZipCodeRepository: zipCodeRepository,
	}
}

func (u *ZipCodeUseCase) ProcessZipCode(ctx context.Context, zipCode string) (*entity.Weather, error) {
	zip := entity.ZipCode{Value: zipCode}

	if err := zip.IsValid(); err != nil {
		return nil, err
	}

	weather, err := u.ZipCodeRepository.GetWeather(ctx, zip)
	if err != nil {
		return nil, err
	}

	return weather, nil
}

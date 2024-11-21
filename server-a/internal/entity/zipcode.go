package entity

import (
	"errors"
	"regexp"
)

type ZipCode struct {
	Value string
}

type Weather struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

func (z *ZipCode) IsValid() error {
	if len(z.Value) != 8 {
		return errors.New("invalid length")
	}

	matched, _ := regexp.MatchString(`^\d{8}$`, z.Value)
	if !matched {
		return errors.New("invalid format")
	}

	return nil
}

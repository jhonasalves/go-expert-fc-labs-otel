package infra

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jhonasalves/go-expert-fc-labs-otel/zipcode-api/internal/entity"
)

type ZipCodeRepository interface {
	GetWeather(zipCode entity.ZipCode) (*entity.Weather, error)
}

type HTTPRepository struct {
	ServiceBURL string
}

func NewHTTPRepository(serviceBURL string) *HTTPRepository {
	return &HTTPRepository{ServiceBURL: serviceBURL}
}

func (r *HTTPRepository) GetWeather(zipCode entity.ZipCode) (*entity.Weather, error) {
	url := fmt.Sprintf("%s/%s", r.ServiceBURL, zipCode.Value)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making GET request to ServiceB: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ServiceB returned status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var weather entity.Weather
	if err := json.Unmarshal(body, &weather); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	return &weather, nil
}

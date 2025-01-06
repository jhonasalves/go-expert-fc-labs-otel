package infra

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jhonasalves/go-expert-fc-labs-otel/zipcode-api/internal/entity"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
)

type ZipCodeRepository interface {
	GetWeather(ctx context.Context, zipCode entity.ZipCode) (*entity.Weather, error)
}

type HTTPRepository struct {
	ServiceBURL string
}

var tracer = otel.Tracer("ZipCodeAPI")

func NewHTTPRepository(serviceBURL string) *HTTPRepository {
	return &HTTPRepository{ServiceBURL: serviceBURL}
}

func (r *HTTPRepository) GetWeather(ctx context.Context, zipCode entity.ZipCode) (*entity.Weather, error) {
	ctx, span := tracer.Start(ctx, "GetWeather")
	defer span.End()

	url := fmt.Sprintf("%s/%s", r.ServiceBURL, zipCode.Value)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error making GET request to ServiceB: %w", err)
	}

	client := http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
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

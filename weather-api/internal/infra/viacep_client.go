package infra

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jhonasalves/go-expert-fc-labs-otel/weather-api/internal/entity"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
)

type ViaCepClient struct{}

var tracer = otel.Tracer("ServiceWeatherAPI")

func (c *ViaCepClient) FetchLocation(ctx context.Context, zipCode string) (*entity.Location, error) {
	ctx, span := tracer.Start(ctx, "FetchLocation")
	defer span.End()

	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", zipCode)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var location entity.Location
	if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
		return nil, err
	}

	if location.City == "" {
		return nil, fmt.Errorf("can not find zipcode")
	}

	return &location, nil
}

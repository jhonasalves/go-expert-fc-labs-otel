package infra

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jhonasalves/go-expert-fc-labs-otel/server-b/internal/entity"
)

type ViaCepClient struct{}

func (c *ViaCepClient) FetchLocation(zipCode string) (*entity.Location, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", zipCode)
	resp, err := http.Get(url)
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

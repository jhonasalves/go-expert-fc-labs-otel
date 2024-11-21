package main

import (
	"log"
	"net/http"

	"github.com/jhonasalves/go-expert-fc-labs-otel/zipcode-api/configs"
	"github.com/jhonasalves/go-expert-fc-labs-otel/zipcode-api/internal/handler"
	"github.com/jhonasalves/go-expert-fc-labs-otel/zipcode-api/internal/infra"
	"github.com/jhonasalves/go-expert-fc-labs-otel/zipcode-api/internal/usecase"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	serviceBURL := configs.WeatherAPIURL

	repo := infra.NewHTTPRepository(serviceBURL)
	useCase := usecase.NewZipCodeUseCase(repo)
	zipCodeHandler := handler.NewZipCodeHandler(useCase)

	http.HandleFunc("/zipcode", zipCodeHandler.HandleZipCode)

	log.Println("Server is running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

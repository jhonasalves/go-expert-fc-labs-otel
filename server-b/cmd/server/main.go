package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jhonasalves/go-expert-fc-labs-otel/server-b/configs"
	"github.com/jhonasalves/go-expert-fc-labs-otel/server-b/internal/handler"
	"github.com/jhonasalves/go-expert-fc-labs-otel/server-b/internal/infra"
	"github.com/jhonasalves/go-expert-fc-labs-otel/server-b/internal/usecase"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	weatherAPIClient := infra.WeatherAPIClient{APIKey: configs.WeatherAPIKey}
	viaCepClient := infra.ViaCepClient{}

	weatherRepository := infra.NewWeatherRepository(weatherAPIClient, viaCepClient)

	weatherUseCase := usecase.NewWeatherUseCase(weatherRepository)
	locationUseCase := usecase.NewLocationUseCase(weatherRepository)

	weatherHandler := handler.NewWeatherHandler(weatherUseCase, locationUseCase)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/weather/{zipCode}", weatherHandler.GetWeatherByZip)

	http.ListenAndServe(":8080", r)
}

package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/jhonasalves/go-expert-fc-labs-otel/pkg/opentelemetry"
	"github.com/jhonasalves/go-expert-fc-labs-otel/weather-api/configs"
	"github.com/jhonasalves/go-expert-fc-labs-otel/weather-api/internal/handler"
	"github.com/jhonasalves/go-expert-fc-labs-otel/weather-api/internal/infra"

	"github.com/jhonasalves/go-expert-fc-labs-otel/weather-api/internal/usecase"
)

func main() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	otelShutdown, err := opentelemetry.InitTracer(ctx, configs.ZipkinURL, "weather-api")
	if err != nil {
		return
	}

	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

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

	select {
	case <-sigCh:
		log.Println("Shutting down gracefully, CTRL+C pressed...")
	case <-ctx.Done():
		log.Println("Shutting down due to other reason...")
	}
}

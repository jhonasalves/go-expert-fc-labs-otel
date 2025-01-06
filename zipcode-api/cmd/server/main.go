package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/jhonasalves/go-expert-fc-labs-otel/pkg/opentelemetry"
	"github.com/jhonasalves/go-expert-fc-labs-otel/zipcode-api/configs"
	"github.com/jhonasalves/go-expert-fc-labs-otel/zipcode-api/internal/handler"
	"github.com/jhonasalves/go-expert-fc-labs-otel/zipcode-api/internal/infra"
	"github.com/jhonasalves/go-expert-fc-labs-otel/zipcode-api/internal/usecase"
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

	otelShutdown, err := opentelemetry.InitTracer(ctx, configs.ZipkinURL, "zipcode-api")
	if err != nil {
		return
	}

	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	serviceBURL := configs.WeatherAPIURL

	repo := infra.NewHTTPRepository(serviceBURL)
	useCase := usecase.NewZipCodeUseCase(repo)
	zipCodeHandler := handler.NewZipCodeHandler(useCase)

	http.HandleFunc("/zipcode", zipCodeHandler.HandleZipCode)

	log.Println("Server is running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))

	select {
	case <-sigCh:
		log.Println("Shutting down gracefully, CTRL+C pressed...")
	case <-ctx.Done():
		log.Println("Shutting down due to other reason...")
	}
}

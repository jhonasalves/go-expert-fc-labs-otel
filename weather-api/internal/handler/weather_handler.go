package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jhonasalves/go-expert-fc-labs-otel/weather-api/internal/usecase"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type WeatherHandler struct {
	WeatherUseCase  *usecase.WeatherUseCase
	LocationUseCase *usecase.LocationUseCase
}

var tracer = otel.Tracer("ServiceWeatherAPI")

func NewWeatherHandler(weatherUseCase *usecase.WeatherUseCase, locationUseCase *usecase.LocationUseCase) *WeatherHandler {
	return &WeatherHandler{
		WeatherUseCase:  weatherUseCase,
		LocationUseCase: locationUseCase,
	}
}

func (h *WeatherHandler) GetWeatherByZip(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	carrier := propagation.HeaderCarrier(r.Header)
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	ctx, span := tracer.Start(ctx, "GetWeatherByZip")
	defer span.End()

	zipCode := chi.URLParam(r, "zipCode")

	if len(zipCode) != 8 || zipCode == "" {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	location, err := h.LocationUseCase.GetLocation(ctx, zipCode)
	if err != nil {
		http.Error(w, "could not fetch location", http.StatusInternalServerError)
		return
	}

	weather, err := h.WeatherUseCase.GetWeather(ctx, location.City)
	if err != nil {
		http.Error(w, "could not fetch weather data", http.StatusInternalServerError)
		return
	}

	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(r.Header))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(weather)
}

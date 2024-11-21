package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jhonasalves/go-expert-fc-labs-otel/weather-api/internal/usecase"
)

type WeatherHandler struct {
	WeatherUseCase  *usecase.WeatherUseCase
	LocationUseCase *usecase.LocationUseCase
}

func NewWeatherHandler(weatherUseCase *usecase.WeatherUseCase, locationUseCase *usecase.LocationUseCase) *WeatherHandler {
	return &WeatherHandler{
		WeatherUseCase:  weatherUseCase,
		LocationUseCase: locationUseCase,
	}
}

func (h *WeatherHandler) GetWeatherByZip(w http.ResponseWriter, r *http.Request) {
	zipCode := chi.URLParam(r, "zipCode")

	if len(zipCode) != 8 || zipCode == "" {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	location, err := h.LocationUseCase.GetLocation(zipCode)
	if err != nil {
		http.Error(w, "could not fetch location", http.StatusInternalServerError)
		return
	}

	weather, err := h.WeatherUseCase.GetWeather(location.City)
	if err != nil {
		http.Error(w, "could not fetch weather data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(weather)
}

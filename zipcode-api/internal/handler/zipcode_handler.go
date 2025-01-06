package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jhonasalves/go-expert-fc-labs-otel/zipcode-api/internal/usecase"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type ZipCodeHandler struct {
	UseCase *usecase.ZipCodeUseCase
}

var tracer = otel.Tracer("ZipCodeAPI")

func NewZipCodeHandler(useCase *usecase.ZipCodeUseCase) *ZipCodeHandler {
	return &ZipCodeHandler{UseCase: useCase}
}

func (h *ZipCodeHandler) HandleZipCode(w http.ResponseWriter, r *http.Request) {
	var input struct {
		CEP string `json:"cep"`
	}

	ctx := r.Context()

	carrier := propagation.HeaderCarrier(r.Header)
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	ctx, span := tracer.Start(ctx, "HandleZipCode")
	defer span.End()

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	fmt.Printf("received request with cep: %s\n", input.CEP)

	weather, err := h.UseCase.ProcessZipCode(ctx, input.CEP)
	if err != nil {
		fmt.Printf("error processing zip code: %v\n", err)
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(r.Header))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(weather); err != nil {
		fmt.Printf("error encoding response: %v\n", err)
		http.Error(w, "error generating response", http.StatusInternalServerError)
		return
	}
}

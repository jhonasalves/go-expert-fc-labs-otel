package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jhonasalves/go-expert-fc-labs-otel/server-a/internal/usecase"
)

type ZipCodeHandler struct {
	UseCase *usecase.ZipCodeUseCase
}

func NewZipCodeHandler(useCase *usecase.ZipCodeUseCase) *ZipCodeHandler {
	return &ZipCodeHandler{UseCase: useCase}
}

func (h *ZipCodeHandler) HandleZipCode(w http.ResponseWriter, r *http.Request) {
	var input struct {
		CEP string `json:"cep"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	fmt.Printf("received request with cep: %s\n", input.CEP)

	weather, err := h.UseCase.ProcessZipCode(input.CEP)
	if err != nil {
		fmt.Printf("error processing zip code: %v\n", err)
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(weather); err != nil {
		fmt.Printf("error encoding response: %v\n", err)
		http.Error(w, "error generating response", http.StatusInternalServerError)
		return
	}
}

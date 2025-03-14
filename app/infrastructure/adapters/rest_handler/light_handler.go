package rest_handler

import (
	"encoding/json"
	"net/http"

	"github.com/SUJEY/PUBLICADOR/app/application/usecase"
	"github.com/SUJEY/PUBLICADOR/app/infrastructure/adapters/simulator"
)

type LightHandler struct {
	useCase *usecase.LightUseCase
}

func NewLightHandler(useCase *usecase.LightUseCase) *LightHandler {
	return &LightHandler{useCase: useCase}
}

func (h *LightHandler) HandleLight(w http.ResponseWriter, r *http.Request) {
	luminosity := simulator.SimulateLight()

	if err := h.useCase.Execute(luminosity); err != nil {
		http.Error(w, "Error processing light sensor", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"luminosity": luminosity,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

package rest_handler

import (
	"encoding/json"
	"net/http"

	"github.com/SUJEY/PUBLICADOR/app/application/usecase"
	"github.com/SUJEY/PUBLICADOR/app/infrastructure/adapters/simulator"
)

type SmokeHandler struct {
	useCase *usecase.SmokeUseCase
}

func NewSmokeHandler(useCase *usecase.SmokeUseCase) *SmokeHandler {
	return &SmokeHandler{useCase: useCase}
}

func (h *SmokeHandler) HandleSmoke(w http.ResponseWriter, r *http.Request) {
	smokeLevel, alarm := simulator.SimulateSmoke()

	if err := h.useCase.Execute(smokeLevel, alarm); err != nil {
		http.Error(w, "Error processing smoke sensor", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"smoke_level": smokeLevel,
		"alarm":       alarm,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

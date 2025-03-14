package rest_handler

import (
	"encoding/json"
	"net/http"

	"github.com/SUJEY/PUBLICADOR/app/application/usecase"
	"github.com/SUJEY/PUBLICADOR/app/infrastructure/adapters/simulator"
)

type MotionHandler struct {
	useCase *usecase.MotionUseCase
}

func NewMotionHandler(useCase *usecase.MotionUseCase) *MotionHandler {
	return &MotionHandler{useCase: useCase}
}

func (h *MotionHandler) HandleMotion(w http.ResponseWriter, r *http.Request) {
	motionDetected, intensity := simulator.SimulateMotion()

	if err := h.useCase.Execute(motionDetected, intensity); err != nil {
		http.Error(w, "Error processing motion sensor", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"motion_detected": motionDetected,
		"intensity":       intensity,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

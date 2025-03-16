package rest_handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/SUJEY/PUBLICADOR/app/application/usecase"
	"github.com/SUJEY/PUBLICADOR/app/domain"
)

type MotionHandler struct {
	useCase *usecase.MotionUseCase
}

func NewMotionHandler(useCase *usecase.MotionUseCase) *MotionHandler {
	return &MotionHandler{useCase: useCase}
}

func (h *MotionHandler) HandleMotion(w http.ResponseWriter, r *http.Request) {

	var sensor domain.MotionSensor
	if err := json.NewDecoder(r.Body).Decode(&sensor); err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}


	if sensor.Timestamp.IsZero() {
		sensor.Timestamp = time.Now()
	}

	
	if err := h.useCase.Execute(sensor.MotionDetected, sensor.Intensity); err != nil {
		http.Error(w, "Error processing motion sensor", http.StatusInternalServerError)
		return
	}

	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sensor)
}

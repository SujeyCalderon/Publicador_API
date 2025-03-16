package rest_handler

import (
	"encoding/json"
	"net/http"

	"github.com/SUJEY/PUBLICADOR/app/application/usecase"
	"github.com/SUJEY/PUBLICADOR/app/domain"
)

type DoorHandler struct {
	useCase *usecase.DoorUseCase
}

func NewDoorHandler(useCase *usecase.DoorUseCase) *DoorHandler {
	return &DoorHandler{useCase: useCase}
}

func (h *DoorHandler) HandleDoor(w http.ResponseWriter, r *http.Request) {

	var sensor domain.DoorSensor
	if err := json.NewDecoder(r.Body).Decode(&sensor); err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}


	if err := h.useCase.Execute(sensor.IsOpen); err != nil {
		http.Error(w, "Error processing door sensor", http.StatusInternalServerError)
		return
	}


	response := map[string]interface{}{
		"is_open":   sensor.IsOpen,     
		"deviceToken": sensor.DeviceToken, 
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
package rest_handler

import (
	"encoding/json"
	"net/http"

	"github.com/SUJEY/PUBLICADOR/app/application/usecase"
	"github.com/SUJEY/PUBLICADOR/app/domain"
)

type LightHandler struct {
	useCase *usecase.LightUseCase
}

func NewLightHandler(useCase *usecase.LightUseCase) *LightHandler {
	return &LightHandler{useCase: useCase}
}

func (h *LightHandler) HandleLight(w http.ResponseWriter, r *http.Request) {

	var sensor domain.LightSensor
	if err := json.NewDecoder(r.Body).Decode(&sensor); err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	
	if err := h.useCase.Execute(sensor.Luminosity); err != nil {
		http.Error(w, "Error processing light sensor", http.StatusInternalServerError)
		return
	}


	response := map[string]interface{}{
		"luminosity": sensor.Luminosity,  
		"deviceToken": sensor.DeviceToken, 
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
package rest_handler

import (
	"encoding/json"
	"net/http"

	"github.com/SUJEY/PUBLICADOR/app/application/usecase"
	"github.com/SUJEY/PUBLICADOR/app/domain"
)

type SmokeHandler struct {
	useCase *usecase.SmokeUseCase
}

func NewSmokeHandler(useCase *usecase.SmokeUseCase) *SmokeHandler {
	return &SmokeHandler{useCase: useCase}
}

func (h *SmokeHandler) HandleSmoke(w http.ResponseWriter, r *http.Request) {
   
    var sensor domain.SmokeSensor
    if err := json.NewDecoder(r.Body).Decode(&sensor); err != nil {
        http.Error(w, "Error decoding request body", http.StatusBadRequest)
        return
    }

  
    if err := h.useCase.Execute(sensor.SmokeLevel, sensor.Alarm); err != nil {
        http.Error(w, "Error processing smoke sensor", http.StatusInternalServerError)
        return
    }

   
    response := map[string]interface{}{
        "smoke_level": sensor.SmokeLevel, 
        "alarm":       sensor.Alarm,      
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

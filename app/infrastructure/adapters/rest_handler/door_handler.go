package rest_handler

import (
	"encoding/json"
	"net/http"

	"github.com/SUJEY/PUBLICADOR/app/application/usecase"
	"github.com/SUJEY/PUBLICADOR/app/infrastructure/adapters/simulator"
)

type DoorHandler struct {
	useCase *usecase.DoorUseCase
}

func NewDoorHandler(useCase *usecase.DoorUseCase) *DoorHandler {
	return &DoorHandler{useCase: useCase}
}

func (h *DoorHandler) HandleDoor(w http.ResponseWriter, r *http.Request) {
	isOpen := simulator.SimulateDoor()

	if err := h.useCase.Execute(isOpen); err != nil {
		http.Error(w, "Error processing door sensor", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"is_open": isOpen,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

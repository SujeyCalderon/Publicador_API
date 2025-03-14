package usecase

import "github.com/SUJEY/PUBLICADOR/app/application/services"

type DoorUseCase struct {
	service services.DoorService
}

func NewDoorUseCase(service services.DoorService) *DoorUseCase {
	return &DoorUseCase{service: service}
}

func (uc *DoorUseCase) Execute(isOpen bool) error {
	return uc.service.ProcessDoor(isOpen)
}

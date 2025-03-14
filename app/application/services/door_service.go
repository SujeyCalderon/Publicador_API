package services

import (
	"time"

	"github.com/SUJEY/PUBLICADOR/app/application/repositories"
	"github.com/SUJEY/PUBLICADOR/app/domain"
)

type DoorService interface {
	ProcessDoor(isOpen bool) error
}

type DoorServiceImpl struct {
	repo repositories.DoorSensorRepository
}

func NewDoorService(repo repositories.DoorSensorRepository) DoorService {
	return &DoorServiceImpl{repo: repo}
}

func (s *DoorServiceImpl) ProcessDoor(isOpen bool) error {
	sensor := domain.DoorSensor{
		ID:        "door1",              // Puedes generar un ID único si lo requieres.
		Timestamp: time.Now(),     // Timestamp en segundos.
		IsOpen:    isOpen,
	}
	return s.repo.Save(sensor)
}

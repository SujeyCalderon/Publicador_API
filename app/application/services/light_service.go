package services

import (
	"time"
	"github.com/SUJEY/PUBLICADOR/app/application/repositories"
	"github.com/SUJEY/PUBLICADOR/app/domain"
)

type LightService interface {
	ProcessLight(luminosity float64) error
}

type LightServiceImpl struct {
	repo repositories.LightSensorRepository
}

func NewLightService(repo repositories.LightSensorRepository) LightService {
	return &LightServiceImpl{repo: repo}
}

func (s *LightServiceImpl) ProcessLight(luminosity float64) error {
	sensor := domain.LightSensor{
		ID:         "light1",
		Timestamp: time.Now(),
		Luminosity: luminosity,
	}
	return s.repo.Save(sensor)
}

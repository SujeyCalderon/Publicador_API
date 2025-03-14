package services

import (
	"time"

	"github.com/SUJEY/PUBLICADOR/app/application/repositories"
	"github.com/SUJEY/PUBLICADOR/app/domain"
)

type SmokeService interface {
	ProcessSmoke(smokeLevel float64, alarm bool) error
}

type SmokeServiceImpl struct {
	repo repositories.SmokeSensorRepository
}

func NewSmokeService(repo repositories.SmokeSensorRepository) SmokeService {
	return &SmokeServiceImpl{repo: repo}
}

func (s *SmokeServiceImpl) ProcessSmoke(smokeLevel float64, alarm bool) error {
	sensor := domain.SmokeSensor{
		ID:         "smoke1",
		Timestamp: time.Now(),
		SmokeLevel: smokeLevel,
		Alarm:      alarm,
	}
	return s.repo.Save(sensor)
}

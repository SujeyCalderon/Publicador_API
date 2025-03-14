package services

import (
	"time"

	"github.com/SUJEY/PUBLICADOR/app/application/repositories"
	"github.com/SUJEY/PUBLICADOR/app/domain"
)

type MotionService interface {
	ProcessMotion(motionDetected bool, intensity float64) error
}

type MotionServiceImpl struct {
	repo repositories.MotionSensorRepository
}

func NewMotionService(repo repositories.MotionSensorRepository) MotionService {
	return &MotionServiceImpl{repo: repo}
}

func (s *MotionServiceImpl) ProcessMotion(motionDetected bool, intensity float64) error {
	sensor := domain.MotionSensor{
		ID:             "motion1",
		Timestamp: time.Now(),
		MotionDetected: motionDetected,
		Intensity:      intensity,
	}
	return s.repo.Save(sensor)
}

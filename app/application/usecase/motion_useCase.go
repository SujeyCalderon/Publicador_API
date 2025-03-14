package usecase

import "github.com/SUJEY/PUBLICADOR/app/application/services"

type MotionUseCase struct {
	service services.MotionService
}

func NewMotionUseCase(service services.MotionService) *MotionUseCase {
	return &MotionUseCase{service: service}
}

func (uc *MotionUseCase) Execute(motionDetected bool, intensity float64) error {
	return uc.service.ProcessMotion(motionDetected, intensity)
}

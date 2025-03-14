package usecase

import "github.com/SUJEY/PUBLICADOR/app/application/services"

type SmokeUseCase struct {
	service services.SmokeService
}

func NewSmokeUseCase(service services.SmokeService) *SmokeUseCase {
	return &SmokeUseCase{service: service}
}

func (uc *SmokeUseCase) Execute(smokeLevel float64, alarm bool) error {
	return uc.service.ProcessSmoke(smokeLevel, alarm)
}

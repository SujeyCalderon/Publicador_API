package usecase

import "github.com/SUJEY/PUBLICADOR/app/application/services"

type LightUseCase struct {
	service services.LightService
}

func NewLightUseCase(service services.LightService) *LightUseCase {
	return &LightUseCase{service: service}
}

func (uc *LightUseCase) Execute(luminosity float64) error {
	return uc.service.ProcessLight(luminosity)
}

package usecase

import (
	"fmt"
	"rlservice/internal/domain/models"
	"rlservice/pkg/logger"
)

type ServiceRepository interface {
	GetServiceByToken(token string) *models.Service
	AddNewService(token string) error
	RemoveService(token string) error
	UpdateServiceBalance(token string, newBalance int) error
}

type ServiceUsecase struct {
	log               *logger.Logger
	serviceRepository ServiceRepository // psql
}

func NewServiceUsecase(log *logger.Logger, ser ServiceRepository) *ServiceUsecase {
	return &ServiceUsecase{
		log:               log,
		serviceRepository: ser,
	}
}

func (su *ServiceUsecase) UpdateService(token string, balanceChange int) error {
	const op = "usecase.UpdateService"
	service := su.serviceRepository.GetServiceByToken(token)
	if service == nil {
		su.serviceRepository.AddNewService(token)
	}
	if err := su.serviceRepository.UpdateServiceBalance(token, balanceChange); err != nil {
		su.log.Error(op, fmt.Sprintf("cannon update service balance: %v", err))
		return ErrCannotUpdateService
	}
	return nil
}

func (su *ServiceUsecase) RemoveService(token string) error {
	return su.serviceRepository.RemoveService(token)
}

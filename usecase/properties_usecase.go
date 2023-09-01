package usecase

import (
	"context"
	"time"

	"github.com/Piyawat-T/go-sample-1/domain"
)

type propertiesUsecase struct {
	contextTimeout time.Duration
}

func NewPropertiesUsecase(timeout time.Duration) domain.PropertiesUseCase {
	return &propertiesUsecase{
		contextTimeout: timeout,
	}
}

func (usecase *propertiesUsecase) GetProperties(c context.Context) ([]domain.Properties, error) {
	properties := []domain.Properties{}
	return properties, nil
}

func (usecase *propertiesUsecase) GetByApplicationAndProfile(c context.Context, application string, profile string) ([]domain.Properties, error) {
	properties := []domain.Properties{}
	return properties, nil
}

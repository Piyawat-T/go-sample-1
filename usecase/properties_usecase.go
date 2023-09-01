package usecase

import (
	"context"
	"time"

	"github.com/Piyawat-T/go-sample-1/domain"
	"github.com/gin-gonic/gin"
)

type propertiesUsecase struct {
	propertiesClient domain.PropertiesClient
	contextTimeout   time.Duration
}

func NewPropertiesUsecase(propertiesClient domain.PropertiesClient, timeout time.Duration) domain.PropertiesUseCase {
	return &propertiesUsecase{
		propertiesClient: propertiesClient,
		contextTimeout:   timeout,
	}
}

func (usecase *propertiesUsecase) GetProperties(c *gin.Context) ([]domain.Properties, error) {
	properties := []domain.Properties{}
	return properties, nil
}

func (usecase *propertiesUsecase) GetByApplicationAndProfile(c *gin.Context, application string, profile string) ([]domain.Properties, error) {
	_, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	properties, err := usecase.propertiesClient.GetByApplicationAndProfile(c, application, profile)
	return properties, err
}

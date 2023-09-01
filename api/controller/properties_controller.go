package controller

import (
	"net/http"

	"github.com/Piyawat-T/go-sample-1/bootstrap"
	"github.com/Piyawat-T/go-sample-1/domain"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
)

type PropertiesController struct {
	PropertiesUsecase domain.PropertiesUseCase
	Env               *bootstrap.Env
}

func (controller *PropertiesController) GetProperties(c *gin.Context) {
	properties, err := controller.PropertiesUsecase.GetProperties(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, properties)
}

func (controller *PropertiesController) GetConfiguration(c *gin.Context) {
	application := c.Param("application")
	profile := c.Param("profile")

	ctx := c.Request.Context()
	log := otelzap.L()
	log.DebugContext(ctx, "Start Get Configuration")

	properties, err := controller.PropertiesUsecase.GetByApplicationAndProfile(c, application, profile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, properties)
}

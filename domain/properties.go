package domain

import (
	"github.com/gin-gonic/gin"
)

const (
	CollectionProperties = "properties"
)

type Properties struct {
	Id          uint   `json:"id"`
	Application string `json:"application"`
	Profile     string `json:"profile"`
	Key         string `json:"key"`
	Value       string `json:"value"`
}

type PropertiesRepository interface {
	Fetch(c *gin.Context) ([]Properties, error)
	FetchByApplicationAndProfile(c *gin.Context, application string, profile string) ([]Properties, error)
}

type PropertiesUseCase interface {
	GetProperties(c *gin.Context) ([]Properties, error)
	GetByApplicationAndProfile(c *gin.Context, application string, profile string) ([]Properties, error)
}

type PropertiesClient interface {
	GetByApplicationAndProfile(c *gin.Context, application string, profile string) ([]Properties, error)
}

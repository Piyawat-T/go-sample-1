package client

import (
	"encoding/json"
	"io"

	"github.com/Piyawat-T/go-sample-1/domain"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type propertiesClient struct {
}

func NewPropertiesClient() domain.PropertiesClient {
	return &propertiesClient{}
}

func (client *propertiesClient) GetByApplicationAndProfile(c *gin.Context, application string, profile string) ([]domain.Properties, error) {
	ctx := c.Request.Context()
	log := otelzap.L()
	log.DebugContext(ctx, "Start Get Configuration")

	var properties []domain.Properties
	resp, err := otelhttp.Get(c, "http://localhost:8100/go-centralize-configuration/deposit/default")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)

	log.DebugContext(ctx, resp.Status)
	json.Unmarshal(b, &properties)
	return properties, err
}

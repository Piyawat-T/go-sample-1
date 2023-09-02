package client

import (
	"encoding/json"
	"io"
	"net/http"

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

func (pc *propertiesClient) GetByApplicationAndProfile(c *gin.Context, application string, profile string) ([]domain.Properties, error) {
	ctx := c.Request.Context()
	log := otelzap.L()
	log.DebugContext(ctx, "Start Get Configuration")

	// client := http.Client{
	// 	Transport: otelhttp.NewTransport(
	// 		http.DefaultTransport,
	// 		otelhttp.WithClientTrace(func(ctx context.Context) *httptrace.ClientTrace {
	// 			return otelhttptrace.NewClientTrace(ctx)
	// 		}),
	// 	),
	// }
	// resp, err := client.Get("http://localhost:8100/go-centralize-configuration/deposit/default")

	// clientTrace := otelhttptrace.NewClientTrace(ctx)
	// ctx = httptrace.WithClientTrace(ctx, clientTrace)

	// tracer := otel.Tracer("hello-world-frontend")
	// otel.SetTextMapPropagator(propagation.TraceContext{})
	// ctx, span := tracer.Start(ctx, "client")
	// defer span.End()

	// span.AddEvent("Launching Request to backend")

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8100/go-centralize-configuration/deposit/default", nil)
	if err != nil {
		return nil, err
	}
	// resp, err := http.DefaultClient.Do(req)
	httpclient := &http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}
	resp, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)

	log.DebugContext(ctx, resp.Status)
	var properties []domain.Properties
	json.Unmarshal(b, &properties)
	return properties, err
}

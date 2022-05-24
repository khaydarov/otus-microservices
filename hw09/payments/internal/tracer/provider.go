package tracer

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
)

type ProviderConfig struct {
	JaegerEndpoint 	string
	ServiceName 	string
	Environment		string
}

type Provider struct {
	provider *tracesdk.TracerProvider
}

func (p *Provider) Close(context context.Context) error {
	err := p.provider.Shutdown(context)
	if err != nil {
		return err
	}

	return nil
}

func NewProvider(config ProviderConfig) (Provider, error) {
	exporter, _ := newExporter(config.JaegerEndpoint)

	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exporter),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(config.ServiceName),
			attribute.String("environment", config.Environment),
		)),
	)
	otel.SetTracerProvider(tp)

	return Provider{
		tp,
	}, nil
}

// newExporter returns a console exporter.
func newExporter(url string) (*jaeger.Exporter, error) {
	// Create the Jaeger exporter
	exporter, err := jaeger.New(
		jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)),
	)

	if err != nil {
		return nil, err
	}

	return exporter, nil
}

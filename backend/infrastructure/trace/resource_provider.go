package trace

import (
	"payd/config"

	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.32.0"
)

func getResource(config *config.Config) *resource.Resource {
	res := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String("rms-service"),
		semconv.ServiceVersionKey.String(config.Version()),
	)

	return res
}

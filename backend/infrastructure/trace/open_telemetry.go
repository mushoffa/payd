package trace

import (
	"context"
	"log"

	"payd/config"

	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

type TracerService interface {
	Shutdown(context.Context)
}

type client struct {
	trace *sdktrace.TracerProvider
	meter *sdkmetric.MeterProvider
}

func NewOpenTelemetry(config *config.Config) TracerService {
	ctx := context.Background()

	traceProvider, err := getTraceProvider(ctx, config)
	if err != nil {
		panic(err)
	}

	meterProvider, err := getMeterProvider(ctx, config)
	if err != nil {
		panic(err)
	}

	return &client{
		trace: traceProvider,
		meter: meterProvider,
	}
}

func (otel *client) Shutdown(ctx context.Context) {
	if err := otel.trace.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	if err := otel.meter.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}

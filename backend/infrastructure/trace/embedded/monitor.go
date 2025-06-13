package embedded

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

type Monitor struct {
	attrs  []attribute.KeyValue
	tracer trace.Tracer
	meter  metric.Meter
}

func (m *Monitor) Init(keys ...attribute.KeyValue) {
	m.attrs = append(m.attrs, keys...)
	m.tracer = otel.Tracer("")
}

func (m *Monitor) Trace(ctx context.Context, spanName string, kv ...attribute.KeyValue) (context.Context, trace.Span) {
	childCtx, span := m.tracer.Start(ctx, spanName, trace.WithAttributes(m.attrs...))
	span.SetAttributes(kv...)

	return childCtx, span
}

func (m Monitor) Attribute(key, value string) attribute.KeyValue {
	return attribute.String(key, value)
}

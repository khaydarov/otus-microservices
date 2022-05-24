package tracer

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func NewSpan(context context.Context, name string) (context.Context, trace.Span) {
	return otel.Tracer("").Start(context, name)
}

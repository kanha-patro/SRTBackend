package observability

import (
	"context"
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

// Tracer wraps OpenTelemetry tracer.
type Tracer struct {
	tracer trace.Tracer
}

// NewTracer creates a new Tracer instance.
func NewTracer(serviceName string) *Tracer {
	return &Tracer{
		tracer: otel.Tracer(serviceName),
	}
}

// StartSpan starts a new span.
func (t *Tracer) StartSpan(ctx context.Context, name string) (context.Context, trace.Span) {
	return t.tracer.Start(ctx, name)
}

// TraceHTTPMiddleware instruments HTTP handlers using otelhttp.
func (t *Tracer) TraceHTTPMiddleware(next http.Handler) http.Handler {
	return otelhttp.NewHandler(next, "http_request")
}

// NewHTTPClient returns an instrumented HTTP client.
func NewHTTPClient() *http.Client {
	return &http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}
}
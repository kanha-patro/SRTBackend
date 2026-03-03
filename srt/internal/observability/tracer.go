package observability

import (
	"go.opentelemetry.io/contrib/instrumentation/net/http/httptrace"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"net/http"
)

// Tracer is a wrapper around OpenTelemetry's tracer.
type Tracer struct {
	tracer trace.Tracer
}

// NewTracer creates a new Tracer instance.
func NewTracer(serviceName string) *Tracer {
	tracer := otel.Tracer(serviceName)
	return &Tracer{tracer: tracer}
}

// StartSpan starts a new span for tracing.
func (t *Tracer) StartSpan(ctx context.Context, name string) (context.Context, trace.Span) {
	return t.tracer.Start(ctx, name)
}

// TraceHTTPMiddleware is a middleware for tracing HTTP requests.
func (t *Tracer) TraceHTTPMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, span := t.tracer.Start(r.Context(), "http_request")
		defer span.End()

		// Add HTTP trace to the span
		httptrace.Inject(ctx, r)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
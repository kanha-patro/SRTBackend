package observability

import (
	"context"
	"fmt"
	"sync"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

var (
	meter metric.Meter

	tripCount       metric.Int64Counter
	locationUpdates metric.Int64Counter
	staleTrips      metric.Int64Counter
	driverDropouts  metric.Int64Counter

	initOnce sync.Once
	initErr  error
)

// InitMetrics initializes all OpenTelemetry metrics.
// Call this AFTER your OpenTelemetry MeterProvider is configured.
func InitMetrics() error {
	initOnce.Do(func() {
		meter = otel.GetMeterProvider().Meter("shuttle-tracking")

		var err error

		tripCount, err = meter.Int64Counter(
			"active_trips",
			metric.WithDescription("Count of active trips"),
		)
		if err != nil {
			initErr = fmt.Errorf("failed to create active_trips counter: %w", err)
			return
		}

		locationUpdates, err = meter.Int64Counter(
			"location_updates",
			metric.WithDescription("Count of location updates received"),
		)
		if err != nil {
			initErr = fmt.Errorf("failed to create location_updates counter: %w", err)
			return
		}

		staleTrips, err = meter.Int64Counter(
			"stale_trips",
			metric.WithDescription("Count of trips that have become stale"),
		)
		if err != nil {
			initErr = fmt.Errorf("failed to create stale_trips counter: %w", err)
			return
		}

		driverDropouts, err = meter.Int64Counter(
			"driver_dropouts",
			metric.WithDescription("Count of driver dropouts"),
		)
		if err != nil {
			initErr = fmt.Errorf("failed to create driver_dropouts counter: %w", err)
			return
		}
	})

	return initErr
}

func RecordTripCount(ctx context.Context, orgID string) {
	if tripCount == nil {
		return
	}
	tripCount.Add(ctx, 1,
		metric.WithAttributes(attribute.String("org_id", orgID)),
	)
}

func RecordLocationUpdate(ctx context.Context, orgID string) {
	if locationUpdates == nil {
		return
	}
	locationUpdates.Add(ctx, 1,
		metric.WithAttributes(attribute.String("org_id", orgID)),
	)
}

func RecordStaleTrip(ctx context.Context, orgID string) {
	if staleTrips == nil {
		return
	}
	staleTrips.Add(ctx, 1,
		metric.WithAttributes(attribute.String("org_id", orgID)),
	)
}

func RecordDriverDropout(ctx context.Context, orgID string) {
	if driverDropouts == nil {
		return
	}
	driverDropouts.Add(ctx, 1,
		metric.WithAttributes(attribute.String("org_id", orgID)),
	)
}
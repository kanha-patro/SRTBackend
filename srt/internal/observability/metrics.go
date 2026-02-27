package observability

import (
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/attribute"
)

var (
	// Metrics
	tripCount       metric.Int64Counter
	locationUpdates metric.Int64Counter
	staleTrips      metric.Int64Counter
	driverDropouts  metric.Int64Counter
)

func init() {
	meter := global.Meter("shuttle-tracking")

	var err error
	tripCount, err = meter.Int64Counter("active_trips", metric.WithDescription("Count of active trips"))
	if err != nil {
		panic(err)
	}

	locationUpdates, err = meter.Int64Counter("location_updates", metric.WithDescription("Count of location updates received"))
	if err != nil {
		panic(err)
	}

	staleTrips, err = meter.Int64Counter("stale_trips", metric.WithDescription("Count of trips that have become stale"))
	if err != nil {
		panic(err)
	}

	driverDropouts, err = meter.Int64Counter("driver_dropouts", metric.WithDescription("Count of driver dropouts"))
	if err != nil {
		panic(err)
	}
}

// RecordTripCount increments the active trip counter.
func RecordTripCount(orgID string) {
	tripCount.Add(context.Background(), 1, attribute.String("org_id", orgID))
}

// RecordLocationUpdate increments the location update counter.
func RecordLocationUpdate(orgID string) {
	locationUpdates.Add(context.Background(), 1, attribute.String("org_id", orgID))
}

// RecordStaleTrip increments the stale trip counter.
func RecordStaleTrip(orgID string) {
	staleTrips.Add(context.Background(), 1, attribute.String("org_id", orgID))
}

// RecordDriverDropout increments the driver dropout counter.
func RecordDriverDropout(orgID string) {
	driverDropouts.Add(context.Background(), 1, attribute.String("org_id", orgID))
}
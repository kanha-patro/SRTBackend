package repository

import (
	"context"
	"time"

	"github.com/akpatri/srt/internal/domain"
)

// OrgRepository defines organization DB operations.
type OrgRepository interface {
	Create(org *domain.Org) error
	GetByID(id string) (*domain.Org, error)
	Update(org *domain.Org) error
	Delete(id string) error
	GetAll() ([]domain.Org, error)
}

// RouteRepository defines route DB operations.
type RouteRepository interface {
	CreateRoute(route *domain.Route) error
	GetRouteByID(id string) (*domain.Route, error)
	UpdateRoute(route *domain.Route) error
	DeleteRoute(id string) error
	GetAllRoutes(orgID string) ([]domain.Route, error)
}

// StopRepository defines stop DB operations.
type StopRepository interface {
	Create(stop *domain.Stop) error
	Update(stop *domain.Stop) error
	Delete(stopID string) error
	FindByID(stopID string) (*domain.Stop, error)
	FindAllByRouteID(routeID string) ([]domain.Stop, error)
	// FindAll returns all stops in the system.
	FindAll(ctx context.Context) ([]domain.Stop, error)
	// FindStopsWithinRadius returns stops within the radius (meters) from provided lat/lon.
	FindStopsWithinRadius(ctx context.Context, latitude, longitude, radius float64) ([]domain.Stop, error)
}

// TripRepository defines trip DB operations.
type TripRepository interface {
	CreateTrip(ctx context.Context, trip *domain.Trip) error
	GetTripByID(ctx context.Context, id string) (*domain.Trip, error)
	UpdateTrip(ctx context.Context, trip *domain.Trip) error
	DeleteTrip(ctx context.Context, id string) error
	GetActiveTripsByOrgID(ctx context.Context, orgID string) ([]*domain.Trip, error)
	GetStaleTrips(ctx context.Context, threshold time.Duration) ([]*domain.Trip, error)
}

// DriverRepository defines driver DB operations.
type DriverRepository interface {
	Create(driver *domain.Driver) error
	Update(driver *domain.Driver) error
	Delete(driverID string) error
	FindByID(driverID string) (*domain.Driver, error)
	FindAll() ([]domain.Driver, error)
}

// LocationRepository defines location DB operations.
type LocationRepository interface {
	SaveLocation(ctx context.Context, location *domain.Location) error
	GetLatestLocationByTripID(ctx context.Context, tripID string) (*domain.Location, error)
	GetLocationsByTripID(ctx context.Context, tripID string, startTime, endTime time.Time) ([]domain.Location, error)
	// FindActiveLocations returns latest locations for all active trips in an org.
	FindActiveLocations(ctx context.Context, orgID string) ([]domain.Location, error)
}

// AuditRepository defines audit log DB operations.
type AuditRepository interface {
	CreateAuditLog(ctx context.Context, log *domain.Audit) error
	GetAuditLogs(ctx context.Context, orgID string, limit, offset int) ([]domain.Audit, error)
}

// OTPRepository stores transient OTPs for driver sessions.
type OTPEntry struct {
	Code   string
	Expiry time.Time
}

type OTPRepository interface {
	StoreOTP(orgCode, routeCode, driverCode, code string, expiry time.Time) error
	GetOTP(orgCode, routeCode, driverCode string) (*OTPEntry, error)
}

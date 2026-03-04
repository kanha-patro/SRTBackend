package service

import (
	"context"
	"time"

	"github.com/akpatri/srt/internal/domain"
)

// OrgService defines organization business logic.
type OrgService interface {
	RegisterOrg(org *domain.Org) error
	ApproveOrg(orgID string) error
	SuspendOrg(orgID string) error
	GetOrg(orgID string) (*domain.Org, error)
	GetActiveTrips() ([]*domain.Trip, error)
	ForceStopTrip(tripID string) error
	RevokeOTPSession(sessionID string) error
	GetActiveOrgs() ([]domain.Org, error)
	UpdateOrg(orgID string, org *domain.Org) error
}

// RouteService defines route business logic.
type RouteService interface {
	CreateRoute(route *domain.Route) error
	UpdateRoute(route *domain.Route) error
	DeleteRoute(routeID string) error
	GetRoute(routeID string) (*domain.Route, error)
	GetAllRoutes(orgID string) ([]domain.Route, error)
}

// TripService defines trip lifecycle operations.
type TripService interface {
	StartTrip(ctx context.Context, trip *domain.Trip) error
	UpdateLocation(ctx context.Context, tripID string, location *domain.Location) error
	EndTrip(ctx context.Context, tripID string) error
	AutoEndStaleTrips(ctx context.Context, threshold time.Duration) error
	GetActiveTrips(ctx context.Context) ([]*domain.Trip, error)
}

// DriverService defines driver-facing operations (OTP start, location updates).
type DriverService interface {
	StartTrip(ctx context.Context, driverID string, otp string) (*domain.Trip, error)
	UpdateLocation(ctx context.Context, tripID string, location domain.Location) error
	EndTrip(ctx context.Context, tripID string) error
}

// LocationService defines location-related operations.
type LocationService interface {
	UpdateLocation(ctx context.Context, tripID string, location domain.Location) error
	GetActiveLocations(ctx context.Context, orgID string) ([]domain.Location, error)
}

// UserService defines public-facing APIs for users.
type UserService interface {
	GetActiveShuttles(orgCode, routeCode, nearbyLocation string) ([]domain.Trip, error)
}

// GeoService defines geo utilities exposed to handlers.
type GeoService interface {
	SnapLocation(lat string, lon string) (interface{}, error)
	SearchNearbyStops(lat string, lon string) ([]domain.Location, error)
}

// OTPService defines OTP generation/validation.
type OTPService interface {
	GenerateOTP(orgCode, routeCode, driverCode string) (string, error)
	ValidateOTP(orgCode, routeCode, driverCode, otpCode string) (bool, error)
}

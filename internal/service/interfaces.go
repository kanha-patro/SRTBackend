package service

import "github.com/akpatri/srt/internal/domain"

// OrgService defines the interface for organization-related business logic.
type OrgService interface {
    CreateOrg(org *domain.Org) error
    ApproveOrg(orgID string) error
    SuspendOrg(orgID string) error
    GetOrg(orgID string) (*domain.Org, error)
}

// RouteService defines the interface for route-related business logic.
type RouteService interface {
    CreateRoute(route *domain.Route) error
    EditRoute(routeID string, route *domain.Route) error
    GetRoute(routeID string) (*domain.Route, error)
    DeleteRoute(routeID string) error
}

// TripService defines the interface for trip-related business logic.
type TripService interface {
    StartTrip(trip *domain.Trip) error
    EndTrip(tripID string) error
    GetActiveTrips(orgID string) ([]*domain.Trip, error)
}

// DriverService defines the interface for driver-related business logic.
type DriverService interface {
    AssignDriver(driver *domain.Driver) error
    UnassignDriver(driverID string) error
    GetDriver(driverID string) (*domain.Driver, error)
}

// LocationService defines the interface for location-related business logic.
type LocationService interface {
    UpdateLocation(location *domain.Location) error
    GetLastKnownLocation(driverID string) (*domain.Location, error)
}

// OTPService defines the interface for OTP-related business logic.
type OTPService interface {
    GenerateOTP(orgCode, routeCode, driverCode string) (string, error)
    ValidateOTP(otp string) (bool, error)
}
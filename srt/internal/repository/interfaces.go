package repository

import "github.com/akpatri/srt/internal/domain"

// OrgRepository defines the interface for organization-related database operations.
type OrgRepository interface {
    Create(org *domain.Org) error
    GetByID(id string) (*domain.Org, error)
    Update(org *domain.Org) error
    Delete(id string) error
    GetAll() ([]*domain.Org, error)
}

// RouteRepository defines the interface for route-related database operations.
type RouteRepository interface {
    Create(route *domain.Route) error
    GetByID(id string) (*domain.Route, error)
    Update(route *domain.Route) error
    Delete(id string) error
    GetAllByOrgID(orgID string) ([]*domain.Route, error)
}

// StopRepository defines the interface for stop-related database operations.
type StopRepository interface {
    Create(stop *domain.Stop) error
    GetByID(id string) (*domain.Stop, error)
    Update(stop *domain.Stop) error
    Delete(id string) error
    GetAllByRouteID(routeID string) ([]*domain.Stop, error)
}

// TripRepository defines the interface for trip-related database operations.
type TripRepository interface {
    Create(trip *domain.Trip) error
    GetByID(id string) (*domain.Trip, error)
    Update(trip *domain.Trip) error
    Delete(id string) error
    GetActiveByDriverID(driverID string) (*domain.Trip, error)
}

// DriverRepository defines the interface for driver-related database operations.
type DriverRepository interface {
    Create(driver *domain.Driver) error
    GetByID(id string) (*domain.Driver, error)
    Update(driver *domain.Driver) error
    Delete(id string) error
    GetAllByOrgID(orgID string) ([]*domain.Driver, error)
}

// LocationRepository defines the interface for location-related database operations.
type LocationRepository interface {
    Create(location *domain.Location) error
    GetByID(id string) (*domain.Location, error)
    Update(location *domain.Location) error
    Delete(id string) error
}

// AuditRepository defines the interface for audit log-related database operations.
type AuditRepository interface {
    Create(audit *domain.Audit) error
    GetAll() ([]*domain.Audit, error)
}
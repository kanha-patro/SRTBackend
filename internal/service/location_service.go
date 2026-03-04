package service

import (
	"context"
	"time"

	"github.com/akpatri/srt/internal/domain"
	"github.com/akpatri/srt/internal/repository"
	"github.com/akpatri/srt/pkg/errors"
)

type LocationService interface {
	UpdateLocation(ctx context.Context, tripID string, location domain.Location) error
	GetActiveLocations(ctx context.Context, orgID string) ([]domain.Location, error)
}

type locationService struct {
	locationRepo repository.LocationRepository
}

func NewLocationService(locationRepo repository.LocationRepository) LocationService {
	return &locationService{
		locationRepo: locationRepo,
	}
}

func (s *locationService) UpdateLocation(ctx context.Context, tripID string, location domain.Location) error {
	// Validate location data
	if err := validateLocation(location); err != nil {
		return err
	}

	// Set timestamp
	location.Timestamp = time.Now().UTC()

	// Save location to repository
	if err := s.locationRepo.SaveLocation(ctx, tripID, location); err != nil {
		return errors.NewInternalServerError("failed to update location")
	}

	return nil
}

func (s *locationService) GetActiveLocations(ctx context.Context, orgID string) ([]domain.Location, error) {
	locations, err := s.locationRepo.FindActiveLocations(ctx, orgID)
	if err != nil {
		return nil, errors.NewInternalServerError("failed to retrieve active locations")
	}
	return locations, nil
}

func validateLocation(location domain.Location) error {
	// Implement validation logic for location
	if location.Latitude == 0 || location.Longitude == 0 {
		return errors.NewBadRequestError("invalid location coordinates")
	}
	// Additional validation rules can be added here
	return nil
}
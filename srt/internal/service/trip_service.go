package service

import (
	"context"
	"errors"
	"time"

	"github.com/akpatri/srt/internal/domain"
	"github.com/akpatri/srt/internal/repository"
	"github.com/akpatri/srt/internal/event"
	"github.com/akpatri/srt/internal/observability"
)

type TripService struct {
	tripRepo       repository.TripRepository
	locationRepo   repository.LocationRepository
	eventPublisher  event.Publisher
	logger         observability.Logger
}

func NewTripService(tripRepo repository.TripRepository, locationRepo repository.LocationRepository, eventPublisher event.Publisher, logger observability.Logger) *TripService {
	return &TripService{
		tripRepo:      tripRepo,
		locationRepo:  locationRepo,
		eventPublisher: eventPublisher,
		logger:        logger,
	}
}

func (s *TripService) StartTrip(ctx context.Context, trip *domain.Trip) error {
	if trip == nil {
		return errors.New("trip cannot be nil")
	}

	// Validate trip details
	if err := trip.Validate(); err != nil {
		return err
	}

	// Save trip to the repository
	if err := s.tripRepo.Create(ctx, trip); err != nil {
		return err
	}

	// Publish TripStarted event
	s.eventPublisher.Publish(event.TripStarted{TripID: trip.ID})

	return nil
}

func (s *TripService) UpdateLocation(ctx context.Context, tripID string, location *domain.Location) error {
	if location == nil {
		return errors.New("location cannot be nil")
	}

	// Update location in the repository
	if err := s.locationRepo.Update(ctx, tripID, location); err != nil {
		return err
	}

	// Publish LocationUpdated event
	s.eventPublisher.Publish(event.LocationUpdated{TripID: tripID, Location: *location})

	return nil
}

func (s *TripService) EndTrip(ctx context.Context, tripID string) error {
	trip, err := s.tripRepo.GetByID(ctx, tripID)
	if err != nil {
		return err
	}

	if trip == nil {
		return errors.New("trip not found")
	}

	// End the trip
	trip.EndedAt = time.Now()
	if err := s.tripRepo.Update(ctx, trip); err != nil {
		return err
	}

	// Publish TripEnded event
	s.eventPublisher.Publish(event.TripEnded{TripID: tripID})

	return nil
}

func (s *TripService) AutoEndStaleTrips(ctx context.Context, threshold time.Duration) error {
	staleTrips, err := s.tripRepo.GetStaleTrips(ctx, threshold)
	if err != nil {
		return err
	}

	for _, trip := range staleTrips {
		if err := s.EndTrip(ctx, trip.ID); err != nil {
			s.logger.Error("failed to end stale trip", err)
		}
	}

	return nil
}
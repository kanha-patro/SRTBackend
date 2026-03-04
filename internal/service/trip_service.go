package service

import (
	"context"
	"time"

	"github.com/akpatri/srt/internal/domain"
	"github.com/akpatri/srt/internal/event"
	"github.com/akpatri/srt/internal/observability"
	"github.com/akpatri/srt/internal/repository"
	pkgerrors "github.com/akpatri/srt/pkg/errors"
	"go.uber.org/zap"
)

type tripServiceImpl struct {
	tripRepo       repository.TripRepository
	locationRepo   repository.LocationRepository
	eventPublisher event.Publisher
	logger         observability.Logger
}

func NewTripService(tripRepo repository.TripRepository, locationRepo repository.LocationRepository, eventPublisher event.Publisher, logger observability.Logger) TripService {
	return &tripServiceImpl{
		tripRepo:       tripRepo,
		locationRepo:   locationRepo,
		eventPublisher: eventPublisher,
		logger:         logger,
	}
}

func (s tripServiceImpl) StartTrip(ctx context.Context, trip *domain.Trip) error {
	if trip == nil {
		return pkgerrors.NewBadRequestError("trip cannot be nil")
	}

	// Basic validation: require RouteID and DriverID
	if trip.RouteID == "" || trip.DriverID == "" {
		return pkgerrors.NewBadRequestError("route_id and driver_id are required")
	}

	trip.State = domain.CREATED
	trip.StartTime = time.Now()

	if err := s.tripRepo.CreateTrip(ctx, trip); err != nil {
		return err
	}

	_ = s.eventPublisher.Publish("trip.started", event.TripStarted{Event: event.Event{Timestamp: time.Now(), Type: "TripStarted"}, TripID: trip.ID})
	return nil
}

func (s tripServiceImpl) UpdateLocation(ctx context.Context, tripID string, location *domain.Location) error {
	if location == nil {
		return pkgerrors.NewBadRequestError("location cannot be nil")
	}

	// set trip association and timestamp
	location.TripID = tripID
	location.Timestamp = time.Now().UTC()

	if err := s.locationRepo.SaveLocation(ctx, location); err != nil {
		return err
	}

	_ = s.eventPublisher.Publish("location.updated", event.LocationUpdated{Event: event.Event{Timestamp: time.Now(), Type: "LocationUpdated"}, TripID: tripID, Latitude: location.Latitude, Longitude: location.Longitude})
	return nil
}

func (s tripServiceImpl) EndTrip(ctx context.Context, tripID string) error {
	trip, err := s.tripRepo.GetTripByID(ctx, tripID)
	if err != nil {
		return err
	}
	if trip == nil {
		return pkgerrors.NewNotFoundError("trip not found")
	}
	trip.EndTime = time.Now()
	trip.State = domain.ENDED
	if err := s.tripRepo.UpdateTrip(ctx, trip); err != nil {
		return err
	}
	_ = s.eventPublisher.Publish("trip.ended", event.TripAutoEnded{Event: event.Event{Timestamp: time.Now(), Type: "TripEnded"}, TripID: tripID})
	return nil
}

func (s tripServiceImpl) AutoEndStaleTrips(ctx context.Context, threshold time.Duration) error {
	staleTrips, err := s.tripRepo.GetStaleTrips(ctx, threshold)
	if err != nil {
		return err
	}
	for _, tr := range staleTrips {
		if err := s.EndTrip(ctx, tr.ID); err != nil {
			s.logger.Error("failed to end stale trip", zap.Any("error", err))
		}
	}
	return nil
}

func (s tripServiceImpl) GetActiveTrips(ctx context.Context) ([]*domain.Trip, error) {
	// Delegate to repository. Org scoping can be added later via context metadata.
	return s.tripRepo.GetActiveTripsByOrgID(ctx, "")
}

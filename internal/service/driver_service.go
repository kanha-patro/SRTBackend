package service

import (
	"context"
	"time"

	"github.com/akpatri/srt/internal/domain"
	"github.com/akpatri/srt/internal/event"
	"github.com/akpatri/srt/internal/repository"
	"github.com/akpatri/srt/pkg/errors"
)

type driverService struct {
	tripRepo     repository.TripRepository
	driverRepo   repository.DriverRepository
	locationRepo repository.LocationRepository
	otpSvc       OTPService
	publisher    event.Publisher
}

func NewDriverService(tripRepo repository.TripRepository, driverRepo repository.DriverRepository, locationRepo repository.LocationRepository, otpSvc OTPService, publisher event.Publisher) DriverService {
	return &driverService{
		tripRepo:     tripRepo,
		driverRepo:   driverRepo,
		locationRepo: locationRepo,
		otpSvc:       otpSvc,
		publisher:    publisher,
	}
}

func (s *driverService) StartTrip(ctx context.Context, driverID string, otp string) (*domain.Trip, error) {
	driver, err := s.driverRepo.FindByID(driverID)
	if err != nil {
		return nil, errors.NewNotFoundError("Driver not found")
	}

	// Validate via OTPService
	ok, err := s.otpSvc.ValidateOTP(driver.OrgID, driver.RouteID, driver.Code, otp)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.NewUnauthorizedError("Invalid OTP")
	}

	trip := &domain.Trip{
		DriverID:  driver.ID,
		State:     domain.STARTED,
		StartTime: time.Now(),
	}

	if err := s.tripRepo.CreateTrip(ctx, trip); err != nil {
		return nil, err
	}

	// Publish trip.started event
	_ = s.publisher.Publish("trip.started", event.TripStarted{Event: event.Event{Timestamp: time.Now(), Type: "TripStarted"}, TripID: trip.ID})

	return trip, nil
}

func (s *driverService) UpdateLocation(ctx context.Context, tripID string, location domain.Location) error {
	trip, err := s.tripRepo.GetTripByID(ctx, tripID)
	if err != nil {
		return errors.NewNotFoundError("Trip not found")
	}

	if trip.State != domain.ACTIVE {
		return errors.NewInvalidStateError("Trip is not active")
	}

	location.TripID = trip.ID
	if err := s.locationRepo.SaveLocation(ctx, &location); err != nil {
		return err
	}

	_ = s.publisher.Publish("location.updated", event.LocationUpdated{Event: event.Event{Timestamp: time.Now(), Type: "LocationUpdated"}, TripID: tripID, Latitude: location.Latitude, Longitude: location.Longitude})

	return nil
}

func (s *driverService) EndTrip(ctx context.Context, tripID string) error {
	trip, err := s.tripRepo.GetTripByID(ctx, tripID)
	if err != nil {
		return errors.NewNotFoundError("Trip not found")
	}

	if trip.State != domain.ACTIVE {
		return errors.NewInvalidStateError("Trip is not active")
	}

	trip.State = domain.ENDED
	if err := s.tripRepo.UpdateTrip(ctx, trip); err != nil {
		return err
	}

	_ = s.publisher.Publish("trip.ended", event.TripAutoEnded{Event: event.Event{Timestamp: time.Now(), Type: "TripEnded"}, TripID: tripID})

	return nil
}

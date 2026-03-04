package service

import (
	"context"
	"time"

	"github.com/akpatri/srt/internal/domain"
	"github.com/akpatri/srt/internal/repository"
	"github.com/akpatri/srt/pkg/errors"
)

type DriverService interface {
	StartTrip(ctx context.Context, driverID string, otp string) (*domain.Trip, error)
	UpdateLocation(ctx context.Context, tripID string, location domain.Location) error
	EndTrip(ctx context.Context, tripID string) error
}

type driverService struct {
	tripRepo    repository.TripRepository
	driverRepo  repository.DriverRepository
	locationRepo repository.LocationRepository
}

func NewDriverService(tripRepo repository.TripRepository, driverRepo repository.DriverRepository, locationRepo repository.LocationRepository) DriverService {
	return &driverService{
		tripRepo:    tripRepo,
		driverRepo:  driverRepo,
		locationRepo: locationRepo,
	}
}

func (s *driverService) StartTrip(ctx context.Context, driverID string, otp string) (*domain.Trip, error) {
	driver, err := s.driverRepo.FindByID(ctx, driverID)
	if err != nil {
		return nil, errors.NewNotFoundError("Driver not found")
	}

	if !driver.ValidateOTP(otp) {
		return nil, errors.NewUnauthorizedError("Invalid OTP")
	}

	trip := &domain.Trip{
		DriverID: driver.ID,
		State:    domain.TripStateCreated,
		StartedAt: time.Now(),
	}

	if err := s.tripRepo.Create(ctx, trip); err != nil {
		return nil, err
	}

	return trip, nil
}

func (s *driverService) UpdateLocation(ctx context.Context, tripID string, location domain.Location) error {
	trip, err := s.tripRepo.FindByID(ctx, tripID)
	if err != nil {
		return errors.NewNotFoundError("Trip not found")
	}

	if trip.State != domain.TripStateActive {
		return errors.NewInvalidStateError("Trip is not active")
	}

	location.TripID = trip.ID
	if err := s.locationRepo.Save(ctx, location); err != nil {
		return err
	}

	return nil
}

func (s *driverService) EndTrip(ctx context.Context, tripID string) error {
	trip, err := s.tripRepo.FindByID(ctx, tripID)
	if err != nil {
		return errors.NewNotFoundError("Trip not found")
	}

	if trip.State != domain.TripStateActive {
		return errors.NewInvalidStateError("Trip is not active")
	}

	trip.State = domain.TripStateEnded
	if err := s.tripRepo.Update(ctx, trip); err != nil {
		return err
	}

	return nil
}
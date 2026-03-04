package trip

import (
	"time"

	"github.com/akpatri/srt/internal/domain"
	"github.com/akpatri/srt/internal/event"
	"github.com/akpatri/srt/internal/repository"
)

type TripLifecycle struct {
	tripRepo      repository.TripRepository
	locationRepo  repository.LocationRepository
	eventPublisher event.Publisher
}

func NewTripLifecycle(tripRepo repository.TripRepository, locationRepo repository.LocationRepository, eventPublisher event.Publisher) *TripLifecycle {
	return &TripLifecycle{
		tripRepo:      tripRepo,
		locationRepo:  locationRepo,
		eventPublisher: eventPublisher,
	}
}

func (tl *TripLifecycle) StartTrip(trip *domain.Trip) error {
	trip.State = domain.TripStateStarted
	trip.StartedAt = time.Now()

	if err := tl.tripRepo.Save(trip); err != nil {
		return err
	}

	tl.eventPublisher.Publish(event.TripStarted{TripID: trip.ID})
	return nil
}

func (tl *TripLifecycle) ActivateTrip(trip *domain.Trip) error {
	if trip.State != domain.TripStateStarted {
		return domain.ErrInvalidTripState
	}

	trip.State = domain.TripStateActive
	trip.ActivatedAt = time.Now()

	if err := tl.tripRepo.Save(trip); err != nil {
		return err
	}

	tl.eventPublisher.Publish(event.TripActivated{TripID: trip.ID})
	return nil
}

func (tl *TripLifecycle) EndTrip(trip *domain.Trip) error {
	if trip.State != domain.TripStateActive {
		return domain.ErrInvalidTripState
	}

	trip.State = domain.TripStateEnded
	trip.EndedAt = time.Now()

	if err := tl.tripRepo.Save(trip); err != nil {
		return err
	}

	tl.eventPublisher.Publish(event.TripEnded{TripID: trip.ID})
	return nil
}

func (tl *TripLifecycle) AutoEndTrip(trip *domain.Trip) error {
	if trip.State == domain.TripStateActive {
		trip.State = domain.TripStateEnded
		trip.EndedAt = time.Now()

		if err := tl.tripRepo.Save(trip); err != nil {
			return err
		}

		tl.eventPublisher.Publish(event.TripAutoEnded{TripID: trip.ID})
	}
	return nil
}

func (tl *TripLifecycle) ValidateTrip(trip *domain.Trip) error {
	if trip.StartedAt.IsZero() || trip.ActivatedAt.IsZero() {
		return domain.ErrTripNotStarted
	}
	return nil
}
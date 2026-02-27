package trip

import (
	"context"
	"time"

	"github.com/akpatri/srt/internal/repository"
	"github.com/akpatri/srt/internal/event"
	"github.com/akpatri/srt/internal/cache"
	"github.com/akpatri/srt/internal/observability"
)

type AutoEndService struct {
	tripRepo      repository.TripRepository
	cache         cache.Cache
	eventPublisher event.Publisher
	logger        observability.Logger
}

func NewAutoEndService(tripRepo repository.TripRepository, cache cache.Cache, eventPublisher event.Publisher, logger observability.Logger) *AutoEndService {
	return &AutoEndService{
		tripRepo:      tripRepo,
		cache:         cache,
		eventPublisher: eventPublisher,
		logger:        logger,
	}
}

func (s *AutoEndService) AutoEndTrips(ctx context.Context) {
	activeTrips, err := s.tripRepo.GetActiveTrips(ctx)
	if err != nil {
		s.logger.Error("Failed to fetch active trips", err)
		return
	}

	for _, trip := range activeTrips {
		if s.shouldAutoEndTrip(trip) {
			err := s.tripRepo.EndTrip(ctx, trip.ID)
			if err != nil {
				s.logger.Error("Failed to end trip", err)
				continue
			}
			s.eventPublisher.PublishTripAutoEnded(trip.ID)
		}
	}
}

func (s *AutoEndService) shouldAutoEndTrip(trip Trip) bool {
	lastUpdateTime := trip.LastLocationUpdateTime
	idleDuration := time.Since(lastUpdateTime)

	return idleDuration > time.Duration(trip.IdleTimeout)*time.Minute
}
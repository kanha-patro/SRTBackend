package trip

import (
	"context"
	"time"

	"github.com/akpatri/srt/internal/event"
	"github.com/akpatri/srt/internal/observability"
	"github.com/akpatri/srt/internal/repository"
	"go.uber.org/zap"
)

type AutoEndService struct {
	tripRepo       repository.TripRepository
	eventPublisher event.Publisher
	logger         observability.Logger
}

func NewAutoEndService(tripRepo repository.TripRepository, eventPublisher event.Publisher, logger observability.Logger) *AutoEndService {
	return &AutoEndService{
		tripRepo:       tripRepo,
		eventPublisher: eventPublisher,
		logger:         logger,
	}
}

func (s *AutoEndService) AutoEndTrips(ctx context.Context, threshold time.Duration) {
	staleTrips, err := s.tripRepo.GetStaleTrips(ctx, threshold)
	if err != nil {
		s.logger.Error("Failed to fetch stale trips", zap.Error(err))
		return
	}

	for _, trip := range staleTrips {
		trip.EndTime = time.Now()
		trip.State = "ENDED"
		if err := s.tripRepo.UpdateTrip(ctx, trip); err != nil {
			s.logger.Error("Failed to end trip", zap.Error(err))
			continue
		}
		_ = s.eventPublisher.Publish("trip.autoended", event.TripAutoEnded{Event: event.Event{Timestamp: time.Now(), Type: "TripAutoEnded"}, TripID: trip.ID})
	}
}

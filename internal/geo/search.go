package geo

import (
	"context"
	"errors"
	"github.com/akpatri/srt/internal/domain"
	"github.com/akpatri/srt/internal/repository"
	"github.com/akpatri/srt/pkg/utils"
)

type StopSearchService struct {
	stopRepo repository.StopRepository
}

func NewStopSearchService(stopRepo repository.StopRepository) *StopSearchService {
	return &StopSearchService{stopRepo: stopRepo}
}

// SearchNearbyStops searches for stops within a specified radius from a given location.
func (s *StopSearchService) SearchNearbyStops(ctx context.Context, latitude, longitude float64, radius float64) ([]domain.Stop, error) {
	if radius <= 0 {
		return nil, errors.New("radius must be greater than zero")
	}

	stops, err := s.stopRepo.FindStopsWithinRadius(ctx, latitude, longitude, radius)
	if err != nil {
		return nil, err
	}

	return stops, nil
}

// NearestStop finds the nearest stop to a given location.
func (s *StopSearchService) NearestStop(ctx context.Context, latitude, longitude float64) (domain.Stop, error) {
	stops, err := s.stopRepo.FindAll(ctx)
	if err != nil {
		return domain.Stop{}, err
	}

	var nearestStop domain.Stop
	minDistance := utils.MaxFloat64

	for _, stop := range stops {
		distance := utils.CalculateDistance(latitude, longitude, stop.Latitude, stop.Longitude)
		if distance < minDistance {
			minDistance = distance
			nearestStop = stop
		}
	}

	if minDistance == utils.MaxFloat64 {
		return domain.Stop{}, errors.New("no stops found")
	}

	return nearestStop, nil
}

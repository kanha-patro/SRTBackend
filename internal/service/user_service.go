package service

import (
	"context"

	"github.com/akpatri/srt/internal/domain"
	"github.com/akpatri/srt/internal/repository"
)

type userService struct {
	tripRepo repository.TripRepository
}

// NewUserService creates a new UserService.
func NewUserService(tripRepo repository.TripRepository) UserService {
	return &userService{tripRepo: tripRepo}
}

// GetActiveShuttles returns active trips for an org/route filter. This is a lightweight implementation that delegates to TripRepository.
func (s *userService) GetActiveShuttles(orgCode, routeCode, nearbyLocation string) ([]domain.Trip, error) {
	// For now, treat orgCode as orgID and return active trips for that org.
	if orgCode == "" {
		return []domain.Trip{}, nil
	}
	trips, err := s.tripRepo.GetActiveTripsByOrgID(context.Background(), orgCode)
	if err != nil {
		return nil, err
	}
	// Convert []*domain.Trip to []domain.Trip
	result := make([]domain.Trip, 0, len(trips))
	for _, t := range trips {
		result = append(result, *t)
	}
	return result, nil
}

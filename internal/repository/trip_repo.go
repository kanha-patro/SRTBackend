package repository

import (
	"context"
	"time"

	"github.com/akpatri/srt/internal/domain"
	"gorm.io/gorm"
)

type tripRepository struct {
	db *gorm.DB
}

func NewTripRepository(db *gorm.DB) TripRepository {
	return &tripRepository{db: db}
}

func (r *tripRepository) CreateTrip(ctx context.Context, trip *domain.Trip) error {
	return r.db.WithContext(ctx).Create(trip).Error
}

func (r *tripRepository) GetTripByID(ctx context.Context, id string) (*domain.Trip, error) {
	var trip domain.Trip
	err := r.db.WithContext(ctx).First(&trip, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &trip, nil
}

func (r *tripRepository) UpdateTrip(ctx context.Context, trip *domain.Trip) error {
	return r.db.WithContext(ctx).Save(trip).Error
}

func (r *tripRepository) DeleteTrip(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&domain.Trip{}, id).Error
}

func (r *tripRepository) GetActiveTripsByOrgID(ctx context.Context, orgID string) ([]*domain.Trip, error) {
	var trips []*domain.Trip
	err := r.db.WithContext(ctx).Where("org_id = ? AND state = ?", orgID, "ACTIVE").Find(&trips).Error
	if err != nil {
		return nil, err
	}
	return trips, nil
}

func (r *tripRepository) GetStaleTrips(ctx context.Context, threshold time.Duration) ([]*domain.Trip, error) {
	var trips []*domain.Trip
	cutoff := time.Now().Add(-threshold)
	err := r.db.WithContext(ctx).Where("last_updated < ? AND state != ?", cutoff, "ENDED").Find(&trips).Error
	if err != nil {
		return nil, err
	}
	return trips, nil
}

package repository

import (
	"context"
	"time"

	"github.com/akpatri/srt/internal/domain"
	"gorm.io/gorm"
)

type locationRepository struct {
	db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) LocationRepository {
	return &locationRepository{db: db}
}

func (r *locationRepository) SaveLocation(ctx context.Context, location *domain.Location) error {
	return r.db.WithContext(ctx).Create(location).Error
}

func (r *locationRepository) GetLatestLocationByTripID(ctx context.Context, tripID string) (*domain.Location, error) {
	var location domain.Location
	err := r.db.WithContext(ctx).
		Where("trip_id = ?", tripID).
		Order("timestamp DESC").
		First(&location).Error
	if err != nil {
		return nil, err
	}
	return &location, nil
}

func (r *locationRepository) GetLocationsByTripID(ctx context.Context, tripID string, startTime time.Time, endTime time.Time) ([]domain.Location, error) {
	var locations []domain.Location
	err := r.db.WithContext(ctx).
		Where("trip_id = ? AND timestamp BETWEEN ? AND ?", tripID, startTime, endTime).
		Find(&locations).Error
	if err != nil {
		return nil, err
	}
	return locations, nil
}

func (r *locationRepository) FindActiveLocations(ctx context.Context, orgID string) ([]domain.Location, error) {
	// join trips and locations to get latest location per active trip for the org
	var locations []domain.Location
	sub := r.db.WithContext(ctx).Table("trips").Select("id").Where("org_id = ? AND state = ?", orgID, "ACTIVE")
	err := r.db.WithContext(ctx).
		Where("trip_id IN (?)", sub).
		Order("timestamp DESC").
		Find(&locations).Error
	if err != nil {
		return nil, err
	}
	return locations, nil
}

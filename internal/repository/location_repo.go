package repository

import (
	"context"
	"time"

	"github.com/akpatri/srt/internal/domain"
	"gorm.io/gorm"
)

type LocationRepository interface {
	SaveLocation(ctx context.Context, location *domain.Location) error
	GetLatestLocationByTripID(ctx context.Context, tripID string) (*domain.Location, error)
	GetLocationsByTripID(ctx context.Context, tripID string, startTime time.Time, endTime time.Time) ([]domain.Location, error)
}

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
package repository

import (
	"context"

	"github.com/akpatri/srt/internal/domain"
	"github.com/akpatri/srt/pkg/utils"
	"gorm.io/gorm"
)

type stopRepository struct {
	db *gorm.DB
}

func NewStopRepository(db *gorm.DB) StopRepository {
	return &stopRepository{db: db}
}

func (r *stopRepository) Create(stop *domain.Stop) error {
	return r.db.Create(stop).Error
}

func (r *stopRepository) Update(stop *domain.Stop) error {
	return r.db.Save(stop).Error
}

func (r *stopRepository) Delete(stopID string) error {
	return r.db.Delete(&domain.Stop{}, stopID).Error
}

func (r *stopRepository) FindByID(stopID string) (*domain.Stop, error) {
	var stop domain.Stop
	err := r.db.First(&stop, stopID).Error
	if err != nil {
		return nil, err
	}
	return &stop, nil
}

func (r *stopRepository) FindAllByRouteID(routeID string) ([]domain.Stop, error) {
	var stops []domain.Stop
	err := r.db.Where("route_id = ?", routeID).Find(&stops).Error
	if err != nil {
		return nil, err
	}
	return stops, nil
}

func (r *stopRepository) FindAll(ctx context.Context) ([]domain.Stop, error) {
	var stops []domain.Stop
	if err := r.db.WithContext(ctx).Find(&stops).Error; err != nil {
		return nil, err
	}
	return stops, nil
}

func (r *stopRepository) FindStopsWithinRadius(ctx context.Context, latitude, longitude, radius float64) ([]domain.Stop, error) {
	// Naive implementation: load all stops and filter by distance.
	var results []domain.Stop
	var stops []domain.Stop
	if err := r.db.WithContext(ctx).Find(&stops).Error; err != nil {
		return nil, err
	}

	for _, s := range stops {
		if utils.CalculateDistance(latitude, longitude, s.Latitude, s.Longitude) <= radius {
			results = append(results, s)
		}
	}
	return results, nil
}

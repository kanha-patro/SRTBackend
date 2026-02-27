package repository

import (
	"github.com/akpatri/srt/internal/domain"
	"gorm.io/gorm"
)

type StopRepository interface {
	Create(stop *domain.Stop) error
	Update(stop *domain.Stop) error
	Delete(stopID string) error
	FindByID(stopID string) (*domain.Stop, error)
	FindAllByRouteID(routeID string) ([]domain.Stop, error)
}

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
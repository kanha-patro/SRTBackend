package repository

import (
	"github.com/akpatri/srt/internal/domain"
	"gorm.io/gorm"
)

type RouteRepository interface {
	CreateRoute(route *domain.Route) error
	GetRouteByID(id string) (*domain.Route, error)
	UpdateRoute(route *domain.Route) error
	DeleteRoute(id string) error
	GetAllRoutes(orgID string) ([]domain.Route, error)
}

type routeRepository struct {
	db *gorm.DB
}

func NewRouteRepository(db *gorm.DB) RouteRepository {
	return &routeRepository{db: db}
}

func (r *routeRepository) CreateRoute(route *domain.Route) error {
	return r.db.Create(route).Error
}

func (r *routeRepository) GetRouteByID(id string) (*domain.Route, error) {
	var route domain.Route
	if err := r.db.First(&route, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &route, nil
}

func (r *routeRepository) UpdateRoute(route *domain.Route) error {
	return r.db.Save(route).Error
}

func (r *routeRepository) DeleteRoute(id string) error {
	return r.db.Delete(&domain.Route{}, id).Error
}

func (r *routeRepository) GetAllRoutes(orgID string) ([]domain.Route, error) {
	var routes []domain.Route
	if err := r.db.Where("org_id = ?", orgID).Find(&routes).Error; err != nil {
		return nil, err
	}
	return routes, nil
}
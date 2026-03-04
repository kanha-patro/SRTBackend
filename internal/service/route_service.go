package service

import (
	"errors"

	"github.com/akpatri/srt/internal/domain"
	"github.com/akpatri/srt/internal/repository"
)

type routeService struct {
	routeRepo repository.RouteRepository
}

func NewRouteService(routeRepo repository.RouteRepository) RouteService {
	return &routeService{
		routeRepo: routeRepo,
	}
}

func (s *routeService) CreateRoute(route *domain.Route) error {
	if route == nil {
		return errors.New("route cannot be nil")
	}
	return s.routeRepo.CreateRoute(route)
}

func (s *routeService) UpdateRoute(route *domain.Route) error {
	if route == nil {
		return errors.New("route cannot be nil")
	}
	return s.routeRepo.UpdateRoute(route)
}

func (s *routeService) DeleteRoute(routeID string) error {
	if routeID == "" {
		return errors.New("routeID cannot be empty")
	}
	return s.routeRepo.DeleteRoute(routeID)
}

func (s *routeService) GetRoute(routeID string) (*domain.Route, error) {
	if routeID == "" {
		return nil, errors.New("routeID cannot be empty")
	}
	return s.routeRepo.GetRouteByID(routeID)
}

func (s *routeService) GetAllRoutes(orgID string) ([]domain.Route, error) {
	if orgID == "" {
		return nil, errors.New("orgID cannot be empty")
	}
	return s.routeRepo.GetAllRoutes(orgID)
}

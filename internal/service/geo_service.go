package service

import (
	"context"
	"strconv"

	"github.com/akpatri/srt/internal/domain"
	"github.com/akpatri/srt/internal/repository"
	"github.com/akpatri/srt/pkg/errors"
	"github.com/akpatri/srt/pkg/utils"
)

type geoServiceImpl struct {
	stopRepo repository.StopRepository
}

func NewGeoService(stopRepo repository.StopRepository) GeoService {
	return &geoServiceImpl{stopRepo: stopRepo}
}

func (s *geoServiceImpl) SnapLocation(lat string, lon string) (interface{}, error) {
	lf, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		return nil, errors.NewBadRequestError("invalid latitude")
	}
	lf2, err := strconv.ParseFloat(lon, 64)
	if err != nil {
		return nil, errors.NewBadRequestError("invalid longitude")
	}
	stops, err := s.stopRepo.FindAll(context.Background())
	if err != nil {
		return nil, err
	}
	// find nearest
	var nearest domain.Stop
	min := utils.MaxFloat64
	for _, st := range stops {
		d := utils.CalculateDistance(lf, lf2, st.Latitude, st.Longitude)
		if d < min {
			min = d
			nearest = st
		}
	}
	return nearest, nil
}

func (s *geoServiceImpl) SearchNearbyStops(lat string, lon string) ([]domain.Location, error) {
	lf, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		return nil, errors.NewBadRequestError("invalid latitude")
	}
	lf2, err := strconv.ParseFloat(lon, 64)
	if err != nil {
		return nil, errors.NewBadRequestError("invalid longitude")
	}
	stops, err := s.stopRepo.FindStopsWithinRadius(context.Background(), lf, lf2, 500)
	if err != nil {
		return nil, err
	}
	var locs []domain.Location
	for _, st := range stops {
		locs = append(locs, domain.Location{Latitude: st.Latitude, Longitude: st.Longitude})
	}
	return locs, nil
}

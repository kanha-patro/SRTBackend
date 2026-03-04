package geo

import (
	"context"

	"github.com/akpatri/srt/internal/domain"
	"github.com/akpatri/srt/pkg/utils"
)

// SnapToNearestStop snaps a given location to the nearest stop using Haversine distance.
func SnapToNearestStop(ctx context.Context, location domain.Location, stops []domain.Stop) (domain.Stop, error) {
	var nearestStop domain.Stop
	minDistance := utils.MaxFloat64

	for _, stop := range stops {
		distance := utils.CalculateDistance(location.Latitude, location.Longitude, stop.Latitude, stop.Longitude)
		if distance < minDistance {
			minDistance = distance
			nearestStop = stop
		}
	}

	return nearestStop, nil
}

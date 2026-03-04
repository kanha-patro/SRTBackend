package geo

import (
	"github.com/akpatri/srt/internal/domain"
)

// SnapToNearestStop snaps a given location to the nearest stop.
// It returns the nearest stop and any error encountered during the process.
func SnapToNearestStop(location domain.Location, stops []domain.Stop) (domain.Stop, error) {
	var nearestStop domain.Stop
	minDistance := float64(-1)

	for _, stop := range stops {
		distance := calculateDistance(location, stop.Location)
		if minDistance == -1 || distance < minDistance {
			minDistance = distance
			nearestStop = stop
		}
	}

	return nearestStop, nil
}

// calculateDistance calculates the distance between two locations.
// This is a placeholder function and should be replaced with an actual distance calculation logic.
func calculateDistance(loc1, loc2 domain.Location) float64 {
	// Implement the distance calculation logic here (e.g., Haversine formula).
	return 0.0
}
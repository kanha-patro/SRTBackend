package domain

import (
	"time"
)

// Location represents the geographical location of a shuttle.
type Location struct {
	ID        string    `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	TripID    string    `json:"trip_id" gorm:"index"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Timestamp time.Time `json:"timestamp" gorm:"autoUpdateTime"`
}

// NewLocation creates a new Location instance.
func NewLocation(tripID string, latitude float64, longitude float64) *Location {
	return &Location{
		TripID:    tripID,
		Latitude:  latitude,
		Longitude: longitude,
		Timestamp: time.Now(),
	}
}
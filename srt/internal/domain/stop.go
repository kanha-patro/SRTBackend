package domain

import (
	"time"
)

// Stop represents a stop in the shuttle tracking system.
type Stop struct {
	ID        string    `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	RouteID   string    `json:"route_id" gorm:"type:uuid;not null"`
	Name      string    `json:"name" gorm:"not null"`
	Latitude  float64   `json:"latitude" gorm:"not null"`
	Longitude float64   `json:"longitude" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// NewStop creates a new Stop instance.
func NewStop(routeID, name string, latitude, longitude float64) *Stop {
	return &Stop{
		RouteID:   routeID,
		Name:      name,
		Latitude:  latitude,
		Longitude: longitude,
	}
}

// Update updates the stop's details.
func (s *Stop) Update(name string, latitude, longitude float64) {
	s.Name = name
	s.Latitude = latitude
	s.Longitude = longitude
}
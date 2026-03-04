package domain

import (
	"time"

	"github.com/google/uuid"
)

// Stop represents a stop in the shuttle tracking system.
// Stop represents a stop in a shuttle route.
type Stop struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	RouteID   uuid.UUID `json:"route_id" gorm:"type:uuid;not null"`
	Name      string    `json:"name" gorm:"size:100;not null"`
	Latitude  float64   `json:"latitude" gorm:"not null"`
	Longitude float64   `json:"longitude" gorm:"not null"`
	Order     int       `json:"order" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// NewStop creates a new Stop instance.
func NewStop(routeID uuid.UUID, name string, latitude, longitude float64) *Stop {
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
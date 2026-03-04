package domain

import "time"

// Location represents the geographical location of a shuttle.
type Location struct {
	ID        string    `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	TripID    string    `json:"trip_id" gorm:"type:uuid;index"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Timestamp time.Time `json:"timestamp" gorm:"autoCreateTime"`
	Accuracy  float64   `json:"accuracy,omitempty"`
	Speed     float64   `json:"speed,omitempty"`
}

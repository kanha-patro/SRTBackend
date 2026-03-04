package domain

import (
	"time"
)

// Trip represents a shuttle trip with its associated details.
type Trip struct {
	ID            string    `json:"id"`
	RouteID       string    `json:"route_id"`
	DriverID      string    `json:"driver_id"`
	State         TripState `json:"state"`
	StartTime     time.Time `json:"start_time"`
	EndTime       time.Time `json:"end_time,omitempty"`
	LastUpdated   time.Time `json:"last_updated"`
	Location      Location  `json:"location"`
	OTP           string    `json:"otp"`
}

// TripState represents the state of a trip in its lifecycle.
type TripState string

const (
	// Trip states
	CREATED TripState = "CREATED"
	STARTED TripState = "STARTED"
	ACTIVE  TripState = "ACTIVE"
	ENDED   TripState = "ENDED"
)

// NewTrip creates a new trip instance.
func NewTrip(id, routeID, driverID, otp string) *Trip {
	return &Trip{
		ID:          id,
		RouteID:     routeID,
		DriverID:    driverID,
		State:       CREATED,
		StartTime:   time.Time{},
		LastUpdated: time.Now(),
		OTP:         otp,
	}
}

// StartTrip sets the trip state to STARTED and records the start time.
func (t *Trip) StartTrip() {
	t.State = STARTED
	t.StartTime = time.Now()
	t.LastUpdated = time.Now()
}

// ActivateTrip sets the trip state to ACTIVE.
func (t *Trip) ActivateTrip() {
	t.State = ACTIVE
	t.LastUpdated = time.Now()
}

// EndTrip sets the trip state to ENDED and records the end time.
func (t *Trip) EndTrip() {
	t.State = ENDED
	t.EndTime = time.Now()
	t.LastUpdated = time.Now()
}

// UpdateLocation updates the trip's last known location and timestamp.
func (t *Trip) UpdateLocation(location Location) {
	t.Location = location
	t.LastUpdated = time.Now()
}

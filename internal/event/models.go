package event

import "time"

// Event represents a generic event in the system.
type Event struct {
	ID        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Type      string    `json:"type"`
	Payload   any       `json:"payload"`
}

// TripStarted represents an event when a trip starts.
type TripStarted struct {
	Event
	TripID string `json:"trip_id"`
}

// TripActivated represents an event when a trip becomes active.
type TripActivated struct {
	Event
	TripID string `json:"trip_id"`
}

// LocationUpdated represents an event when a driver's location is updated.
type LocationUpdated struct {
	Event
	TripID    string  `json:"trip_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// TripAutoEnded represents an event when a trip is automatically ended.
type TripAutoEnded struct {
	Event
	TripID string `json:"trip_id"`
}

// OrgSuspended represents an event when an organization is suspended.
type OrgSuspended struct {
	Event
	OrgID string `json:"org_id"`
}

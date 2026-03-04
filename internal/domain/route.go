package domain

import (
	"time"

	"github.com/google/uuid"
)

// Route represents a shuttle route in the system.
type Route struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	OrgID     uuid.UUID `json:"org_id" gorm:"type:uuid;not null"`
	Name      string    `json:"name" gorm:"size:100;not null"`
	Stops     []Stop    `json:"stops" gorm:"foreignKey:RouteID"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// NewRoute creates a new Route instance.
func NewRoute(orgID uuid.UUID, name string) *Route {
	return &Route{
		ID:    uuid.New(),
		OrgID: orgID,
		Name:  name,
		Stops: []Stop{},
	}
}

// AddStop adds a new stop to the route.
func (r *Route) AddStop(name string, latitude, longitude float64, order int) {
	stop := Stop{
		ID:        uuid.New(),
		RouteID:   r.ID,
		Name:      name,
		Latitude:  latitude,
		Longitude: longitude,
		Order:     order,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	r.Stops = append(r.Stops, stop)
}

// UpdateStop updates an existing stop in the route.
func (r *Route) UpdateStop(stopID uuid.UUID, name string, latitude, longitude float64) {
	for i, stop := range r.Stops {
		if stop.ID == stopID {
			r.Stops[i].Name = name
			r.Stops[i].Latitude = latitude
			r.Stops[i].Longitude = longitude
			r.Stops[i].UpdatedAt = time.Now()
			break
		}
	}
}

// RemoveStop removes a stop from the route.
func (r *Route) RemoveStop(stopID uuid.UUID) {
	for i, stop := range r.Stops {
		if stop.ID == stopID {
			r.Stops = append(r.Stops[:i], r.Stops[i+1:]...)
			break
		}
	}
}

package domain

import (
	"time"

	"github.com/google/uuid"
)

// Org represents an organization in the system.
type Org struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsActive  bool      `json:"is_active"`
}

// NewOrg creates a new organization instance.
func NewOrg(name, code string) *Org {
	return &Org{
		ID:        uuid.New(),
		Name:      name,
		Code:      code,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsActive:  true,
	}
}

// Update updates the organization details.
func (o *Org) Update(name, code string) {
	o.Name = name
	o.Code = code
	o.UpdatedAt = time.Now()
}

// Deactivate marks the organization as inactive.
func (o *Org) Deactivate() {
	o.IsActive = false
	o.UpdatedAt = time.Now()
}

// Activate marks the organization as active.
func (o *Org) Activate() {
	o.IsActive = true
	o.UpdatedAt = time.Now()
}
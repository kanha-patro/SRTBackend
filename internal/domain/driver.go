package domain

import (
	"fmt"
	"time"
)

// Driver represents a driver in the shuttle tracking system.
type Driver struct {
	ID        string    `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	OrgID     string    `json:"org_id" gorm:"type:uuid;not null"`
	RouteID   string    `json:"route_id" gorm:"type:uuid;not null"`
	Code      string    `json:"code" gorm:"unique;not null"`
	OTP       string    `json:"otp" gorm:"not null"`
	DeviceID  string    `json:"device_id" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// NewDriver creates a new Driver instance.
func NewDriver(orgID, routeID, code, otp, deviceID string) *Driver {
	return &Driver{
		OrgID:    orgID,
		RouteID:  routeID,
		Code:     code,
		OTP:      otp,
		DeviceID: deviceID,
	}
}

// Validate checks if the driver fields are valid.
func (d *Driver) Validate() error {
	if d.OrgID == "" || d.RouteID == "" || d.Code == "" || d.OTP == "" || d.DeviceID == "" {
		return fmt.Errorf("all fields must be filled")
	}
	return nil
}
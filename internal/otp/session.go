package otp

import (
	"errors"
	"sync"
	"time"
)

// Session represents an OTP session for a driver.
type Session struct {
	OTP        string
	OrgCode    string
	RouteCode  string
	DriverCode string
	DeviceID   string
	ExpiresAt  time.Time
	Used       bool
	mu         sync.Mutex
}

// NewSession creates a new OTP session with the specified parameters.
func NewSession(orgCode, routeCode, driverCode, deviceID string, expiryDuration time.Duration) *Session {
	return &Session{
		OTP:        generateOTP(), // Assume generateOTP is a function that generates a secure OTP
		OrgCode:    orgCode,
		RouteCode:  routeCode,
		DriverCode: driverCode,
		DeviceID:   deviceID,
		ExpiresAt:  time.Now().Add(expiryDuration),
		Used:       false,
	}
}

// Validate checks if the OTP session is valid and not expired.
func (s *Session) Validate(otp string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.Used {
		return errors.New("OTP has already been used")
	}
	if time.Now().After(s.ExpiresAt) {
		return errors.New("OTP has expired")
	}
	if s.OTP != otp {
		return errors.New("invalid OTP")
	}
	s.Used = true // Mark the OTP as used
	return nil
}

// IsExpired checks if the OTP session has expired.
func (s *Session) IsExpired() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return time.Now().After(s.ExpiresAt)
}

// Reset resets the OTP session, allowing it to be reused.
func (s *Session) Reset() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Used = false
	s.OTP = generateOTP() // Regenerate OTP
	s.ExpiresAt = time.Now().Add(time.Duration(5) * time.Minute) // Reset expiry
}
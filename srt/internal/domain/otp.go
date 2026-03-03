package domain

import (
	"errors"
	"time"
)

type OTP struct {
	Code      string
	OrgCode   string
	RouteCode string
	DriverCode string
	ExpiresAt time.Time
	Used      bool
	DeviceID  string
}

var (
	ErrOTPExpired   = errors.New("otp has expired")
	ErrOTPUsed      = errors.New("otp has already been used")
	ErrOTPInvalid   = errors.New("invalid otp")
)

// NewOTP generates a new OTP with a specified expiry duration.
func NewOTP(orgCode, routeCode, driverCode, deviceID string, expiryDuration time.Duration) OTP {
	return OTP{
		Code:      generateRandomOTP(), // Implement this function to generate a random OTP
		OrgCode:   orgCode,
		RouteCode: routeCode,
		DriverCode: driverCode,
		ExpiresAt: time.Now().Add(expiryDuration),
		Used:      false,
		DeviceID:  deviceID,
	}
}

// Validate checks if the OTP is valid based on its state and expiry.
func (otp *OTP) Validate(inputCode string) error {
	if otp.Used {
		return ErrOTPUsed
	}
	if time.Now().After(otp.ExpiresAt) {
		return ErrOTPExpired
	}
	if otp.Code != inputCode {
		return ErrOTPInvalid
	}
	return nil
}

// MarkAsUsed marks the OTP as used.
func (otp *OTP) MarkAsUsed() {
	otp.Used = true
}
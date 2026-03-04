package domain

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"time"
)

type OTP struct {
	Code       string
	OrgCode    string
	RouteCode  string
	DriverCode string
	ExpiresAt  time.Time
	Used       bool
	DeviceID   string
}

var (
	ErrOTPExpired = errors.New("otp has expired")
	ErrOTPUsed    = errors.New("otp has already been used")
	ErrOTPInvalid = errors.New("invalid otp")
)

// NewOTP generates a new OTP with a specified expiry duration.
func NewOTP(orgCode, routeCode, driverCode, deviceID string, expiryDuration time.Duration) OTP {
	return OTP{
		Code:       generateRandomOTP(),
		OrgCode:    orgCode,
		RouteCode:  routeCode,
		DriverCode: driverCode,
		ExpiresAt:  time.Now().Add(expiryDuration),
		Used:       false,
		DeviceID:   deviceID,
	}
}

// generateRandomOTP returns a cryptographically secure 6-digit numeric OTP.
func generateRandomOTP() string {
	n, _ := rand.Int(rand.Reader, big.NewInt(1000000))
	return fmt.Sprintf("%06d", n.Int64())
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

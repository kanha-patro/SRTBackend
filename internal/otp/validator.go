package otp

import (
	"errors"
	"time"
)

// OTPValidator validates OTPs based on defined rules.
type OTPValidator struct {
	otpExpiryDuration time.Duration
}

// NewOTPValidator creates a new instance of OTPValidator.
func NewOTPValidator(expiryDuration time.Duration) *OTPValidator {
	return &OTPValidator{
		otpExpiryDuration: expiryDuration,
	}
}

// Validate checks if the provided OTP is valid based on its expiry and session binding.
func (v *OTPValidator) Validate(otp string, createdAt time.Time, sessionID string, currentSessionID string) error {
	if otp == "" {
		return errors.New("OTP cannot be empty")
	}

	if currentSessionID != sessionID {
		return errors.New("OTP is not bound to this session")
	}

	if time.Since(createdAt) > v.otpExpiryDuration {
		return errors.New("OTP has expired")
	}

	return nil
}
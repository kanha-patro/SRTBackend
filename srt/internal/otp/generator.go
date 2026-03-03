package otp

import (
	"crypto/rand"
	"encoding/base32"
	"fmt"
	"time"
)

// Generator is responsible for generating OTPs.
type Generator struct {
	otpLength int
	expiry    time.Duration
}

// NewGenerator creates a new OTP generator with the specified length and expiry duration.
func NewGenerator(length int, expiry time.Duration) *Generator {
	return &Generator{
		otpLength: length,
		expiry:    expiry,
	}
}

// Generate generates a new OTP and returns it as a string.
func (g *Generator) Generate() (string, error) {
	otp := make([]byte, g.otpLength)
	_, err := rand.Read(otp)
	if err != nil {
		return "", err
	}
	return base32.StdEncoding.EncodeToString(otp), nil
}

// IsExpired checks if the given timestamp is expired based on the generator's expiry duration.
func (g *Generator) IsExpired(timestamp time.Time) bool {
	return time.Since(timestamp) > g.expiry
}
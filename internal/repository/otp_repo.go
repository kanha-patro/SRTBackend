package repository

import (
	"sync"
	"time"
)

// inMemoryOTPRepo is a simple thread-safe in-memory implementation of OTPRepository.
type inMemoryOTPRepo struct {
	mu sync.RWMutex
	// key: org|route|driver
	store map[string]OTPEntry
}

// NewInMemoryOTPRepository creates a new in-memory OTP repository.
func NewInMemoryOTPRepository() OTPRepository {
	return &inMemoryOTPRepo{store: make(map[string]OTPEntry)}
}

func makeKey(org, route, driver string) string {
	return org + "|" + route + "|" + driver
}

func (r *inMemoryOTPRepo) StoreOTP(orgCode, routeCode, driverCode, code string, expiry time.Time) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.store[makeKey(orgCode, routeCode, driverCode)] = OTPEntry{Code: code, Expiry: expiry}
	return nil
}

func (r *inMemoryOTPRepo) GetOTP(orgCode, routeCode, driverCode string) (*OTPEntry, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	key := makeKey(orgCode, routeCode, driverCode)
	if v, ok := r.store[key]; ok {
		// Return copy
		entry := v
		// Optionally delete expired entries lazily
		if time.Now().After(entry.Expiry) {
			return nil, nil
		}
		return &entry, nil
	}
	return nil, nil
}

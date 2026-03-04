package utils

import "time"

// GetCurrentUTC returns the current UTC time.
func GetCurrentUTC() time.Time {
    return time.Now().UTC()
}

// ParseTime parses a string representation of time in RFC3339 format.
func ParseTime(value string) (time.Time, error) {
    return time.Parse(time.RFC3339, value)
}

// FormatTime formats a time.Time object into a string in RFC3339 format.
func FormatTime(t time.Time) string {
    return t.Format(time.RFC3339)
}

// IsExpired checks if the given time has expired based on the current UTC time.
func IsExpired(t time.Time) bool {
    return t.Before(GetCurrentUTC())
}
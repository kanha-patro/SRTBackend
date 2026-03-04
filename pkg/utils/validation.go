package utils

import (
	"regexp"
	"time"
)

// ValidateEmail checks if the provided email is valid.
func ValidateEmail(email string) bool {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

// ValidatePhone checks if the provided phone number is valid.
func ValidatePhone(phone string) bool {
	const phoneRegex = `^\+?[1-9]\d{1,14}$`
	re := regexp.MustCompile(phoneRegex)
	return re.MatchString(phone)
}

// ValidateUUID checks if the provided UUID is valid.
func ValidateUUID(uuid string) bool {
	const uuidRegex = `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`
	re := regexp.MustCompile(uuidRegex)
	return re.MatchString(uuid)
}

// ValidateTimestamp checks if the provided timestamp is valid.
func ValidateTimestamp(ts time.Time) bool {
	return !ts.IsZero()
}

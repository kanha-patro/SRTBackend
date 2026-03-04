package errors

import "fmt"

// CustomError defines a custom error type for the application.
type CustomError struct {
	Code    int
	Message string
}

// New creates a new CustomError with the given code and message.
func New(code int, message string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}

// Error implements the error interface for CustomError.
func (e *CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

// Is checks if the error matches a specific code.
func Is(err error, code int) bool {
	if customErr, ok := err.(*CustomError); ok {
		return customErr.Code == code
	}
	return false
}

// Convenience constructors for common HTTP-style errors.
func NewNotFoundError(message string) error {
	return New(404, message)
}

func NewUnauthorizedError(message string) error {
	return New(401, message)
}

func NewInvalidStateError(message string) error {
	return New(409, message)
}

func NewInternalServerError(message string) error {
	return New(500, message)
}

func NewBadRequestError(message string) error {
	return New(400, message)
}

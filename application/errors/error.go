package apperrors

import "fmt"

type ErrorType string

const (
	NotFound     ErrorType = "NOT_FOUND"
	BadRequest   ErrorType = "BAD_REQUEST"
	Internal     ErrorType = "INTERNAL"
	Unauthorized ErrorType = "UNAUTHORIZED"
)

type AppError struct {
	Type    ErrorType // Machine-readable code
	Message string    // Human-readable message for the client
	Err     error     // The actual internal error (for logging)
}

// Implementing the 'error' interface
func (e *AppError) Error() string {
	return fmt.Sprintf("[%s] %s: %v", e.Type, e.Message, e.Err)
}

// Unwrap allows errors.Is() and errors.As() to work
func (e *AppError) Unwrap() error {
	return e.Err
}

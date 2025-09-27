package web

import "errors"

// shutdownError is a type used to help with graceful termination of the service
type shutdownError struct {
	Message string
}

// NewShutdownError return an error that causes the frameowork to signal a graceful shutdown
func NewShutDownError(message string) error {
	return &shutdownError{message}
}

// Error is the implementation of the error interface
func (se *shutdownError) Error() string {
	return se.Message
}

// IsShutdown checks to see if the shutdown error is contained
func IsShutdown(err error) bool {
	var se *shutdownError
	return errors.As(err, &se)
}

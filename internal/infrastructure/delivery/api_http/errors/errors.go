package errors

import (
	"errors"
	"time"
)

var (
	ErrUnauthorized   = errors.New("unauthorized")
	ErrForbidden      = errors.New("forbidden")
	ErrInvalidPayload = errors.New("invalid payload")
)

type HttpConcreteError struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

func BuildHttpError(err error) HttpConcreteError {
	return HttpConcreteError{
		Message:   err.Error(),
		Timestamp: time.Now(),
	}
}

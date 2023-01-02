package codeflowauth

import "errors"

var (
	ErrCodeInvalid          = errors.New("auth code is invalid")
	ErrTokenInvalid         = errors.New("token is invalid")
	ErrFailedToFetchProfile = errors.New("failed to fetch user profile")
)

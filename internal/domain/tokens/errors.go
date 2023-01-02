package tokens

import "errors"

var (
	ErrAccessTokenExpired = errors.New("access token expired")
	ErrInvalidToken       = errors.New("invalid token")
)

package tokens

import "time"

type AccessTokenClaims struct {
	UserID    string    `json:"user_id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

func NewAccessTokenClaims(ID string, duration time.Duration) *AccessTokenClaims {
	return &AccessTokenClaims{
		UserID:    ID,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(duration),
	}
}

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

func NewTokenPair(accessToken, refreshToken string) *TokenPair {
	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

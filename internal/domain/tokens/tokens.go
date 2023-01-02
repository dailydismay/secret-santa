package tokens

import "context"

type SignTokensPayload struct {
	UserID string
}

type TokensProvider interface {
	VerifyAccessToken(context.Context, string) (*AccessTokenClaims, error)
	SignAccessToken(context.Context, *SignTokensPayload) (string, error)
	SignTokens(context.Context, *SignTokensPayload) (*TokenPair, error)
}

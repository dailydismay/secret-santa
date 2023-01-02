package paseto

import (
	"context"
	"fmt"
	"secretsanta/internal/config"
	"secretsanta/internal/domain/tokens"
	"time"

	paseto_lib "github.com/o1egl/paseto"

	"go.uber.org/fx"
	"golang.org/x/crypto/chacha20poly1305"
)

type PasetoOptions struct {
	fx.In

	Cfg config.Config
}

type implementation struct {
	secretKey      []byte
	pst            *paseto_lib.V2
	accessTokenTTL time.Duration
}

func (i *implementation) signAccessToken(claims *tokens.AccessTokenClaims) (string, error) {
	token, err := i.pst.Encrypt(i.secretKey, claims, nil)
	if err != nil {
		return "", fmt.Errorf("%w: failed to sign token", err)
	}

	return token, nil
}

func NewPasetoTokenService(opts PasetoOptions) (tokens.TokensProvider, error) {
	if len(opts.Cfg.AccessTokenSecret) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("token secret key size is incorrect")
	}

	return &implementation{
		secretKey:      []byte(opts.Cfg.AccessTokenSecret),
		pst:            paseto_lib.NewV2(),
		accessTokenTTL: opts.Cfg.AccessTokenTTL,
	}, nil
}

func (i *implementation) VerifyAccessToken(ctx context.Context, token string) (*tokens.AccessTokenClaims, error) {
	claims := &tokens.AccessTokenClaims{}

	err := i.pst.Decrypt(token, i.secretKey, claims, nil)
	if err != nil {
		return nil, tokens.ErrInvalidToken
	}

	if time.Now().After(claims.ExpiresAt) {
		return nil, tokens.ErrAccessTokenExpired
	}

	return claims, nil
}

func (i *implementation) SignAccessToken(ctx context.Context, payload *tokens.SignTokensPayload) (string, error) {
	claims := tokens.NewAccessTokenClaims(payload.UserID, i.accessTokenTTL)
	return i.signAccessToken(claims)
}

func (i *implementation) SignTokens(ctx context.Context, payload *tokens.SignTokensPayload) (*tokens.TokenPair, error) {
	accessToken, err := i.SignAccessToken(ctx, payload)
	if err != nil {
		return nil, err
	}

	return tokens.NewTokenPair(accessToken, ""), nil
}

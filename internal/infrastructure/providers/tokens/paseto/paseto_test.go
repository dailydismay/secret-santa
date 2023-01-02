package paseto_test

import (
	"context"
	"math/rand"
	"secretsanta/internal/config"
	"secretsanta/internal/domain/tokens"
	"secretsanta/internal/infrastructure/providers/tokens/paseto"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPasetoTokenServiceInitialization(t *testing.T) {
	assert := assert.New(t)

	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 32)
	rand.Read(b)

	svc, err := paseto.NewPasetoTokenService(paseto.PasetoOptions{
		Cfg: config.Config{
			AuthConfig: config.AuthConfig{
				AccessTokenSecret: string(b),
			},
		},
	})

	assert.NoError(err)
	assert.NotNil(svc)
}

func TestPasetoTokenServiceSign(t *testing.T) {
	assert := assert.New(t)
	userID := "random-user-id"
	ttl := time.Minute

	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 32)
	rand.Read(b)

	svc, err := paseto.NewPasetoTokenService(paseto.PasetoOptions{
		Cfg: config.Config{
			AuthConfig: config.AuthConfig{
				AccessTokenSecret: string(b),
				AccessTokenTTL:    ttl,
			},
		},
	})

	assert.NoError(err)
	assert.NotNil(svc)

	at, err := svc.SignAccessToken(context.Background(), &tokens.SignTokensPayload{
		UserID: userID,
	})
	assert.NoError(err)
	assert.NotEmpty(at)
}

func TestPasetoTokenServiceVerifyAccessToken(t *testing.T) {
	assert := assert.New(t)
	userID := "random-user-id"
	ttl := time.Minute

	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 32)
	rand.Read(b)

	svc, err := paseto.NewPasetoTokenService(paseto.PasetoOptions{
		Cfg: config.Config{
			AuthConfig: config.AuthConfig{
				AccessTokenSecret: string(b),
				AccessTokenTTL:    ttl,
			},
		},
	})
	assert.NoError(err)
	assert.NotNil(svc)

	at, err := svc.SignAccessToken(context.Background(), &tokens.SignTokensPayload{
		UserID: userID,
	})
	assert.NoError(err)
	assert.NotEmpty(at)

	claims, err := svc.VerifyAccessToken(context.Background(), at)
	assert.NoError(err)
	assert.NotEmpty(claims)
	assert.Equal(claims.UserID, userID)
}

func TestPasetoTokenServiceAccessTokenExpired(t *testing.T) {
	assert := assert.New(t)
	userID := "random-user-id"
	ttl := -time.Minute

	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 32)
	rand.Read(b)

	svc, err := paseto.NewPasetoTokenService(paseto.PasetoOptions{
		Cfg: config.Config{
			AuthConfig: config.AuthConfig{
				AccessTokenSecret: string(b),
				AccessTokenTTL:    ttl,
			},
		},
	})
	assert.NoError(err)
	assert.NotNil(svc)

	at, err := svc.SignAccessToken(context.Background(), &tokens.SignTokensPayload{
		UserID: userID,
	})
	assert.NoError(err)
	assert.NotEmpty(at)

	_, err = svc.VerifyAccessToken(context.Background(), at)
	assert.Error(err)
	assert.ErrorIs(err, tokens.ErrAccessTokenExpired)
}

func TestPasetoTokenServiceAccessTokenInvalid(t *testing.T) {
	assert := assert.New(t)
	ttl := -time.Minute

	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 32)
	rand.Read(b)

	svc, err := paseto.NewPasetoTokenService(paseto.PasetoOptions{
		Cfg: config.Config{
			AuthConfig: config.AuthConfig{
				AccessTokenSecret: string(b),
				AccessTokenTTL:    ttl,
			},
		},
	})
	assert.NoError(err)
	assert.NotNil(svc)

	_, err = svc.VerifyAccessToken(context.Background(), "unknown-token")
	assert.Error(err)
	assert.ErrorIs(err, tokens.ErrInvalidToken)
}

package auth

import (
	"secretsanta/internal/domain/tokens"
	"secretsanta/internal/infrastructure/delivery/api_http/errors"

	"github.com/gofiber/fiber/v2"
)

func GetClaims(c *fiber.Ctx) (*tokens.AccessTokenClaims, error) {
	claims := c.Locals(ContextUserClaimsKey)
	if claims == nil {
		return nil, errors.ErrUnauthorized
	}

	clmz, ok := claims.(tokens.AccessTokenClaims)
	if !ok {
		return nil, errors.ErrUnauthorized
	}

	return &clmz, nil
}

func MustGetClaims(c *fiber.Ctx) *tokens.AccessTokenClaims {
	if climes, err := GetClaims(c); err != nil {
		panic(err)
	} else {
		return climes
	}
}

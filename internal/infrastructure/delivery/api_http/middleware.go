package api_http

import (
	"net/http"
	"secretsanta/internal/domain/tokens"
	"secretsanta/internal/infrastructure/delivery/api_http/auth"
	"secretsanta/internal/infrastructure/delivery/api_http/errors"
	"strings"

	"github.com/gofiber/fiber/v2"
)

const (
	HeaderAuthorization = "Authorization"
)

func buildAuthenticationMiddleware(tokenService tokens.TokensProvider) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get(HeaderAuthorization)
		if authHeader == "" {
			return c.Next()
		}

		authHeaderContent := strings.Split(authHeader, auth.TokenPrefix)
		if len(authHeaderContent) != 2 {
			return c.Status(http.StatusUnauthorized).JSON(errors.BuildHttpError(tokens.ErrInvalidToken))
		}

		claims, err := tokenService.VerifyAccessToken(c.Context(), authHeaderContent[1])
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(errors.BuildHttpError(err))
		}

		c.Locals(auth.ContextUserClaimsKey, *claims)
		return c.Next()
	}
}

func authRequiredMiddleware(c *fiber.Ctx) error {
	claims := c.Locals(auth.ContextUserClaimsKey)
	if claims == nil {
		return c.Status(http.StatusUnauthorized).JSON(errors.BuildHttpError(errors.ErrForbidden))
	}

	return c.Next()
}

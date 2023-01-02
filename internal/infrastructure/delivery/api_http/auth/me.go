package auth

import (
	"net/http"
	"secretsanta/internal/domain/user/entity"
	"secretsanta/internal/domain/user/usecase"
	"secretsanta/internal/infrastructure/delivery/api_http/errors"

	"github.com/gofiber/fiber/v2"
)

func NewMeHandler(uc usecase.UserMeUsecase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims := MustGetClaims(c)
		user, err := uc.Execute(c.Context(), entity.UserID(claims.UserID))
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(errors.BuildHttpError(err))
		}

		return c.JSON(UserROFromDoman(user))
	}
}

package auth

import (
	"net/http"
	"secretsanta/internal/domain/user/usecase"
	"secretsanta/internal/infrastructure/delivery/api_http/errors"

	"github.com/gofiber/fiber/v2"
)

type LoginRO struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	User         UserRO `json:"user"`
}

func NewLoginHandler(uc usecase.LoginUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		code := c.Query("code")
		tokens, err := uc.Execute(c.Context(), &usecase.LoginUserParams{
			Code: code,
		})
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(errors.BuildHttpError(err))
		}

		resp := LoginRO{
			AccessToken:  tokens.AccessToken,
			RefreshToken: tokens.RefreshToken,
			User:         UserROFromDoman(tokens.User),
		}
		return c.JSON(resp)
	}
}

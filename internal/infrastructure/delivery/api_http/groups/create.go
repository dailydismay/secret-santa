package groups

import (
	"encoding/json"
	"net/http"
	"secretsanta/internal/application/groups"
	"secretsanta/internal/domain/user/entity"
	"secretsanta/internal/infrastructure/delivery/api_http/auth"
	"secretsanta/internal/infrastructure/delivery/api_http/errors"

	"github.com/gofiber/fiber/v2"
)

type CreateGroupBody struct {
	Title string `json:"title"`
}

type CreateGroupResponse struct {
	ID string `json:"id"`
}

func NewCreateGroupHandler(
	createUC groups.CreateGroupUseCase,
) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body CreateGroupBody
		if err := json.Unmarshal(c.Body(), &body); err != nil {
			return c.Status(http.StatusUnprocessableEntity).JSON(errors.BuildHttpError(errors.ErrInvalidPayload))
		}

		claims := auth.MustGetClaims(c)
		result, err := createUC.Execute(c.Context(), &groups.CreateGroupPayload{
			Title:   body.Title,
			OwnerID: entity.UserID(claims.UserID),
		})
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(errors.BuildHttpError(err))
		}

		return c.JSON(&CreateGroupResponse{
			ID: string(result.Group.GetID()),
		})
	}
}

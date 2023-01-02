package groups

import (
	"net/http"
	"secretsanta/internal/application/groups"
	"secretsanta/internal/domain/group/entity"
	userDomain "secretsanta/internal/domain/user/entity"
	"secretsanta/internal/infrastructure/delivery/api_http/auth"
	"secretsanta/internal/infrastructure/delivery/api_http/errors"

	"github.com/gofiber/fiber/v2"
)

func NewGetGroupByIDHandler(getByIDUC groups.GetGroupByIDUsecase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		claims := auth.MustGetClaims(c)

		group, err := getByIDUC.Execute(c.Context(), &groups.GetGroupByIDPayload{
			ID:               entity.ID(id),
			AuthorizedUserID: userDomain.UserID(claims.UserID),
		})
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(errors.BuildHttpError(err))
		}

		return c.JSON(GroupDetailedFromDomain(group.Group))
	}
}

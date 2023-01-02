package groups

import (
	"context"
	"secretsanta/internal/domain/group/entity"
	userEntity "secretsanta/internal/domain/user/entity"
)

type CreateGroupPayload struct {
	Title   string
	OwnerID userEntity.UserID
}

type CreateGroupResult struct {
	Group *entity.Group
}

type CreateGroupUseCase interface {
	Execute(context.Context, *CreateGroupPayload) (*CreateGroupResult, error)
}

type GetGroupByIDPayload struct {
	ID               entity.ID
	AuthorizedUserID userEntity.UserID
}
type GetGroupByIDResult struct {
	Group *entity.GroupDetailed
}

type GetGroupByIDUsecase interface {
	Execute(context.Context, *GetGroupByIDPayload) (*GetGroupByIDResult, error)
}

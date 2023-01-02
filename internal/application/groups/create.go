package groups

import (
	"context"
	"secretsanta/internal/domain/group/entity"
	"secretsanta/internal/domain/group/persistence"
	userEntity "secretsanta/internal/domain/user/entity"

	"go.uber.org/fx"
)

type createGroupUseCase struct {
	groupRepo persistence.GroupRepository
}

type CreateGroupOptions struct {
	fx.In

	GroupRepo persistence.GroupRepository
}

func NewCreateGroupUseCase(opts CreateGroupOptions) CreateGroupUseCase {
	return &createGroupUseCase{
		groupRepo: opts.GroupRepo,
	}
}

func (i *createGroupUseCase) Execute(ctx context.Context, payload *CreateGroupPayload) (*CreateGroupResult, error) {
	newGroup := entity.New(
		entity.NewTitle(payload.Title),
		payload.OwnerID,
		[]userEntity.UserID{payload.OwnerID},
	)

	err := i.groupRepo.Add(ctx, newGroup)

	if err != nil {
		return nil, err
	}

	return &CreateGroupResult{
		Group: newGroup,
	}, nil
}

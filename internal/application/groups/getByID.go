package groups

import (
	"context"
	"secretsanta/internal/core"
	"secretsanta/internal/domain/group/persistence"

	"go.uber.org/fx"
)

type getGroupByIDUseCase struct {
	groupRepo persistence.GroupRepository
}

type GetByIDOptions struct {
	fx.In

	GroupRepo persistence.GroupRepository
}

func NewGetGroupByIDUseCase(opts GetByIDOptions) GetGroupByIDUsecase {
	return &getGroupByIDUseCase{
		groupRepo: opts.GroupRepo,
	}
}

func (i *getGroupByIDUseCase) Execute(ctx context.Context, params *GetGroupByIDPayload) (*GetGroupByIDResult, error) {
	isOwner, err := i.groupRepo.IsMember(ctx, params.ID, params.AuthorizedUserID)
	if err != nil {
		return nil, err
	}

	if !isOwner {
		return nil, core.NewNotEnoughRightsError("group", string(params.ID))
	}

	if group, err := i.groupRepo.GetByID(ctx, params.ID); err != nil {
		return nil, err
	} else {
		return &GetGroupByIDResult{
			Group: group,
		}, nil
	}
}

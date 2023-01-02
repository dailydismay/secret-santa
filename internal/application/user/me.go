package user

import (
	"context"
	"secretsanta/internal/domain/user/entity"
	"secretsanta/internal/domain/user/persistence"
	"secretsanta/internal/domain/user/usecase"

	"go.uber.org/fx"
)

type NonEmptyString string

type UserMeUseCaseOptions struct {
	fx.In
	UserRepo persistence.UserRepository
}

type meImpl struct {
	userRepo persistence.UserRepository
}

func NewMeUseCase(opts LoginUseCaseOptions) usecase.UserMeUsecase {
	return &meImpl{
		userRepo: opts.UserRepo,
	}
}

func (i *meImpl) Execute(ctx context.Context, id entity.UserID) (*entity.User, error) {
	return i.userRepo.GetByID(ctx, id)
}

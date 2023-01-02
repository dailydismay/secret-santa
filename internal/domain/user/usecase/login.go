package usecase

import (
	"context"
	"secretsanta/internal/domain/user/entity"
)

type LoginUserParams struct {
	Code string
}

type LoginUserResult struct {
	AccessToken  string
	RefreshToken string
	User         *entity.User
}

type LoginUseCase interface {
	Execute(context.Context, *LoginUserParams) (*LoginUserResult, error)
}

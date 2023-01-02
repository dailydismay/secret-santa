package usecase

import (
	"context"
	"secretsanta/internal/domain/user/entity"
)

type UserMeUsecase interface {
	Execute(context.Context, entity.UserID) (*entity.User, error)
}

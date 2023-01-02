package persistence

import (
	"context"
	"secretsanta/internal/domain/user/entity"
)

type UserRepository interface {
	GetByAuthProviderID(context.Context, string) (*entity.User, error)
	GetByID(context.Context, entity.UserID) (*entity.User, error)
	AddUser(context.Context, *entity.User) error
}

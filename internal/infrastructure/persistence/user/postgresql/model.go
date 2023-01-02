package postgresql

import (
	"secretsanta/internal/domain/user/entity"
	"time"
)

type Password string

func NewPassword(s string) (Password, error) {
	return Password(s), nil
}

type UserModel struct {
	ID             string `db:"id"`
	AuthProviderID string `db:"auth_provider_id"`
	FirstName      string `db:"first_name"`
	LastName       string `db:"last_name"`
	AvatarURL      string `db:"avatar_url"`
	Password       Password
	CreatedAt      time.Time `db:"created_at"`
}

func (um UserModel) ToDomain() *entity.User {
	return entity.BuildFromData(entity.UserID(um.ID), um.AuthProviderID, um.FirstName, um.LastName, um.AvatarURL, um.CreatedAt)
}

func fromDomain(e entity.User) UserModel {
	return UserModel{
		ID:             string(e.GetID()),
		AuthProviderID: e.GetAuthProviderID(),
		FirstName:      e.GetFirstName(),
		LastName:       e.GetLastName(),
		AvatarURL:      e.GetAvatarURL(),
		CreatedAt:      e.GetCreatedAt(),
	}
}

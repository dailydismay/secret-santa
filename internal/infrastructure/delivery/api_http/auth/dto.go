package auth

import (
	"secretsanta/internal/domain/user/entity"
	"time"
)

type UserRO struct {
	ID             string    `json:"id"`
	AuthProviderID string    `json:"auth_provider_id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	AvatarURL      string    `json:"avatar_url"`
	CreatedAt      time.Time `json:"created_at"`
}

func UserROFromDoman(user *entity.User) UserRO {
	return UserRO{
		ID:             string(user.GetID()),
		AuthProviderID: user.GetAuthProviderID(),
		FirstName:      user.GetFirstName(),
		LastName:       user.GetLastName(),
		AvatarURL:      user.GetAvatarURL(),
		CreatedAt:      user.GetCreatedAt(),
	}
}

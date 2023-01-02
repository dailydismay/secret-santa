package entity

import (
	"time"
)

type User struct {
	id             UserID
	authProviderID string
	firstName      string
	lastName       string
	avatarURL      string
	createdAt      time.Time
}

func New(authProviderID string, firstName string, lastName string, avatarURL string) *User {
	return &User{
		id:             NewUserID(),
		authProviderID: authProviderID,
		firstName:      firstName,
		lastName:       lastName,
		avatarURL:      avatarURL,
		createdAt:      time.Now(),
	}
}

func BuildFromData(id UserID, authProviderID string, firstName string, lastName string, avatarURL string, createdAt time.Time) *User {
	return &User{
		id:             id,
		authProviderID: authProviderID,
		firstName:      firstName,
		lastName:       lastName,
		avatarURL:      avatarURL,
		createdAt:      createdAt,
	}
}

func (u *User) GetID() UserID {
	return u.id
}

func (u *User) GetFirstName() string {
	return u.firstName
}

func (u *User) SetFirstName(firstName string) {
	u.firstName = firstName
}

func (u *User) GetLastName() string {
	return u.lastName
}

func (u *User) SetLastName(lastName string) {
	u.lastName = lastName
}

func (u *User) GetAvatarURL() string {
	return u.avatarURL
}

func (u *User) SetAvatarURL(avatarURL string) {
	u.avatarURL = avatarURL
}

func (u *User) GetAuthProviderID() string {
	return u.authProviderID
}

func (u *User) SetAuthProviderID(authID string) {
	u.authProviderID = authID
}

func (u *User) GetCreatedAt() time.Time {
	return u.createdAt
}

package entity

import userEntity "secretsanta/internal/domain/user/entity"

type GroupMember struct {
	id        userEntity.UserID
	firstName string
	lastName  string
	avatarURL string
	isOwner   bool
}

func NewGroupMember(
	id userEntity.UserID,
	firstName string,
	lastName string,
	avatarURL string,
	isOwner bool,
) *GroupMember {
	return &GroupMember{
		id:        id,
		firstName: firstName,
		lastName:  lastName,
		avatarURL: avatarURL,
		isOwner:   isOwner,
	}
}

func (gm *GroupMember) GetID() userEntity.UserID {
	return gm.id
}
func (gm *GroupMember) GetFirstName() string {
	return gm.firstName
}
func (gm *GroupMember) GetLastName() string {
	return gm.lastName
}
func (gm *GroupMember) GetAvatarURL() string {
	return gm.avatarURL
}
func (gm *GroupMember) GetIsOwner() bool {
	return gm.isOwner
}

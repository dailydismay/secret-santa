package entity

import (
	userEntity "secretsanta/internal/domain/user/entity"
	"secretsanta/pkg/util"
	"time"
)

type Group struct {
	id             ID
	title          Title
	invitationCode string
	createdAt      time.Time
	ownerID        userEntity.UserID
	membersIDs     []userEntity.UserID
}

type GroupDetailed struct {
	Group
	members []*GroupMember
}

func (gd *GroupDetailed) GetDetailedMembers() []*GroupMember {
	return gd.members
}

func New(
	title Title,
	ownerID userEntity.UserID,
	membersIDs []userEntity.UserID,
) *Group {
	return &Group{
		id:             NewID(),
		title:          title,
		invitationCode: util.GetRandomString(6),
		createdAt:      time.Now(),
		ownerID:        ownerID,
		membersIDs:     membersIDs,
	}
}

func FromDataDetailed(
	id ID,
	title Title,
	invitationCode string,
	createdAt time.Time,
	ownerID userEntity.UserID,
	members []*GroupMember,
) *GroupDetailed {
	return &GroupDetailed{
		Group: Group{
			id:             id,
			title:          title,
			invitationCode: invitationCode,
			createdAt:      createdAt,
			ownerID:        ownerID,
		},
		members: members,
	}
}

func (g *Group) GetID() ID {
	return g.id
}

func (g *Group) GetTitle() Title {
	return g.title
}

func (g *Group) GetInvitationCode() string {
	return g.invitationCode
}

func (g *Group) SetInvitationCode(s string) {
	g.invitationCode = s
}

func (g *Group) GetCreatedAt() time.Time {
	return g.createdAt
}

func (g *Group) GetOwnerID() userEntity.UserID {
	return g.ownerID
}

func (g *Group) GetMembers() []userEntity.UserID {
	return g.membersIDs
}

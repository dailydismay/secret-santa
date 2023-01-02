package groups

import (
	"secretsanta/internal/domain/group/entity"
	"time"
)

type GroupRO struct {
	ID             string    `json:"id"`
	Title          string    `json:"title"`
	InvitationCode string    `json:"invitation_code"`
	CreatedAt      time.Time `json:"created_at"`
	OwnerID        string    `json:"owner_id"`
}

func GroupFromDomain(g *entity.Group) *GroupRO {
	return &GroupRO{
		ID:             string(g.GetID()),
		Title:          string(g.GetTitle()),
		InvitationCode: string(g.GetInvitationCode()),
	}
}

type GroupMemberRO struct {
	ID        string `json:"id"`
	FirstName string `json:"fisrt_name"`
	LastName  string `json:"last_name"`
	AvatarURL string `json:"avatar_url"`
	IsOwner   bool   `json:"is_owner"`
}

type GroupDetailedRO struct {
	ID             string           `json:"id"`
	Title          string           `json:"title"`
	InvitationCode string           `json:"invitation_code"`
	CreatedAt      time.Time        `json:"created_at"`
	OwnerID        string           `json:"owner_id"`
	Members        []*GroupMemberRO `json:"members"`
}

func GroupMemberFromDomain(m *entity.GroupMember) *GroupMemberRO {
	return &GroupMemberRO{
		ID:        string(m.GetID()),
		FirstName: m.GetFirstName(),
		LastName:  m.GetLastName(),
		IsOwner:   m.GetIsOwner(),
		AvatarURL: m.GetAvatarURL(),
	}
}

func GroupMembersFromDomain(mm []*entity.GroupMember) []*GroupMemberRO {
	r := make([]*GroupMemberRO, 0)

	for _, m := range mm {
		r = append(r, GroupMemberFromDomain(m))
	}

	return r
}

func GroupDetailedFromDomain(gd *entity.GroupDetailed) *GroupDetailedRO {
	return &GroupDetailedRO{
		ID:             string(gd.GetID()),
		Title:          string(gd.GetTitle()),
		InvitationCode: string(gd.GetInvitationCode()),
		CreatedAt:      gd.GetCreatedAt(),
		Members:        GroupMembersFromDomain(gd.GetDetailedMembers()),
	}
}

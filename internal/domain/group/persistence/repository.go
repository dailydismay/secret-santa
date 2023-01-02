package persistence

import (
	"context"
	"secretsanta/internal/domain/group/entity"
	userEntity "secretsanta/internal/domain/user/entity"
)

type CreateGroupPayload struct{}

type ListForUserPayload struct {
	VisitorID userEntity.UserID
	Limit     int
	Offset    int
}

type ListForUserResult struct {
	Items   entity.Group
	Total   int
	HasNext bool
}

type GroupRepository interface {
	GetByID(context.Context, entity.ID) (*entity.GroupDetailed, error)
	ListForUser(context.Context, *ListForUserPayload) (*ListForUserResult, error)
	AddMemberByInvitationCode(ctx context.Context, id entity.ID, userID userEntity.UserID) error
	IsMember(context.Context, entity.ID, userEntity.UserID) (bool, error)
	Add(context.Context, *entity.Group) error
}

//
// select count(*) from group_members gm where gm.id = $1 LIMIT 1

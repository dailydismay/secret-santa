package postgresql

import (
	"secretsanta/internal/domain/group/entity"
	userEntity "secretsanta/internal/domain/user/entity"
	"time"
)

type Group struct {
	ID             string    `db:"id"`
	Title          string    `db:"title"`
	InvitationCode string    `db:"title"`
	OwnerID        string    `db:"owner_id"`
	CreatedAt      time.Time `db:"created_at"`
}

type GroupMemberToInsert struct {
	GroupID string `db:"group_id"`
	UserID  string `db:"user_id"`
}

func MembersToDataInsert(id entity.ID, ids []userEntity.UserID) []GroupMemberToInsert {
	e := make([]GroupMemberToInsert, 0)
	groupID := string(id)

	for _, id := range ids {
		e = append(e, GroupMemberToInsert{
			GroupID: groupID,
			UserID:  string(id),
		})
	}

	return e
}

type GroupMember struct {
	ID        string `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	AvatarURL string `db:"avatar_url"`
	IsOwner   bool   `db:"is_owner"`
}

func (gm *GroupMember) ToDomain() *entity.GroupMember {
	return entity.NewGroupMember(
		userEntity.UserID(gm.ID),
		gm.FirstName,
		gm.LastName,
		gm.AvatarURL,
		gm.IsOwner,
	)
}

type JoinedGroupRow struct {
	ID             string    `db:"group_id"`
	Title          string    `db:"title"`
	InvitationCode string    `db:"invitation_code"`
	OwnerID        string    `db:"owner_id"`
	CreatedAt      time.Time `db:"created_at"`
	MemberID       string    `db:"id"`
	FirstName      string    `db:"first_name"`
	LastName       string    `db:"last_name"`
	AvatarURL      string    `db:"avatar_url"`
	IsOwner        bool      `db:"is_owner"`
}

type GroupWithMembers struct {
	ID             string         `db:"id"`
	Title          string         `db:"title"`
	InvitationCode string         `db:"invitation_code"`
	OwnerID        string         `db:"owner_id"`
	CreatedAt      time.Time      `db:"created_at"`
	Members        []*GroupMember `db:",prefix=u"`
}

func mapMembers(mm []*GroupMember) []*entity.GroupMember {
	e := make([]*entity.GroupMember, len(mm))

	for _, m := range mm {
		e = append(e, m.ToDomain())
	}

	return e
}

func (gwm *GroupWithMembers) ToDomainDetailed() *entity.GroupDetailed {
	return entity.FromDataDetailed(
		entity.ID(gwm.ID),
		entity.Title(gwm.Title),
		gwm.InvitationCode,
		gwm.CreatedAt,
		userEntity.UserID(gwm.OwnerID),
		mapMembers(gwm.Members),
	)
}

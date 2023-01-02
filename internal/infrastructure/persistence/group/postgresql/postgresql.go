package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"secretsanta/internal/core"
	"secretsanta/internal/domain/group/entity"
	"secretsanta/internal/domain/group/persistence"
	userEntity "secretsanta/internal/domain/user/entity"
	"secretsanta/internal/infrastructure/crosscutting/postgresql"

	"go.uber.org/fx"
)

type Options struct {
	fx.In

	Pg *postgresql.PostgresAdapter
}

type implementation struct {
	pg *postgresql.PostgresAdapter
}

func NewPostgresGroupRepo(opts Options) persistence.GroupRepository {
	return &implementation{
		pg: opts.Pg,
	}
}

func (i implementation) ListForUser(ctx context.Context, p *persistence.ListForUserPayload) (*persistence.ListForUserResult, error) {
	return nil, nil
}

func (i implementation) GetByID(ctx context.Context, id entity.ID) (*entity.GroupDetailed, error) {
	var rows []JoinedGroupRow
	err := i.pg.DB.SelectContext(ctx, &rows, `
		SELECT 
			g.id as group_id,
			g.title,
			g.invitation_code,
			g.created_at, 
			u.id,
			u.first_name,
			u.last_name,
			u.avatar_url,
			u.id = g.owner_id as is_owner
		FROM groups g
		LEFT JOIN group_members gm ON gm.group_id = g.id
		LEFT JOIN users u ON u.id = gm.user_id
		WHERE g.id = $1;
	`, id)

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, core.NewEntityNotFoundError("group", "id")
	} else if err != nil {
		return nil, err
	}

	var group *GroupWithMembers
	for _, r := range rows {
		if group == nil {
			group = &GroupWithMembers{
				ID:             r.ID,
				Title:          r.Title,
				InvitationCode: r.InvitationCode,
				OwnerID:        r.OwnerID,
				CreatedAt:      r.CreatedAt,
				Members:        []*GroupMember{},
			}
		}

		group.Members = append(group.Members, &GroupMember{
			ID:        r.MemberID,
			FirstName: r.FirstName,
			LastName:  r.LastName,
			AvatarURL: r.AvatarURL,
			IsOwner:   r.IsOwner,
		})
	}

	return group.ToDomainDetailed(), nil
}

func (i *implementation) IsMember(ctx context.Context, id entity.ID, userID userEntity.UserID) (bool, error) {
	var digit int

	err := i.pg.DB.GetContext(ctx, &digit, "SELECT 1 FROM group_members gm WHERE gm.user_id = $1 AND gm.group_id $2 LIMIT 1")
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

func (i *implementation) AddMemberByInvitationCode(ctx context.Context, id entity.ID, userID userEntity.UserID) error {
	return nil
}

func (i *implementation) Add(ctx context.Context, e *entity.Group) error {
	tx, err := i.pg.DB.Beginx()
	if err != nil {
		return err
	}
	var count int
	rows, err := tx.QueryContext(ctx, "SELECT count(*) as cnt FROM groups WHERE title = $1 OR invitation_code = $2", string(e.GetTitle()), e.GetInvitationCode())
	if err != nil {
		return err
	}

	if rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return err
		}

		if count != 0 {
			if err := tx.Rollback(); err != nil {
				return err
			}
			return core.ErrResourceDuplication
		}
	}
	rows.Close()

	_, err = tx.ExecContext(ctx, "INSERT INTO groups (id, title, invitation_code, owner_id, created_at) VALUES ($1, $2, $3, $4, $5)",
		e.GetID(),
		e.GetTitle(),
		e.GetInvitationCode(),
		e.GetOwnerID(),
		e.GetCreatedAt(),
	)
	if err != nil {
		return err
	}
	_, err = tx.NamedExecContext(ctx, "INSERT INTO group_members (user_id, group_id) VALUES (:user_id, :group_id)",
		MembersToDataInsert(e.GetID(), e.GetMembers()),
	)
	if err != nil {
		return err
	}

	return tx.Commit()
}

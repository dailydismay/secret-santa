package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"secretsanta/internal/domain/user/entity"
	"secretsanta/internal/domain/user/persistence"
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

func NewPostgresUserRepo(opts Options) persistence.UserRepository {
	return &implementation{
		pg: opts.Pg,
	}
}

func (i *implementation) GetByAuthProviderID(ctx context.Context, authProviderID string) (*entity.User, error) {
	var user UserModel
	err := i.pg.DB.GetContext(ctx, &user, `SELECT * FROM users WHERE auth_provider_id = $1;`, authProviderID)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, persistence.ErrUserNotFound
	} else if err != nil {
		return nil, err
	}
	return user.ToDomain(), nil
}

func (i *implementation) GetByID(ctx context.Context, id entity.UserID) (*entity.User, error) {
	var user UserModel
	err := i.pg.DB.GetContext(ctx, &user, `SELECT * FROM users WHERE id=$1;`, id)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, persistence.ErrUserNotFound
	} else if err != nil {
		return nil, err
	}

	return user.ToDomain(), nil
}

func (i *implementation) AddUser(ctx context.Context, user *entity.User) error {
	tx, err := i.pg.DB.Begin()
	if err != nil {
		return err
	}

	var existingUser UserModel
	row := tx.QueryRowContext(ctx, "SELECT * FROM users WHERE auth_provider_id = $1", user.GetAuthProviderID())
	if err := row.Scan(&existingUser); err != nil && !errors.Is(err, sql.ErrNoRows) {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	if existingUser.ID != "" {
		if err := tx.Rollback(); err != nil {
			return err
		}

		return persistence.ErrUserExists
	}

	usr := fromDomain(*user)
	_, err = tx.ExecContext(ctx, `
		INSERT INTO users(id, auth_provider_id, first_name, last_name, avatar_url, created_at)
		VALUES($1, $2, $3, $4, $5, $6);
	`, usr.ID, usr.AuthProviderID, usr.FirstName, usr.LastName, usr.AvatarURL, usr.CreatedAt)
	if err != nil {
		return err
	}

	return tx.Commit()
}

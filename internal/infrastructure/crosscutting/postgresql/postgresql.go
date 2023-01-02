package postgresql

import (
	"context"
	"secretsanta/internal/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
)

type PostgresAdapter struct {
	DB *sqlx.DB
}

type PostgresOptions struct {
	fx.In
	LC  fx.Lifecycle
	Cfg config.Config
}

func NewPostgresAdapter(opts PostgresOptions) (*PostgresAdapter, error) {
	adapter := &PostgresAdapter{}
	db, err := sqlx.ConnectContext(context.Background(), "postgres", opts.Cfg.PGConnectionString)
	if err != nil {
		return nil, err
	}
	adapter.DB = db

	// opts.LC.Append(fx.Hook{
	// 	OnStart: func(c context.Context) error {

	// 		return db.PingContext(c)
	// 	},
	// 	OnStop: func(c context.Context) error {
	// 		return adapter.DB.Close()
	// 	},
	// })

	return adapter, nil
}

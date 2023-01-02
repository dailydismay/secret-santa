package main

import (
	"context"
	"secretsanta/internal/application/groups"
	"secretsanta/internal/application/user"
	"secretsanta/internal/config"
	"secretsanta/internal/infrastructure/crosscutting/postgresql"
	"secretsanta/internal/infrastructure/delivery/api_http"
	groupsPostgresRepo "secretsanta/internal/infrastructure/persistence/group/postgresql"
	userPostgresRepo "secretsanta/internal/infrastructure/persistence/user/postgresql"
	"secretsanta/internal/infrastructure/providers/codeflowauth/vk"
	"secretsanta/internal/infrastructure/providers/tokens/paseto"

	"go.uber.org/fx"
)

func main() {
	appCtx, cancel := context.WithCancel(context.Background())
	app := fx.New(
		fx.Provide(func() context.Context {
			return appCtx
		}),
		fx.Provide(config.New),
		fx.Provide(postgresql.NewPostgresAdapter),
		fx.Provide(userPostgresRepo.NewPostgresUserRepo),
		fx.Provide(groupsPostgresRepo.NewPostgresGroupRepo),
		fx.Provide(
			paseto.NewPasetoTokenService,
		),
		fx.Provide(
			vk.NewVKCodeFlowAuthProvider,
		),
		fx.Provide(
			groups.NewCreateGroupUseCase,
			groups.NewGetGroupByIDUseCase,
		),
		fx.Provide(
			user.NewLoginUseCase,
			user.NewMeUseCase,
		),
		fx.Invoke(
			api_http.NewAPIHttp,
		),
	)

	go func() {
		<-app.Done()
		cancel()
	}()

	app.Run()
}

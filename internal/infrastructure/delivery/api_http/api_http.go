package api_http

import (
	"context"
	"fmt"
	applicationGroups "secretsanta/internal/application/groups"
	"secretsanta/internal/config"
	"secretsanta/internal/domain/tokens"
	"secretsanta/internal/domain/user/usecase"
	"secretsanta/internal/infrastructure/delivery/api_http/auth"
	"secretsanta/internal/infrastructure/delivery/api_http/groups"

	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/fx"
)

type APIHttpAdapter struct {
	app            *fiber.App
	cfg            config.Config
	loginUC        usecase.LoginUseCase
	meUC           usecase.UserMeUsecase
	createGroupUC  applicationGroups.CreateGroupUseCase
	getGroupByIDUC applicationGroups.GetGroupByIDUsecase
	tokensProvider tokens.TokensProvider
}

type APIHttpOptions struct {
	fx.In

	LC             fx.Lifecycle
	Cfg            config.Config
	LoginUC        usecase.LoginUseCase
	MeUC           usecase.UserMeUsecase
	CreateGroupUC  applicationGroups.CreateGroupUseCase
	GetGroupByIDUC applicationGroups.GetGroupByIDUsecase

	TokensProvider tokens.TokensProvider
}

func (a *APIHttpAdapter) start(_ context.Context) error {
	go func() {
		err := a.app.Listen(fmt.Sprintf("localhost:%d", a.cfg.Port))
		panic(err)
	}()

	return nil
}
func (a *APIHttpAdapter) stop(_ context.Context) error {
	return a.app.Server().Shutdown()
}

func NewAPIHttp(opts APIHttpOptions) (APIHttpAdapter, error) {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	adapter := &APIHttpAdapter{
		app:            app,
		cfg:            opts.Cfg,
		loginUC:        opts.LoginUC,
		meUC:           opts.MeUC,
		createGroupUC:  opts.CreateGroupUC,
		getGroupByIDUC: opts.GetGroupByIDUC,
		tokensProvider: opts.TokensProvider,
	}

	adapter.mapHandlers()

	opts.LC.Append(fx.Hook{
		OnStart: adapter.start,
		OnStop:  adapter.stop,
	})

	return *adapter, nil
}

func (a *APIHttpAdapter) mapHandlers() {
	authMiddleware := buildAuthenticationMiddleware(a.tokensProvider)

	v1 := a.app.Group("/v1")

	authV1 := v1.Group("/auth")
	authV1.Get("/login", auth.NewLoginHandler(a.loginUC))
	authV1.Get("/me", authMiddleware, authRequiredMiddleware, auth.NewMeHandler(a.meUC))

	groupsV1 := v1.Group("/groups", authMiddleware, authRequiredMiddleware)
	// groupsV1.Get("/")
	groupsV1.Get("/:id", groups.NewGetGroupByIDHandler(a.getGroupByIDUC))
	groupsV1.Post("/", groups.NewCreateGroupHandler(a.createGroupUC))
}

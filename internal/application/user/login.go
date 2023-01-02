package user

import (
	"context"
	"errors"
	"fmt"
	"secretsanta/internal/config"
	"secretsanta/internal/domain/codeflowauth"
	"secretsanta/internal/domain/tokens"
	"secretsanta/internal/domain/user/entity"
	"secretsanta/internal/domain/user/persistence"
	"secretsanta/internal/domain/user/usecase"

	"go.uber.org/fx"
)

type LoginUseCaseOptions struct {
	fx.In
	Cfg                  config.Config
	UserRepo             persistence.UserRepository
	CodeFlowAuthProvider codeflowauth.CodeFlowAuthProvider
	TokensProvider       tokens.TokensProvider
}

type implementation struct {
	userRepo             persistence.UserRepository
	codeflowauthProvider codeflowauth.CodeFlowAuthProvider
	tokensProvider       tokens.TokensProvider
}

func NewLoginUseCase(opts LoginUseCaseOptions) usecase.LoginUseCase {
	return &implementation{
		userRepo:             opts.UserRepo,
		codeflowauthProvider: opts.CodeFlowAuthProvider,
		tokensProvider:       opts.TokensProvider,
	}
}

func (i *implementation) getOrCreateUserByAuthID(ctx context.Context, profile *codeflowauth.UserProfile) (*entity.User, error) {
	existingUser, err := i.userRepo.GetByAuthProviderID(ctx, string(profile.GetID()))
	if err != nil && !errors.Is(err, persistence.ErrUserNotFound) {
		return nil, err
	}

	if existingUser != nil {
		return existingUser, nil
	}

	created := entity.New(
		fmt.Sprint(profile.GetID()),
		profile.GetFirstName(),
		profile.GetLastName(),
		profile.GetPhoto(),
	)

	err = i.userRepo.AddUser(
		ctx,
		created,
	)
	if err != nil {
		return nil, err
	}

	return created, nil
}

func (i *implementation) Execute(ctx context.Context, params *usecase.LoginUserParams) (*usecase.LoginUserResult, error) {
	externalUserProfile, err := i.codeflowauthProvider.GetUserProfile(ctx, params.Code)
	if err != nil {
		return nil, err
	}

	usr, err := i.getOrCreateUserByAuthID(ctx, externalUserProfile)
	if err != nil {
		return nil, err
	}

	tkPair, err := i.tokensProvider.SignTokens(ctx, &tokens.SignTokensPayload{
		UserID: string(usr.GetID()),
	})
	if err != nil {
		return nil, err
	}

	return &usecase.LoginUserResult{
		AccessToken:  tkPair.AccessToken,
		RefreshToken: tkPair.RefreshToken,
		User:         usr,
	}, nil
}

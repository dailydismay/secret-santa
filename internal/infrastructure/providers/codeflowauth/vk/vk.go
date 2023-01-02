package vk

import (
	"context"
	"net/http"
	"secretsanta/internal/config"
	"secretsanta/internal/domain/codeflowauth"

	vkAPI "github.com/go-vk-api/vk"
	"go.uber.org/fx"
	"golang.org/x/oauth2"
	vkAuth "golang.org/x/oauth2/vk"
)

type implementation struct {
	cfg config.Config

	httpClient   *http.Client
	oauth2Config *oauth2.Config
}

type VKCodeFlowAuthProviderOptions struct {
	fx.In
	Config config.Config
}

func NewVKCodeFlowAuthProvider(opts VKCodeFlowAuthProviderOptions) codeflowauth.CodeFlowAuthProvider {
	oauthConf := &oauth2.Config{
		ClientID:     opts.Config.ClientID,
		ClientSecret: opts.Config.ClientSecret,
		RedirectURL:  opts.Config.RedirectURL,
		Scopes:       []string{},
		Endpoint:     vkAuth.Endpoint,
	}

	return &implementation{
		cfg:          opts.Config,
		oauth2Config: oauthConf,
		httpClient: &http.Client{
			Timeout: vkAPIRequestTimeout,
		},
	}
}

func (i *implementation) GetUserProfile(ctx context.Context, code string) (*codeflowauth.UserProfile, error) {
	ctx, cancel := context.WithTimeout(ctx, vkAPIRequestTimeout)
	defer cancel()
	tk, err := i.oauth2Config.Exchange(ctx, code)
	if err != nil {
		return nil, codeflowauth.ErrCodeInvalid
	}

	vkClient, err := vkAPI.NewClientWithOptions(vkAPI.WithToken(tk.AccessToken), vkAPI.WithHTTPClient(i.httpClient))
	if err != nil {
		return nil, codeflowauth.ErrTokenInvalid
	}

	var resp []ProfileResponse
	err = vkClient.CallMethod("users.get", vkAPI.RequestParams{
		"fields": "photo_400_orig",
	}, &resp)
	if err != nil {
		return nil, codeflowauth.ErrFailedToFetchProfile
	}

	return resp[0].ToDomain(), nil
}

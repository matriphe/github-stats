package repository

import (
	"context"

	"github.com/google/go-github/v35/github"
	"golang.org/x/oauth2"
)

type (
	authClient struct {
		ctx context.Context
	}

	AuthClient interface {
		Client(token string) *github.Client
	}
)

func NewAuthClient(ctx context.Context) AuthClient {
	return &authClient{ctx: ctx}
}

func (a *authClient) Client(token string) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(a.ctx, ts)

	return github.NewClient(tc)
}

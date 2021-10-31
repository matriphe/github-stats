package client

import (
	"context"

	"github.com/google/go-github/v35/github"
	"golang.org/x/oauth2"
)

type (
	authClient struct {
		ctx context.Context
	}

	// AuthClient is an interface for GitHub authenticated client.
	AuthClient interface {
		Client(token string) *github.Client
	}
)

// NewGitHubAuthClient creates new authenticated GitHub client.
func NewGitHubAuthClient(ctx context.Context, token string) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc)
}

// NewAuthClient creates new GitHub authenticated client.
func NewAuthClient(ctx context.Context) AuthClient {
	return &authClient{ctx: ctx}
}

// Client returns GitHub authenticate client.
func (a *authClient) Client(token string) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(a.ctx, ts)

	return github.NewClient(tc)
}

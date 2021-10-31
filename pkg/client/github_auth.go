package client

import (
	"context"

	"github.com/google/go-github/v35/github"
	"golang.org/x/oauth2"
)

// NewGitHubAuthClient creates new authenticated GitHub client.
func NewGitHubAuthClient(ctx context.Context, token string) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc)
}

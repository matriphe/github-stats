package repository

import (
	"context"

	"github.com/google/go-github/v35/github"
)

type (
	userRepo struct {
		ctx context.Context
		gc  *github.Client
	}

	// UserRepo is an interface for GitHub user repository.
	UserRepo interface {
		User() (*github.User, error)
	}
)

// NewUserRepo creates new GitHub user repository.
func NewUserRepo(ctx context.Context, gh *github.Client) UserRepo {
	return &userRepo{
		ctx: ctx,
		gc:  gh,
	}
}

// User returns GitHub user info.
func (r *userRepo) User() (*github.User, error) {
	u, _, err := r.gc.Users.Get(r.ctx, "")
	if err != nil {
		return nil, err
	}

	return u, nil
}

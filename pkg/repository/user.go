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

	UserRepo interface {
		User() (*github.User, error)
	}
)

func NewUserRepo(ctx context.Context, gh *github.Client) UserRepo {
	return &userRepo{
		ctx: ctx,
		gc:  gh,
	}
}

// User returns Github user info.
func (r *userRepo) User() (*github.User, error) {
	u, _, err := r.gc.Users.Get(r.ctx, "")
	if err != nil {
		return nil, err
	}

	return u, nil
}

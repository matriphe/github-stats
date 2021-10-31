//go:build unit

package repository

import (
	"github.com/google/go-github/v35/github"
	"github.com/stretchr/testify/mock"
)

// UserRepoMock mocks UserRepo.
type UserRepoMock struct {
	mock.Mock
}

// User mocks User.
func (m *UserRepoMock) User() (*github.User, error) {
	args := m.Called()

	err := args.Error(1)
	if err != nil {
		return nil, err
	}

	return args.Get(0).(*github.User), err
}

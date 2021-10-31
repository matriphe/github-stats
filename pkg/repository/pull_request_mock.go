//go:build unit

package repository

import "github.com/stretchr/testify/mock"

// PullRequestRepoMock mocks PullRequestRepo.
type PullRequestRepoMock struct {
	mock.Mock
}

// GetPullRequests mocks GetPullRequests.
func (m *PullRequestRepoMock) GetPullRequests(q PullRequestQuery) ([]PullRequests, error) {
	args := m.Called(q)

	return args.Get(0).([]PullRequests), args.Error(1)
}

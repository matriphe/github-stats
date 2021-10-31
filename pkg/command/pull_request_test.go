//go:build unit

package command

import (
	"testing"

	"github.com/google/go-github/v35/github"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/matriphe/github-stats/pkg/calculator"
	"github.com/matriphe/github-stats/pkg/repository"
)

func TestPullRequestCommand_NoError(t *testing.T) {
	testCases := map[string]struct {
		user           string
		stubPRs        []repository.PullRequests
		expectedResult calculator.PullRequestCalc
	}{
		"no pull request": {
			user:    "tester",
			stubPRs: []repository.PullRequests{},
			expectedResult: calculator.PullRequestCalc{
				Num: 0,
				Avg: calculator.PrAvg{
					Files:     0,
					Additions: 0,
					Deletions: 0,
					Changes:   0,
					Total:     0,
				},
			},
		},
		"two pull requests": {
			user: "tester",
			stubPRs: []repository.PullRequests{
				{
					Stats: repository.Stats{
						NumFiles:  2,
						Additions: 5,
						Deletions: 10,
						Changes:   20,
						Total:     35,
					},
				},
				{
					Stats: repository.Stats{
						NumFiles:  1,
						Additions: 10,
						Deletions: 5,
						Changes:   25,
						Total:     40,
					},
				},
			},
			expectedResult: calculator.PullRequestCalc{
				Num: 2,
				Avg: calculator.PrAvg{
					Files:     1,
					Additions: 7,
					Deletions: 7,
					Changes:   22,
					Total:     37,
				},
			},
		},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			userRepo := &repository.UserRepoMock{}
			defer userRepo.AssertExpectations(t)

			userRepo.On("User").
				Once().
				Return(&github.User{Login: &tc.user}, nil)

			prRepo := &repository.PullRequestRepoMock{}
			defer prRepo.AssertExpectations(t)

			prQuery := repository.PullRequestQuery{}

			prRepo.On("GetPullRequests", repository.PullRequestQuery{Author: tc.user}).
				Once().
				Return(tc.stubPRs, nil)

			calc := calculator.NewPullRequestCalculator()

			result, err := PullRequest(userRepo, prRepo, prQuery, calc)

			require.NoError(t, err)
			assert.Equal(t, tc.expectedResult, result.Statistics)
		})
	}
}

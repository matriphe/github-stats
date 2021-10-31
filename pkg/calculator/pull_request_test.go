//go:build unit

package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matriphe/github-stats/pkg/repository"
)

func TestNewPullRequestCalculator(t *testing.T) {
	testCases := map[string]struct {
		pullRequests   []repository.PullRequests
		expectedResult PullRequestCalc
	}{
		"no pull request": {
			pullRequests: []repository.PullRequests{},
			expectedResult: PullRequestCalc{
				Num: 0,
				Avg: PrAvg{
					Files:     0,
					Additions: 0,
					Deletions: 0,
					Changes:   0,
					Total:     0,
				},
			},
		},
		"two pull requests": {
			pullRequests: []repository.PullRequests{
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
			expectedResult: PullRequestCalc{
				Num: 2,
				Avg: PrAvg{
					Files:     1,
					Additions: 7,
					Deletions: 7,
					Changes:   22,
					Total:     37,
				},
			},
		},
	}

	calc := NewPullRequestCalculator()

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			result := calc.GetAverage(tc.pullRequests)

			assert.Equal(t, tc.expectedResult, result)
		})
	}
}

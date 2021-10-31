package calculator

import "github.com/matriphe/github-stats/pkg/repository"

type (
	prCalc struct {
	}

	// PrAvg holds the pull request average.
	PrAvg struct {
		Files     int
		Additions int
		Deletions int
		Changes   int
		Total     int
	}

	// PullRequestCalc holds pull request calculator result.
	PullRequestCalc struct {
		Num int
		Avg PrAvg
	}

	// PullRequestCalculator is an interface for pull request calculator.
	PullRequestCalculator interface {
		GetAverage([]repository.PullRequests) PullRequestCalc
	}
)

// NewPullRequestCalculator creates pull request calculator.
func NewPullRequestCalculator() PullRequestCalculator {
	return &prCalc{}
}

// GetAverage returns pull request average.
func (c *prCalc) GetAverage(prs []repository.PullRequests) PullRequestCalc {
	r := PullRequestCalc{Num: len(prs)}
	if r.Num == 0 {
		return r
	}

	additions := 0
	deletions := 0
	changes := 0
	total := 0
	files := 0

	for _, pr := range prs {
		additions += pr.Stats.Additions
		deletions += pr.Stats.Deletions
		changes += pr.Stats.Changes
		total += pr.Stats.Total
		files += pr.Stats.NumFiles
	}

	r.Avg = PrAvg{
		Files:     files / r.Num,
		Additions: additions / r.Num,
		Deletions: deletions / r.Num,
		Changes:   changes / r.Num,
		Total:     total / r.Num,
	}

	return r
}

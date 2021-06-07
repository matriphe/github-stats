package calculator

import "github.com/matriphe/github-stats/pkg/repository"

type (
	prCalc struct {
	}

	prAvg struct {
		Files     int
		Additions int
		Deletions int
		Changes   int
		Total     int
	}

	PullRequestCalc struct {
		Num int
		Avg prAvg
	}

	PullRequestCalculator interface {
		GetAverage([]repository.PullRequests) PullRequestCalc
	}
)

func NewPullRequestCalculator() PullRequestCalculator {
	return &prCalc{}
}

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

	r.Avg = prAvg{
		Files:     files / r.Num,
		Additions: additions / r.Num,
		Deletions: deletions / r.Num,
		Changes:   changes / r.Num,
		Total:     total / r.Num,
	}

	return r
}

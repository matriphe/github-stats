package command

import (
	"github.com/google/go-github/v35/github"
	"github.com/pkg/errors"

	"github.com/matriphe/github-stats/pkg/calculator"
	"github.com/matriphe/github-stats/pkg/repository"
)

type (
	PullRequestCommandResult struct {
		User         *github.User
		PullRequests []repository.PullRequests
		Query        repository.PullRequestQuery
		Statistics   calculator.PullRequestCalc
	}
)

func PullRequest(
	userRepo repository.UserRepo,
	prRepo repository.PullRequestRepo,
	query repository.PullRequestQuery,
	calc calculator.PullRequestCalculator,
) (PullRequestCommandResult, error) {
	result := PullRequestCommandResult{
		Query: query,
	}

	user, err := userRepo.User()
	if err != nil {
		return result, errors.Wrap(err, "failed getting user info")
	}

	result.User = user
	result.Query.Author = result.User.GetLogin()

	prs, err := prRepo.GetPullRequests(result.Query)
	if err != nil {
		return result, errors.Wrap(err, "failed getting PR info")
	}

	result.PullRequests = prs
	result.Statistics = calc.GetAverage(result.PullRequests)

	return result, nil
}

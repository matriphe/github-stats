package repository

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-github/v35/github"
)

type (
	prRepo struct {
		ctx     context.Context
		gh      *github.Client
		perPage int
	}

	ownerRepo struct {
		owner string
		repo  string
	}

	PullRequestQuery struct {
		Author    string
		Org       string
		State     string
		StartDate string
	}

	Stats struct {
		NumFiles  int
		Additions int
		Deletions int
		Changes   int
		Total     int
	}

	PullRequests struct {
		Issue *github.Issue
		Files []*github.CommitFile
		Stats Stats
	}

	PullRequestRepo interface {
		GetPullRequests(PullRequestQuery) ([]PullRequests, error)
	}
)

func NewPullRequestRepo(
	ctx context.Context,
	client *github.Client,
	perPage int,
) PullRequestRepo {
	return &prRepo{
		ctx:     ctx,
		gh:      client,
		perPage: perPage,
	}
}

func (pr *prRepo) GetPullRequests(query PullRequestQuery) ([]PullRequests, error) {
	q := pr.setQuery(query)

	issues, err := pr.getIssues(q)
	if err != nil {
		return nil, err
	}

	return pr.getCommittedFiles(issues)
}

func (pr *prRepo) setQuery(prq PullRequestQuery) string {
	q := []string{
		"type:pr",
	}

	if prq.Author != "" {
		q = append(q, fmt.Sprintf("author:%s", prq.Author))

	}

	if prq.State != "" {
		q = append(q, fmt.Sprintf("state:%s", prq.State))
	}

	if prq.Org != "" {
		q = append(q, fmt.Sprintf("org:%s", prq.Org))
	}

	if prq.StartDate != "" {
		q = append(q, fmt.Sprintf("created:>=%s", prq.StartDate))
	}

	return strings.Join(q, " ")
}

func (pr *prRepo) getOwnerRepoFromPRUrl(s string) (ownerRepo, error) {
	or := ownerRepo{}
	u, err := url.Parse(s)
	if err != nil {
		return or, err
	}

	path := strings.Trim(u.Path, "/")
	paths := strings.Split(path, "/")

	if len(paths) < 3 {
		return or, fmt.Errorf("invalid URL: %s", s)
	}

	or.owner = paths[1]
	or.repo = paths[2]

	return or, nil
}

func (pr *prRepo) getIssues(query string) ([]*github.Issue, error) {
	opt := &github.SearchOptions{
		ListOptions: github.ListOptions{PerPage: pr.perPage},
	}

	allIssues := make([]*github.Issue, 0)
	for {
		issues, resp, err := pr.gh.Search.Issues(pr.ctx, query, opt)
		if err != nil {
			return allIssues, err
		}

		allIssues = append(allIssues, issues.Issues...)
		if resp.NextPage == 0 {
			break
		}

		opt.Page = resp.NextPage
	}

	return allIssues, nil
}

func (pr *prRepo) getCommittedFiles(issues []*github.Issue) ([]PullRequests, error) {
	prs := make([]PullRequests, 0)
	for _, i := range issues {
		ipr := PullRequests{
			Issue: i,
			Stats: Stats{
				NumFiles:  0,
				Additions: 0,
				Deletions: 0,
				Changes:   0,
				Total:     0,
			},
		}

		or, err := pr.getOwnerRepoFromPRUrl(i.GetURL())
		if err != nil {
			return prs, err
		}

		o := &github.ListOptions{PerPage: pr.perPage}
		fpr := make([]*github.CommitFile, 0)
		for {
			files, r, err := pr.gh.PullRequests.ListFiles(
				pr.ctx,
				or.owner,
				or.repo,
				i.GetNumber(),
				o,
			)
			if err != nil {
				return prs, err
			}

			fpr = append(fpr, files...)
			if r.NextPage == 0 {
				break
			}
			o.Page = r.NextPage
		}

		ipr.Files = fpr
		pr.setFileStat(&ipr.Stats, ipr.Files)

		prs = append(prs, ipr)
	}

	return prs, nil
}

func (pr *prRepo) setFileStat(stat *Stats, files []*github.CommitFile) {
	stat.NumFiles = len(files)

	for _, f := range files {
		stat.Additions += *f.Additions
		stat.Deletions += *f.Deletions
		stat.Changes += *f.Changes
		stat.Total += *f.Additions - *f.Deletions
	}
}

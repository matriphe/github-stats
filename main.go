package main

import (
	"os"

	"github.com/matriphe/github-stats/pkg/client"

	"github.com/jedib0t/go-pretty/v6/table"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	"github.com/matriphe/github-stats/pkg/calculator"
	"github.com/matriphe/github-stats/pkg/command"
	"github.com/matriphe/github-stats/pkg/output"
	"github.com/matriphe/github-stats/pkg/repository"
)

const (
	perPage = 30
)

func main() {
	app := &cli.App{
		Name:  "github-stats",
		Usage: "Simple Github stats",
		Commands: []*cli.Command{
			{
				Name:  "pr",
				Usage: "Get statistics of Pull Request created by user",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "token",
						Required: true,
						Usage:    "Github personal access token",
					},
					&cli.StringFlag{
						Name:  "org",
						Usage: "Get PR only for specific organization",
					},
					&cli.StringFlag{
						Name:  "state",
						Usage: "Get PR with specific state 'open' or 'closed'",
					},
					&cli.StringFlag{
						Name:  "start",
						Usage: "Get created PR started from this date, format YYYY-MM-DD",
					},
				},
				Action: func(c *cli.Context) error {
					ctx := c.Context
					token := c.String("token")

					ghAuthClient := client.NewGitHubAuthClient(ctx, token)

					userRepo := repository.NewUserRepo(ctx, ghAuthClient)
					prRepo := repository.NewPullRequestRepo(ctx, ghAuthClient, perPage)
					query := repository.PullRequestQuery{
						Org:       c.String("org"),
						State:     c.String("state"),
						StartDate: c.String("start"),
					}

					calc := calculator.NewPullRequestCalculator()

					result, err := command.PullRequest(userRepo, prRepo, query, calc)
					if err != nil {
						return err
					}

					w := output.NewPullRequestOutput(table.NewWriter(), result)
					w.ShowTitle(c.App.Usage)
					w.ShowPullRequests()

					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.WithError(err).Fatal("failed to run command")
	}
}

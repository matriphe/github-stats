package output

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/matriphe/github-stats/pkg/command"
)

type (
	prOutput struct {
		t      table.Writer
		title  string
		result command.PullRequestCommandResult
	}

	PullRequestOutput interface {
		ShowTitle()
		ShowPullRequests()
	}
)

func NewPullRequestOutput(
	t table.Writer,
	title string,
	r command.PullRequestCommandResult,
) PullRequestOutput {
	return &prOutput{
		t:      t,
		title:  title,
		result: r,
	}
}

func (s *prOutput) ShowTitle() {
	t := table.NewWriter()
	t.SetTitle(s.title)
	t.AppendRow(table.Row{"User", s.result.User.GetLogin()})

	if s.result.Query.Org != "" {
		t.AppendRow(table.Row{"Organization", s.result.Query.Org})
	}
	if s.result.Query.State != "" {
		t.AppendRow(table.Row{"State", s.result.Query.State})
	}
	if s.result.Query.StartDate != "" {
		t.AppendRow(table.Row{"From Date", s.result.Query.StartDate})
	}

	fmt.Println(t.Render())
}

func (s *prOutput) ShowPullRequests() {
	t := table.NewWriter()
	t.Style().Options.SeparateRows = true
	t.AppendHeader(table.Row{"#", "Pull Request", "Files", "Additions", "Deletions", "Changes", "Total"})
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, Align: text.AlignRight},
		{Number: 2, Align: text.AlignLeft},
		{Number: 3, Align: text.AlignRight},
		{Number: 4, Align: text.AlignRight},
		{Number: 5, Align: text.AlignRight},
		{Number: 6, Align: text.AlignRight},
		{Number: 7, Align: text.AlignRight},
	})
	i := 1
	for _, pr := range s.result.PullRequests {
		t.AppendRow(table.Row{
			fmt.Sprintf("%d", i),
			text.WrapSoft(pr.Issue.GetTitle(), 50),
			fmt.Sprintf("%d", pr.Stats.NumFiles),
			fmt.Sprintf("%d", pr.Stats.Additions),
			fmt.Sprintf("%d", pr.Stats.Deletions),
			fmt.Sprintf("%d", pr.Stats.Changes),
			fmt.Sprintf("%d", pr.Stats.Total),
		})
		i++
	}
	t.AppendFooter(table.Row{
		"",
		"Average",
		fmt.Sprintf("%d", s.result.Statistics.Avg.Files),
		fmt.Sprintf("%d", s.result.Statistics.Avg.Additions),
		fmt.Sprintf("%d", s.result.Statistics.Avg.Deletions),
		fmt.Sprintf("%d", s.result.Statistics.Avg.Changes),
		fmt.Sprintf("%d", s.result.Statistics.Avg.Total),
	})

	fmt.Println(t.Render())
}

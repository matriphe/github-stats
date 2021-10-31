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
		result command.PullRequestCommandResult
	}

	PullRequestOutput interface {
		ShowTitle(title string)
		ShowPullRequests()
	}
)

func NewPullRequestOutput(
	t table.Writer,
	r command.PullRequestCommandResult,
) PullRequestOutput {
	return &prOutput{
		t:      t,
		result: r,
	}
}

func (s *prOutput) ShowTitle(title string) {
	s.t.SetTitle(title)
	s.t.AppendRow(table.Row{"User", s.result.User.GetLogin()})

	if s.result.Query.Org != "" {
		s.t.AppendRow(table.Row{"Organization", s.result.Query.Org})
	}
	if s.result.Query.State != "" {
		s.t.AppendRow(table.Row{"State", s.result.Query.State})
	}
	if s.result.Query.StartDate != "" {
		s.t.AppendRow(table.Row{"From Date", s.result.Query.StartDate})
	}

	fmt.Println(s.t.Render())

	s.resetRender()
}

func (s *prOutput) ShowPullRequests() {
	s.t.Style().Options.SeparateRows = true
	s.t.AppendHeader(table.Row{"#", "Pull Request", "Files", "Additions", "Deletions", "Changes", "Total"})
	s.t.SetColumnConfigs([]table.ColumnConfig{
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
		s.t.AppendRow(table.Row{
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
	s.t.AppendFooter(table.Row{
		"",
		"Average",
		fmt.Sprintf("%d", s.result.Statistics.Avg.Files),
		fmt.Sprintf("%d", s.result.Statistics.Avg.Additions),
		fmt.Sprintf("%d", s.result.Statistics.Avg.Deletions),
		fmt.Sprintf("%d", s.result.Statistics.Avg.Changes),
		fmt.Sprintf("%d", s.result.Statistics.Avg.Total),
	})

	fmt.Println(s.t.Render())

	s.resetRender()
}

func (s *prOutput) resetRender() {
	s.t.ResetHeaders()
	s.t.ResetRows()
	s.t.ResetFooters()
}

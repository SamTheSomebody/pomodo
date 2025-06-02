package task

import (
	"pomodo/helpers"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

/*
	ID
  NAME - DUE
	SUMMARY
	Est: TIMEESTIMATE (TIMESPENT spent)
	Enth: #, Pri: #
*/

// TODO width options
type Model struct {
	task          *helpers.RawTask
	isInLineView  bool
	hideEnthusasm bool
	hidePriority  bool
	hideSummary   bool
	hideTimes     bool
	hideDueAt     bool
	showID        bool
}

func New(task *helpers.RawTask, opts ...Option) Model {
	return Model{
		task: task,
	}
}

type Option func(*Model)

func WithInLineView() Option {
	return func(m *Model) {
		m.isInLineView = true
	}
}

func WithoutEnthusiasm() Option {
	return func(m *Model) {
		m.hideEnthusasm = true
	}
}

func WithoutPriority() Option {
	return func(m *Model) {
		m.hidePriority = true
	}
}

func WithoutSummary() Option {
	return func(m *Model) {
		m.hideEnthusasm = true
	}
}

func WithoutDueAt() Option {
	return func(m *Model) {
		m.hideDueAt = true
	}
}

func WithoutTimes() Option {
	return func(m *Model) {
		m.hideTimes = true
	}
}

func WithID() Option {
	return func(m *Model) {
		m.showID = true
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return nil, nil
}

func (m Model) View() string {
	if m.isInLineView {
		return ""
	}
	return m.fullView()
}

func (m Model) fullView() string {
	b := strings.Builder{}
	if m.showID {
		b.WriteString(m.task.ID + "\n")
	}
	b.WriteString(m.task.Name + "\n")
	if !m.hideDueAt {
		b.WriteString(" - " + m.task.DueAt)
	}
	b.WriteString("\n")
	if !m.hideSummary {
		b.WriteString(m.task.Summary + "\n")
	}
	if !m.hideTimes {
		b.WriteString("Est: " + m.task.TimeEstimate + "(" + m.task.TimeSpent + " spent)\n")
	}
	// TODO add priority and enthusasm
	return b.String()
}

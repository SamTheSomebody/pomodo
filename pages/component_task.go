package pages

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
type taskModel struct {
	task          *helpers.RawTask
	isInLineView  bool
	hideEnthusasm bool
	hidePriority  bool
	hideSummary   bool
	hideTimes     bool
	hideDueAt     bool
	showID        bool
}

type taskOption func(*taskModel)

func WithInLineView() taskOption {
	return func(m *taskModel) {
		m.isInLineView = true
	}
}

func WithoutEnthusiasm() taskOption {
	return func(m *taskModel) {
		m.hideEnthusasm = true
	}
}

func WithoutPriority() taskOption {
	return func(m *taskModel) {
		m.hidePriority = true
	}
}

func WithoutSummary() taskOption {
	return func(m *taskModel) {
		m.hideEnthusasm = true
	}
}

func WithoutDueAt() taskOption {
	return func(m *taskModel) {
		m.hideDueAt = true
	}
}

func WithoutTimes() taskOption {
	return func(m *taskModel) {
		m.hideTimes = true
	}
}

func WithID() taskOption {
	return func(m *taskModel) {
		m.showID = true
	}
}

func (m taskModel) Init() tea.Cmd {
	return nil
}

func (m taskModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return nil, nil
}

func (m taskModel) View() string {
	if m.isInLineView {
		return ""
	}
	return m.fullView()
}

func (m taskModel) fullView() string {
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

func InitialTaskModel(t *helpers.RawTask) tea.Model {
	return taskModel{
		task: t,
	}
}

package bubbletea

import tea "github.com/charmbracelet/bubbletea"

// Error Message
type ErrMsg struct {
	Err error
}

func ErrCmd(err error) tea.Cmd {
	return func() tea.Msg {
		return ErrMsg{Err: err}
	}
}

// New Page Message
type NewPageMsg struct {
	Constructor func() (tea.Model, tea.Cmd)
}

func NewPageCmd(constructor func() (tea.Model, tea.Cmd)) tea.Cmd {
	return func() tea.Msg {
		return NewPageMsg{Constructor: constructor}
	}
}

// Log Message for writing to the status bar
type LogMsg struct {
	Message string
}

func LogCmd(message string) tea.Cmd {
	return func() tea.Msg {
		return LogMsg{Message: message}
	}
}

// Message for enabling/disabling navigation
type ItemSelectMsg struct {
	IsSelected bool
}

func ItemSelectCmd(isSelected bool) tea.Cmd {
	return func() tea.Msg {
		return ItemSelectMsg{IsSelected: isSelected}
	}
}

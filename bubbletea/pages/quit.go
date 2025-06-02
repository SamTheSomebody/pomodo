package pages

import tea "github.com/charmbracelet/bubbletea"

type QuitPage struct{}

func (m QuitPage) Init() tea.Cmd {
	return tea.Quit
}

func (m QuitPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, tea.Quit
}

func (m QuitPage) View() string {
	return "Quitting!"
}

func NewQuitPage() tea.Model {
	return QuitPage{}
}

// TODO add a bunch of stats for the date
// time spent, tasks completed, etc

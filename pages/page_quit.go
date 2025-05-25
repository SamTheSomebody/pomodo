package pages

import tea "github.com/charmbracelet/bubbletea"

type quitModel struct {
	state *State
}

func (m quitModel) Init() tea.Cmd {
	return tea.Quit
}

func (m quitModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m.state.ProcessUniversalKeys(msg)
}

func (m quitModel) View() string {
	return m.state.View("Quitting!")
}

func InitialQuitModel(s *State) tea.Model {
	return quitModel{state: s}
}

// TODO add a bunch of stats for the date
// time spent, tasks completed, etc

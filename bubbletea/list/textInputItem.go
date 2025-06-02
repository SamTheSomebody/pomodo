package list

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type TextInputItem struct {
	Input textinput.Model
}

func (m TextInputItem) Init() tea.Cmd {
	return nil
}

func (m TextInputItem) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.Input, _ = m.Input.Update(msg)
	return m, nil
}

func (m TextInputItem) View() string {
	return m.Input.View()
}

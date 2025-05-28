package pages

import (
	tea "github.com/charmbracelet/bubbletea"
)

type buttonModel struct {
	Label   string
	Focused bool
	OnClick func() (tea.Model, tea.Cmd)
}

func InitialButtonModel(label string, onClick func() (tea.Model, tea.Cmd)) buttonModel {
	return buttonModel{
		Label:   label,
		Focused: false,
		OnClick: onClick,
	}
}

func (m buttonModel) Init() tea.Cmd {
	return nil
}

func (m buttonModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if !m.Focused {
			return m, nil
		}
		switch msg.Type {
		case tea.KeyEnter:
			return m.OnClick()
		}
	}
	return m, nil
}

func (m buttonModel) View() string {
	if m.Focused {
		return activeButtonStyle.Render(m.Label)
	}
	return buttonStyle.Render(m.Label)
}

func (m *buttonModel) SetFocused(f bool) {
	m.Focused = f
}

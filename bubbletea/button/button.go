package button

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Label         string
	Focused       bool
	OnClick       func() (tea.Model, tea.Cmd)
	ActiveStyle   lipgloss.Style
	InactiveStyle lipgloss.Style
}

func New(label string, onClick func() (tea.Model, tea.Cmd)) Model {
	return Model{
		Label:   label,
		Focused: false,
		OnClick: onClick,
		ActiveStyle: lipgloss.NewStyle().Bold(true).Italic(true).
			Background(lipgloss.Color("#FF5F87")),
		InactiveStyle: lipgloss.NewStyle(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m Model) View() string {
	if m.Focused {
		return m.ActiveStyle.Render(m.Label)
	}
	return m.InactiveStyle.Render(m.Label)
}

func (m *Model) SetFocused(f bool) {
	m.Focused = f
}

func (m *Model) Focus() {
	m.Focused = true
}

func (m *Model) Blur() {
	m.Focused = false
}

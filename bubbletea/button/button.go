package button

import (
	tea "github.com/charmbracelet/bubbletea"

	"pomodo/bubbletea"
	"pomodo/bubbletea/list"
)

type Model struct {
	Label   string
	OnClick func() (tea.Model, tea.Cmd)
}

func New(label string, onClick func() (tea.Model, tea.Cmd)) Model {
	return Model{
		Label:   label,
		OnClick: onClick,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			return m, bubbletea.NewPageCmd(m.OnClick)
		}
	}
	return m, nil
}

func (m Model) View() string {
	return m.Label
}

func (m Model) OnSelect() (list.Item, tea.Cmd) {
	return m, bubbletea.NewPageCmd(m.OnClick)
}

func (m Model) OnSubmit() (list.Item, tea.Cmd) {
	return m, bubbletea.NewPageCmd(m.OnClick)
}

func (m Model) OnCancel() (list.Item, tea.Cmd) {
	return m, nil
}

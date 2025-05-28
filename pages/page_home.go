package pages

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type homeModel struct {
	state   *State
	focus   int
	buttons []buttonModel
	choice  string
}

func InitialHomeModel(s *State) homeModel {
	m := homeModel{
		buttons: []buttonModel{
			InitialButtonModel("Start Timer", OnConfigureTimerModelButtonClick(s, nil)),
			InitialButtonModel("View Tasks", OnViewTasksButtonClick(s)),
			InitialButtonModel("Add Task", OnEditTaskButtonClick(s, nil)),
		},
	}
	m.buttons[0].SetFocused(true)
	m.state = s
	s.Navigation.Add(m)
	return m
}

func (m homeModel) Init() tea.Cmd {
	m.focus = 0
	for i := range m.buttons {
		m.buttons[i].SetFocused(false)
	}
	m.buttons[m.focus].SetFocused(true)
	return nil
}

func (m homeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.state.Message = msg
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			b := m.buttons[m.focus]
			return b.OnClick()
		case "tab", "down", "j":
			return m.setFocus(1)
		case "shift+tab", "up", "k":
			return m.setFocus(-1)
		case "b", "esc", "q":
			return InitialQuitModel(m.state), nil
		case "ctrl+c":
			return nil, tea.Quit
		}
	}

	var cmd tea.Cmd
	return m, cmd
}

func (m homeModel) View() string {
	b := strings.Builder{}
	for _, button := range m.buttons {
		b.WriteString(button.View() + "\n")
	}
	return m.state.View(b.String())
}

func (m homeModel) setFocus(increment int) (tea.Model, tea.Cmd) {
	m.buttons[m.focus].SetFocused(false)
	m.focus += increment + len(m.buttons)
	m.focus %= len(m.buttons)
	m.buttons[m.focus].SetFocused(true)
	return m, nil
}

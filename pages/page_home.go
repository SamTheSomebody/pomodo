package pages

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"pomodo/internal/database"
)

type homeModel struct {
	state   *State
	focus   int
	buttons []buttonModel
	choice  string
}

func (m homeModel) Init() tea.Cmd {
	return nil
}

func (m homeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.state.Log = fmt.Sprintf("Msg: %v", msg)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			b := m.buttons[m.focus]
			return b.ClickModel, b.ClickCommand
		case "tab", "down", "j":
			return m.setFocus(1)
		case "shift+tab", "up", "k":
			return m.setFocus(-1)
		case "q":
			return InitialQuitModel(m.state), nil
		}
	}

	if mod, cmd := m.state.ProcessUniversalKeys(msg); mod != nil || cmd != nil {
		return mod, cmd
	}

	var cmd tea.Cmd
	return m, cmd
}

func (m homeModel) setFocus(increment int) (tea.Model, tea.Cmd) {
	m.buttons[m.focus].SetFocused(false)
	m.focus += increment + len(m.buttons)
	m.focus %= len(m.buttons)
	m.buttons[m.focus].SetFocused(true)
	return m, nil
}

func (m homeModel) View() string {
	b := strings.Builder{}
	for _, button := range m.buttons {
		b.WriteString(button.View() + "\n")
	}
	return m.state.View(b.String())
}

func InitialHomeModel(s *State) homeModel {
	m := homeModel{
		buttons: []buttonModel{
			InitialButtonModel("Start Timer", InitialConfigureTimerModel(s, nil), nil),
			InitialButtonModel("View Tasks", InitialViewTasksModel(s), nil),
			InitialButtonModel("Add Task", InitialEditTaskModel(s, database.Task{}), nil),
		},
	}
	m.buttons[0].SetFocused(true)
	m.state = s
	s.Navigation.Add(m)
	return m
}

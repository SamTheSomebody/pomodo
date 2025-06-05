package pages

import (
	tea "github.com/charmbracelet/bubbletea"

	"pomodo/bubbletea"
	"pomodo/bubbletea/button"
	"pomodo/bubbletea/list"
)

type HomePage struct {
	List list.Model
}

func NewHomePage(keymap *bubbletea.Keymap) HomePage {
	m := HomePage{
		List: list.New([]list.Item{
			button.New("Allocate Time", func() (tea.Model, tea.Cmd) { return NewAllocateTimePage(keymap), nil }),
			button.New("Start Timer", func() (tea.Model, tea.Cmd) { return NewConfigureTimerPage(nil, keymap), nil }),
			button.New("View Tasks", func() (tea.Model, tea.Cmd) { return NewViewTasksPage(keymap), nil }),
			button.New("Add Task", func() (tea.Model, tea.Cmd) { return NewEditTaskPage(nil, keymap), nil }),
		}, keymap),
	}
	return m
}

func (m HomePage) Init() tea.Cmd {
	return m.List.Init()
}

func (m HomePage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	l, cmd := m.List.Update(msg)
	m.List = l.(list.Model)
	return m, cmd
}

func (m HomePage) View() string {
	return m.List.View()
}

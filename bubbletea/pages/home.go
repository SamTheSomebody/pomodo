package pages

import (
	"pomodo/bubbletea/button"
	"pomodo/bubbletea/list"

	tea "github.com/charmbracelet/bubbletea"
)

type HomePage struct {
	List list.Model
}

func NewHomePage() HomePage {
	m := HomePage{
		List: list.New([]list.Item{
			button.New("Allocate Time", func() (tea.Model, tea.Cmd) { return NewAllocateTimePage(), nil }),
			button.New("Start Timer", func() (tea.Model, tea.Cmd) { return NewConfigureTimerPage(nil), nil }),
			button.New("View Tasks", func() (tea.Model, tea.Cmd) { return NewViewTasksPage(), nil }),
			button.New("Add Task", func() (tea.Model, tea.Cmd) { return NewEditTaskPage(nil), nil }),
		}),
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

package list

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Item fits the abstraction for text input but not for button
// Item needs focus, blur and select (and cancel?)

type Item interface {
	Init() tea.Cmd
	Update(msg tea.Msg) (tea.Model, tea.Cmd)
	View() string
	OnSelect() (Item, tea.Cmd)
	OnCancel() (Item, tea.Cmd)
	OnSubmit() (Item, tea.Cmd)
}

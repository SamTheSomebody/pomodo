package list

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// TODO create a time input field that validates time

type TextInputItem struct {
	Input    textinput.Model
	OldValue string
}

func NewTextInput(input textinput.Model) TextInputItem {
	return TextInputItem{
		Input:    input,
		OldValue: input.Value(),
	}
}

func (m TextInputItem) Init() tea.Cmd {
	return textinput.Blink
}

func (m TextInputItem) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.Input, _ = m.Input.Update(msg)
	return m, nil
}

func (m TextInputItem) View() string {
	return m.Input.View()
}

func (m TextInputItem) OnSelect() (Item, tea.Cmd) {
	m.OldValue = m.Input.Value()
	m.Input.Focus()
	return m, m.Init()
}

func (m TextInputItem) OnSubmit() (Item, tea.Cmd) {
	// TODO validate data
	m.Input.Blur()
	return m, nil
}

func (m TextInputItem) OnCancel() (Item, tea.Cmd) {
	m.Input.SetValue(m.OldValue)
	m.Input.Blur()
	return m, nil
}

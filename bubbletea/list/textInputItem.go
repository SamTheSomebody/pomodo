package list

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"pomodo/bubbletea"
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
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case bubbletea.ItemSelectMsg:
		if msg.IsSelected {
			m.Input.Focus()
			return m, textinput.Blink
		} else {
			m.Input.Blur()
			return m, cmd
		}
	}
	m.Input, cmd = m.Input.Update(msg)
	return m, cmd
}

func (m TextInputItem) View() string {
	return m.Input.View()
}

func (m TextInputItem) OnSelect() (Item, tea.Cmd) {
	m.OldValue = m.Input.Value()
	return m, bubbletea.ItemSelectCmd(true)
}

func (m TextInputItem) OnSubmit() (Item, tea.Cmd) {
	if m.Input.Err != nil {
		return m, bubbletea.ErrCmd(m.Input.Err)
	}
	return m, bubbletea.ItemSelectCmd(false)
}

func (m TextInputItem) OnCancel() (Item, tea.Cmd) {
	m.Input.SetValue(m.OldValue)
	return m, bubbletea.ItemSelectCmd(false)
}

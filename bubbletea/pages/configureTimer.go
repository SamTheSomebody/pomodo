package pages

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

/* Visual
(Error)
How long is the timer? : (ID)
  XXh XXm XXs
*/

type configureTimerModel struct {
	input textinput.Model
	err   error
}

func InitialConfigureTimerModel() configureTimerModel {
	m := configureTimerModel{}
	m.input = textinput.New()
	m.input.Placeholder = "XXh XXm XXs"
	m.input.SetValue("")
	m.input.Prompt = ""
	m.input.Width = 50
	m.input.Focus()
	return m
}

func (m configureTimerModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m configureTimerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			return InitialHomeModel(), nil
		case tea.KeyEnter:
			d, err := time.ParseDuration(m.input.Value())
			if err != nil {
				m.err = err
				return m, nil
			}
			return InitialTimerModel(d), nil
		}
	}
	m.input, _ = m.input.Update(msg)
	return m, nil
}

func (m configureTimerModel) View() string {
	s := "How long is the timer?\n"
	s += fmt.Sprintf("  %s\n", m.input.View())
	if m.err != nil {
		s += fmt.Sprintf("Error! %v", m.err)
	}
	s += "\n"
	return s
}

package pages

import (
	"fmt"
	"pomodo/bubbletea"
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
	nav *bubbletea.Navigation
	input textinput.Model
	err   error
}

func InitialConfigureTimerModel(nav *bubbletea.Navigation) configureTimerModel {
	m := configureTimerModel{}
	m.input = textinput.New()
	m.input.Placeholder = "XXh XXm XXs"
	m.input.SetValue("")
	m.input.Prompt = ""
	m.input.Width = 50
	m.input.Focus()
	m.nav = nav
	nav.Add(m)
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
			return m.nav.Back()
		case tea.KeyEnter:
			d, err := time.ParseDuration(m.input.Value())
			if err != nil {
				m.err = err
				return m, nil
			}
			return InitialTimerModel(m.nav, d), nil
		}
	}
	m.input, _ = m.input.Update(msg)
	return m, nil
}

func (m configureTimerModel) View() string {
	s := "\n" + padding + "How long is the timer?" + "\n"
	s += fmt.Sprint(padding, m.input.View(), "\n")
	if m.err != nil {
		s += fmt.Sprint(padding, "Error! ", m.err)
	}
	s += "\n"
	return s
}

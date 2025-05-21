package pages

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type configureTimerModel struct {
	state *State
	input textinput.Model
	err   error
}

func InitialConfigureTimerModel(s *State) configureTimerModel {
	m := configureTimerModel{}
	m.input = textinput.New()
	m.input.Placeholder = "XXh XXm XXs"
	m.input.SetValue("")
	m.input.Prompt = ""
	m.input.Width = 50
	m.input.Focus()
	m.state = s
	s.Navigation.Add(m)
	return m
}

func (m configureTimerModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m configureTimerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if mod, cmd := m.state.ProcessUniversalKeys(msg); mod != nil || cmd != nil {
		return mod, cmd
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			return m.state.Navigation.Back()
		case tea.KeyEnter:
			d, err := time.ParseDuration(m.input.Value())
			if err != nil {
				m.err = err
				return m, nil
			}
			return InitialTimerModel(m.state, d), nil
		}
	}
	m.input, _ = m.input.Update(msg)
	return m, nil
}

func (m configureTimerModel) View() string {
	s := header + padding + "How long is the timer?" + "\n"
	s += fmt.Sprint(padding, m.input.View(), "\n")
	if m.err != nil {
		s += fmt.Sprint(padding, "Error! ", m.err, "\n")
	}
	m.state.HelpView()
	return s
}

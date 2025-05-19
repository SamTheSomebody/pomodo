package bubbletea

import tea "github.com/charmbracelet/bubbletea"

type State struct {
	History []tea.Model
}

func (s *State) Back() (tea.Model, tea.Cmd) {
	l := len(s.History) - 1
	if l < 0 {
		return nil, tea.Quit
	}
	m := s.History[l]
	s.History = s.History[:l]
	return m, nil
}

func (s *State) Push(m tea.Model, c tea.Cmd) (tea.Model, tea.Cmd) {
	s.History = append(s.History, m)
	return m, c
}

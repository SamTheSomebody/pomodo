package bubbletea

import tea "github.com/charmbracelet/bubbletea"

type Navigation struct {
	History []tea.Model
}

func NewNavigation() *Navigation {
	return &Navigation{
		History: make([]tea.Model, 0),
	}
}

func (s *Navigation) Back() (tea.Model, tea.Cmd) {
	l := len(s.History) - 1
	if l < 0 {
		return nil, tea.Quit
	}
	m := s.History[l]
	s.History = s.History[:l]
	return m, nil
}

func (s *Navigation) Add(m tea.Model) { //If model is home, clear history?
	s.History = append(s.History, m)
}

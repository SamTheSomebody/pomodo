package pages

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type Navigation struct {
	State   *State
	History []*tea.Model
}

func NewNavigation(state *State) *Navigation {
	return &Navigation{
		State:   state,
		History: make([]*tea.Model, 0),
	}
}

func (s *Navigation) Back() (tea.Model, tea.Cmd) {
	l := len(s.History) - 2
	if l < 0 {
		return nil, tea.Quit
	}
	m := s.History[l]
	s.History = s.History[:l]
	s.State.Log = fmt.Sprintf("Back to %T in navigation history [%v]", *m, len(s.History))
	return *m, nil
}

func (s *Navigation) Add(m tea.Model) {
	s.State.Log = fmt.Sprintf("Added %T to navigation history [%v]", m, len(s.History))
	if _, ok := m.(homeModel); ok {
		s.History = s.History[:0]
	}
	s.History = append(s.History, &m)
}

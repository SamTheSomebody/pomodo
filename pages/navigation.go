package pages

import (
	"fmt"
	"reflect"

	tea "github.com/charmbracelet/bubbletea"
)

type Navigation struct {
	State   *State
	History []tea.Model
}

func NewNavigation(state *State) Navigation {
	return Navigation{
		State:   state,
		History: make([]tea.Model, 0),
	}
}

func (n *Navigation) Back() (tea.Model, tea.Cmd) {
	l := len(n.History) - 2
	if l < 0 {
		return InitialQuitModel(n.State), nil
	}
	m := n.History[l]
	m.Init()
	n.History = n.History[:l+1]
	n.State.Log = fmt.Sprintf("Back to %T in navigation history [%v of %v],  %p", m, l+1, len(n.History), &m)
	return m, nil
}

// TODO this isn't being called?
func (n *Navigation) Add(m tea.Model) {
	n.State.Err = nil
	if _, ok := m.(homeModel); ok {
		n.History = n.History[:0]
	} else if reflect.TypeOf(m) == reflect.TypeOf(n.Last()) {
		return
	}
	n.History = append(n.History, m)
	n.State.Log = fmt.Sprintf("Added %T to navigation history [%v], %p", m, len(n.History), &m)
}

func (n *Navigation) Last() tea.Model {
	if len(n.History) == 0 {
		return nil
	}
	return n.History[len(n.History)-1]
}

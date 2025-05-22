package pages

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type State struct {
	Navigation *Navigation
	Help       *help.Model
	Keys       *keymap
}

func NewState() *State {
	h := help.New()
	return &State{
		Navigation: NewNavigation(),
		Help:       &h,
		Keys:       NewKeymap(),
	}
}

func (s *State) HelpView(keys ...key.Binding) string {
	b := []key.Binding{
		s.Keys.Back,
		s.Keys.Kill,
		s.Keys.Enter,
	}
	b = append(b, keys...)
	return "\n\n" + padding + s.Help.ShortHelpView(b)
}

func (s *State) ProcessUniversalKeys(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, s.Keys.Back):
			return s.Navigation.Back()
		case key.Matches(msg, s.Keys.Kill):
			mod := InitialQuitModel(s)
			s.Navigation.Add(mod)
			return mod, nil
		}
	}
	return nil, nil
}

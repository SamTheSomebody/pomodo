package pages

import (
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type State struct {
	Navigation *Navigation
	Help       *help.Model
	Keys       *keymap
	Err        error
	Log        string
}

func NewState() *State {
	h := help.New()
	s := &State{
		Help: &h,
		Keys: NewKeymap(),
	}
	s.Navigation = NewNavigation(s)
	return s
}

func (s *State) View(modelView string, keys ...key.Binding) string {
	t, _ := time.ParseDuration("7h43m")
	x := Header(s.Log, 10, t)
	x += regularStyle.Render(modelView)
	x += helpStyle.Render(s.helpView(keys...))
	x += s.footer()
	return x
}

func (s *State) footer() string {
	x := ""
	if s.Err != nil {
		x += "\n" + errorStyle.Render(s.Err.Error())
	}
	return x
}

func (s *State) helpView(keys ...key.Binding) string {
	b := []key.Binding{
		s.Keys.Back,
		s.Keys.Kill,
		s.Keys.Enter,
	}
	b = append(b, keys...)
	return "\n\n" + s.Help.ShortHelpView(b)
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

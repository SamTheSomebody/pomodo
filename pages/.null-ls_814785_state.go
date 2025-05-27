package pages

import (
	"strings"
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
	b := strings.Builder{}
	b.WriteString(Header(s.Log, 10, t))
	b.WriteString(regularStyle.Render(modelView))
	b.WriteString(helpStyle.Render(s.helpView(keys...)))
	b.WriteString(s.footer())
	return b.String()
}

func (s *State) footer() string {
	x := ""
	if s.Err != nil {
		x += errorStyle.Render(s.Err.Error())
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
	return "\n" + s.Help.ShortHelpView(b)
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

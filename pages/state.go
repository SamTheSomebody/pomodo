package pages

import (
	"context"
	"fmt"
	"pomodo/helpers"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type State struct {
	Navigation Navigation
	Help       help.Model
	Keys       keymap
	Err        error
	Log        string
	Message    tea.Msg
}

func NewState() *State {
	h := help.New()
	s := &State{
		Log:  "",
		Help: h,
		Keys: NewKeymap(),
	}
	s.Navigation = NewNavigation(s)
	return s
}

func (s *State) View(modelView string, keys ...key.Binding) string {
	t, err := time.ParseDuration("7h43m")
	if err != nil {
		s.Err = err
		return s.Err.Error()
	}
	db := helpers.GetDBQueries()
	tasks, err := db.GetTasks(context.TODO())
	if err != nil {
		s.Err = err
		return s.Err.Error()
	}
	taskCount := len(tasks)
	b := strings.Builder{}
	b.WriteString(Header(s.Log, s.Message, taskCount, t))
	b.WriteString(fmt.Sprintf("State: %p, Nav State: %p, Nav length: %v\n", s, s.Navigation.State, len(s.Navigation.History)))
	b.WriteString(regularStyle.Render(modelView))
	if keys != nil {
		b.WriteString(helpStyle.Render(s.helpView(keys...)))
	}
	b.WriteString(s.footer())
	return b.String()
}

func (s *State) footer() string {
	x := "\n"
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

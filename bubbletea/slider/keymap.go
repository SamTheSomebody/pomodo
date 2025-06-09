package slider

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
)

type Keymap struct {
	CursorUp   key.Binding
	CursorDown key.Binding
	NextPage   key.Binding
	PrevPage   key.Binding
	GoToStart  key.Binding
	GoToEnd    key.Binding

	help help.Model
}

func DefaultKeymap() Keymap {
	return Keymap{
		// Browsing.
		CursorUp: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "up"),
		),
		CursorDown: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "down"),
		),
		PrevPage: key.NewBinding(
			key.WithKeys("left", "h", "pgup", "b", "u"),
			key.WithHelp("←/h/pgup", "prev page"),
		),
		NextPage: key.NewBinding(
			key.WithKeys("right", "l", "pgdown", "f", "d"),
			key.WithHelp("→/l/pgdn", "next page"),
		),
		GoToStart: key.NewBinding(
			key.WithKeys("home", "g"),
			key.WithHelp("g/home", "go to start"),
		),
		GoToEnd: key.NewBinding(
			key.WithKeys("end", "G"),
			key.WithHelp("G/end", "go to end"),
		),
		help: help.New(),
	}
}

// Toggle navigation if item is focused
func (k *Keymap) SetNavigationEnabled(enabled bool) {
	k.CursorUp.SetEnabled(enabled)
	k.CursorDown.SetEnabled(enabled)
	k.NextPage.SetEnabled(enabled)
	k.PrevPage.SetEnabled(enabled)
	k.GoToStart.SetEnabled(enabled)
	k.GoToEnd.SetEnabled(enabled)
}

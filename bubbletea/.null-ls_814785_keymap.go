package bubbletea

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
	Return     key.Binding

	Select key.Binding
	Submit key.Binding
	Cancel key.Binding

	ForceQuit key.Binding

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
		Return: key.NewBinding(
			key.WithKeys("esc", "b"),
			key.WithHelp("b/esc", "return"),
		),

		// Selection
		Select: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "select"),
		),
		Submit: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "confirm"),
			key.WithDisabled(),
		),
		Cancel: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "cancel"),
			key.WithDisabled(),
		),

		ForceQuit: key.NewBinding(key.WithKeys("ctrl+c")),

		help: help.New(),
	}
}

// Toggle navigation if item is focused
func (k *Keymap) SetNavigationEnabled(enabled bool) {
	k.Submit.SetEnabled(!enabled)
	k.Cancel.SetEnabled(!enabled)

	k.Select.SetEnabled(enabled)
	k.Return.SetEnabled(enabled)

	k.CursorUp.SetEnabled(enabled)
	k.CursorDown.SetEnabled(enabled)
	k.NextPage.SetEnabled(enabled)
	k.PrevPage.SetEnabled(enabled)
	k.GoToStart.SetEnabled(enabled)
	k.GoToEnd.SetEnabled(enabled)
}

func VerticalNaviagtionKeymap() Keymap {
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
		Return: key.NewBinding(
			key.WithKeys("esc", "b"),
			key.WithHelp("b/esc", "return"),
		),

		ForceQuit: key.NewBinding(key.WithKeys("ctrl+c")),

		help: help.New(),
	}
}

// Toggle navigation if item is focused
func (k *Keymap) SetNavigationEnabled(enabled bool) {
	k.Submit.SetEnabled(!enabled)
	k.Cancel.SetEnabled(!enabled)

	k.Select.SetEnabled(enabled)
	k.Return.SetEnabled(enabled)

	k.CursorUp.SetEnabled(enabled)
	k.CursorDown.SetEnabled(enabled)
	k.NextPage.SetEnabled(enabled)
	k.PrevPage.SetEnabled(enabled)
	k.GoToStart.SetEnabled(enabled)
	k.GoToEnd.SetEnabled(enabled)
}

func (k *Keymap) SetHomeKeysEnabled(enabled bool) {
	if enabled {
		k.Return.SetKeys("esc", "q")
		k.Return.SetHelp("esc/q", "quit")
	} else {
		k.Return.SetKeys("esc")
		k.Return.SetHelp("esc", "return")
	}
}

func (k *Keymap) Help() string {
	keys := []key.Binding{
		k.CursorUp,
		k.CursorDown,
		k.NextPage,
		k.PrevPage,
		k.GoToStart,
		k.GoToEnd,
		k.Return,
		k.Select,
		k.Submit,
		k.Cancel,
		k.ForceQuit,
	}
	return k.help.ShortHelpView(keys)
}

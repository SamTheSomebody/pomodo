package list

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	CursorUp   key.Binding
	CursorDown key.Binding
	NextPage   key.Binding
	PrevPage   key.Binding
	GoToStart  key.Binding
	GoToEnd    key.Binding

	Select  key.Binding
	Confirm key.Binding
	Cancel  key.Binding

	Quit      key.Binding
	ForceQuit key.Binding
}

func DefaultKeyMap() KeyMap {
	return KeyMap{
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

		// Selection
		Select: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "select"),
		),
		Confirm: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "confirm"),
			key.WithDisabled(),
		),
		Cancel: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "cancel"),
			key.WithDisabled(),
		),

		// Quitting.
		Quit: key.NewBinding(
			key.WithKeys("q", "esc"),
			key.WithHelp("q", "quit"),
		),
		ForceQuit: key.NewBinding(key.WithKeys("ctrl+c")),
	}
}

// Disables navigation if item is focused
func (k *KeyMap) SetItemFocus(b bool) {
	k.Confirm.SetEnabled(b)
	k.Cancel.SetEnabled(b)
	k.Select.SetEnabled(!b)

	k.CursorUp.SetEnabled(!b)
	k.CursorDown.SetEnabled(!b)
	k.NextPage.SetEnabled(!b)
	k.PrevPage.SetEnabled(!b)
	k.GoToStart.SetEnabled(!b)
	k.GoToEnd.SetEnabled(!b)

	k.Quit.SetEnabled(!b)
}

package pages

import (
	"github.com/charmbracelet/bubbles/key"
)

// Do I want to have this be a dictionary?
type keymap struct {
	Kill  key.Binding
	Back  key.Binding
	Enter key.Binding
}

func NewKeymap() keymap {
	return keymap{
		Kill: key.NewBinding(
			key.WithKeys("ctrl+c"),
			key.WithHelp("ctrl+c", "kill"),
		),

		Back: key.NewBinding(
			key.WithKeys("b"),
			key.WithHelp("b", "back"),
		),
		Enter: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "continue"),
			key.WithDisabled(),
		),
	}
}

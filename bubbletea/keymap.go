package bubbletea

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
)

type Keymap struct {
	Return    key.Binding
	ForceQuit key.Binding

	help help.Model
}

func DefaultKeymap() Keymap {
	return Keymap{
		Return: key.NewBinding(
			key.WithKeys("esc", "b"),
			key.WithHelp("b/esc", "return"),
		),
		ForceQuit: key.NewBinding(key.WithKeys("ctrl+c")),
		help:      help.New(),
	}
}

func (k *Keymap) SetNavigationEnabled(b bool) {
	k.Return.SetEnabled(b)
}

func (k *Keymap) Help() string {
	keys := []key.Binding{
		k.ForceQuit,
	}
	return k.help.ShortHelpView(keys)
}

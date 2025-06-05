package list

import "github.com/charmbracelet/lipgloss"

type Styles struct {
	ActiveItem   lipgloss.Style
	InactiveItem lipgloss.Style
}

func (s Styles) Render(i Item, isFocused bool) string {
	if isFocused {
		return s.ActiveItem.Render(i.View())
	}
	return s.InactiveItem.Render(i.View())
}

func DefaultStyle() Styles {
	return Styles{
		ActiveItem: lipgloss.NewStyle().Bold(true).Italic(true).
			Background(lipgloss.Color("#A550DF")),
		InactiveItem: lipgloss.NewStyle().Padding(0, 2).Foreground(lipgloss.Color("#343433")),
	}
}

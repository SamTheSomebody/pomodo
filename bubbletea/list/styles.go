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
			PaddingLeft(2).Background(lipgloss.Color("#A550DF")),
		InactiveItem: lipgloss.NewStyle(),
	}
}

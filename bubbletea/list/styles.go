package list

import "github.com/charmbracelet/lipgloss"

type Styles struct {
	ActiveItem   lipgloss.Style
	InactiveItem lipgloss.Style
}

func (s Styles) Render(i Item) string {
	if i.IsFocused {
		return s.ActiveItem.Render(i.Model.View())
	}
	return s.InactiveItem.Render(i.Model.View())
}

func DefaultStyle() Styles {
	return Styles{
		ActiveItem:   lipgloss.NewStyle().Background(lipgloss.Color("#A550DF")),
		InactiveItem: lipgloss.NewStyle(),
	}
}

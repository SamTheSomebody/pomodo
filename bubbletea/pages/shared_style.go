package pages

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	padding  = 2
	maxWidth = 80
)

var (
	regularStyle = lipgloss.NewStyle().
			Padding(1, 2)

	errorStyle = lipgloss.NewStyle().
			Inherit(regularStyle).
			Foreground(lipgloss.AdaptiveColor{Light: "#343433", Dark: "#FF5F87"}).
			Background(lipgloss.AdaptiveColor{Light: "#FF5F87", Dark: "#353533"}).
			Italic(true).
			SetString("Error: ")

	helpStyle = lipgloss.NewStyle().
			Inherit(regularStyle).
			Foreground(lipgloss.Color("8"))

	searchStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("8")).
			Italic(true)

	// Button Styles

	buttonStyle = lipgloss.NewStyle().
			Inherit(regularStyle)

	activeButtonStyle = lipgloss.NewStyle().
				Inherit(buttonStyle).
				Bold(true).
				Italic(true).
				Background(lipgloss.Color("#FF5F87"))

	// Status Bar
	statusNugget = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Padding(0, 1)

	statusBarStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#343433", Dark: "#C1C6B2"}).
			Background(lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#353533"})

	statusStyle = lipgloss.NewStyle().
			Inherit(statusBarStyle).
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#FF5F87")).
			Padding(0, 1).
			MarginRight(1)

	tasksRemainingStyle = statusNugget.
				Background(lipgloss.Color("#A550DF")).
				Align(lipgloss.Right)

	statusText = lipgloss.NewStyle().Inherit(statusBarStyle)

	allocatedTimeStyle = statusNugget.Background(lipgloss.Color("#6124DF"))
)

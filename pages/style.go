package pages

import (
	"fmt"
	"os"
	"strconv"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
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

func Header(log string, message tea.Msg, taskCount int, allocatedTimeLength time.Duration) string {
	w := lipgloss.Width
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))
	statusKey := statusStyle.Render("POMODO")
	tasksRemaining := tasksRemainingStyle.Render(strconv.Itoa(taskCount) + " tasks")
	allocatedTime := allocatedTimeStyle.Render(allocatedTimeLength.String())
	statusMessage := statusText.Render(fmt.Sprintf(" %T: %v ", message, message))
	statusVal := statusText.
		Width(width - w(statusKey) - w(tasksRemaining) - w(allocatedTime) - w(statusMessage)).
		Render(log)

	bar := lipgloss.JoinHorizontal(lipgloss.Top,
		statusKey,
		statusVal,
		statusMessage,
		tasksRemaining,
		allocatedTime,
	)

	return statusBarStyle.Width(width).Render(bar) + "\n"
}

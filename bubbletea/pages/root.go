package pages

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"

	"pomodo/bubbletea"
	"pomodo/helpers"
)

type RootPage struct {
	Pages  []tea.Model
	KeyMap bubbletea.KeyMap
	Log    string
	Msg    tea.Msg
	Err    error
	Help   help.Model
}

func NewRootPage() RootPage {
	keymap := bubbletea.DefaultKeyMap()
	return RootPage{
		Pages:  []tea.Model{NewHomePage(&keymap)},
		KeyMap: keymap,
		Help:   help.New(),
	}
}

func (m RootPage) Init() tea.Cmd {
	return nil
}

func (m RootPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.Msg = msg
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, m.KeyMap.ForceQuit) {
			return m, tea.Quit
		}
		if key.Matches(msg, m.KeyMap.Return) {
			if len(m.Pages) < 2 {
				return NewQuitPage(), nil // TODO new page command
			}
			m.Pages = m.Pages[:len(m.Pages)-1]
			m.Pages[len(m.Pages)-1].Init()
			return m, nil
		}
	case bubbletea.NewPageMsg:
		m.Err = nil
		p, cmd := msg.Constructor()
		m.Log = fmt.Sprintf("Add page %T to pages [len: %v]", p, len(m.Pages))
		if reflect.TypeOf(p) == reflect.TypeOf(m.Pages[len(m.Pages)-1]) {
			m.Pages[len(m.Pages)-1] = p
			return m, cmd
		}
		switch p.(type) {
		case HomePage:
			m.Pages = m.Pages[:1]
			m.Pages[0] = p
			return m, cmd
		default:
			m.Pages = append(m.Pages, p)
			return m, cmd
		}
	case bubbletea.ErrMsg:
		m.Err = msg.Err
		return m, nil
	case bubbletea.LogMsg:
		m.Log = msg.Message
		return m, nil
	case bubbletea.EnableNavigationMsg:
		m.KeyMap.EnableNavigation(msg.Enabled)
		m.Log = fmt.Sprintf("Nagivtaion Enabled: %v", msg.Enabled)
		return m, nil
	}
	m.Pages[len(m.Pages)-1], cmd = m.Pages[len(m.Pages)-1].Update(msg)
	return m, cmd
}

func (m RootPage) View() string {
	b := strings.Builder{}
	// Header (Status Bar)
	t, err := time.ParseDuration("7h43m")
	if err != nil {
		return err.Error()
	}
	db := helpers.GetDBQueries()
	tasks, err := db.GetTasks(context.TODO())
	if err != nil {
		return err.Error()
	}
	w := lipgloss.Width
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))
	statusKey := statusStyle.Render("POMODO")
	tasksRemaining := tasksRemainingStyle.Render(fmt.Sprintf("%v tasks", len(tasks)))
	allocatedTime := allocatedTimeStyle.Render(t.String())
	statusMessage := statusText.Render(fmt.Sprintf(" %T: %v ", m.Msg, m.Msg))
	statusVal := statusText.
		Width(width - w(statusKey) - w(tasksRemaining) - w(allocatedTime) - w(statusMessage)).
		Render(m.Log)
	bar := lipgloss.JoinHorizontal(lipgloss.Top,
		statusKey,
		statusVal,
		statusMessage,
		tasksRemaining,
		allocatedTime,
	)
	b.WriteString(statusBarStyle.Width(width).Render(bar) + "\n")

	// Current Page
	b.WriteString("\n" + m.Pages[len(m.Pages)-1].View() + "\n")

	// Footer
	keys := []key.Binding{
		m.KeyMap.Return,
		m.KeyMap.ForceQuit,
		m.KeyMap.Select,
	}
	b.WriteString("\n" + m.Help.ShortHelpView(keys))
	if m.Err != nil {
		b.WriteString("\n" + errorStyle.Render(m.Err.Error()))
	}
	return b.String()
}

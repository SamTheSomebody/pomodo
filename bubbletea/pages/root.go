package pages

import (
	"context"
	"fmt"
	"log"
	"os"
	"pomodo/bubbletea"
	"pomodo/helpers"
	"reflect"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

type RootPage struct {
	Pages         []tea.Model
	Keymap        bubbletea.Keymap
	Log           string
	Msg           tea.Msg
	Err           error
	Help          help.Model
	AllocatedTime time.Duration
}

func NewRootPage() RootPage {
	var page tea.Model
	page = NewHomePage()
	users, err := helpers.GetDBQueries().GetUsers(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	if len(users) < 1 {
		page = NewAddUserPage()
	}
	allocatedTime, err := helpers.GetAllocatedTime()
	return RootPage{
		Pages:         []tea.Model{page},
		Keymap:        bubbletea.DefaultKeymap(),
		Help:          help.New(),
		AllocatedTime: allocatedTime,
		Err:           err,
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
		if key.Matches(msg, m.Keymap.ForceQuit) {
			return m, tea.Quit
		}
		if key.Matches(msg, m.Keymap.Return) {
			if len(m.Pages) < 2 {
				return m, bubbletea.NewPageCmd(func() (tea.Model, tea.Cmd) { return NewQuitPage(), nil })
			}
			m.Pages = m.Pages[:len(m.Pages)-1]
			m.Pages[len(m.Pages)-1].Init()
			return m, cmd
		}
	case bubbletea.NewPageMsg:
		m.Err = nil
		p, cmd := msg.Constructor()
		m.Log = fmt.Sprintf("Add page '%T' to pages [len: %v]", p, len(m.Pages))
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
		return m, cmd
	case bubbletea.LogMsg:
		m.Log = msg.Message
		return m, cmd
	case bubbletea.ItemSelectMsg:
		m.Keymap.SetNavigationEnabled(!msg.IsSelected)
		// Don't return, the children still need this
	}
	m.Pages[len(m.Pages)-1], cmd = m.Pages[len(m.Pages)-1].Update(msg)
	return m, cmd
}

func (m RootPage) View() string {
	b := strings.Builder{}
	// Header (Status Bar)
	db := helpers.GetDBQueries()
	tasks, err := db.GetTasks(context.TODO())
	if err != nil {
		return err.Error()
	}
	w := lipgloss.Width
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))
	statusKey := statusStyle.Render("POMODO")
	tasksRemaining := tasksRemainingStyle.Render(fmt.Sprintf("%v tasks", len(tasks)))
	allocatedTime := allocatedTimeStyle.Render(m.AllocatedTime.String())
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
	b.WriteString("\n" + m.Keymap.Help())
	if m.Err != nil {
		b.WriteString("\n" + errorStyle.Render(m.Err.Error()))
	}
	return b.String()
}

package pages

import (
	"fmt"
	"io"
	"pomodo/internal/database"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Background(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type item string

func (i item) FilterValue() string {
	return ""
}

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

type homeModel struct {
	state  *State
	list   list.Model
	choice string
}

func (m homeModel) Init() tea.Cmd {
	return nil
}

func (m homeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = string(i)
				switch m.choice {
				case "Start Timer":
					return InitialConfigureTimerModel(m.state, nil), nil
				case "View Tasks":
					return InitialViewTasksModel(m.state), nil
				case "Add Task":
					return InitialEditTaskModel(m.state, database.Task{}), nil
				}
			}
			return m, tea.Quit
		}
	}

	if mod, cmd := m.state.ProcessUniversalKeys(msg); mod != nil || cmd != nil {
		return mod, cmd
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m homeModel) View() string {
	s := m.list.View()
	k := []key.Binding{
		m.list.KeyMap.CursorUp,
		m.list.KeyMap.CursorDown,
		m.list.KeyMap.Filter,
		m.list.KeyMap.AcceptWhileFiltering,
		m.list.KeyMap.Quit,
		m.list.KeyMap.ForceQuit,
	}
	return m.state.View(s, k...)
}

func InitialHomeModel(s *State) homeModel {
	items := []list.Item{
		item("Start Timer"),
		item("View Tasks"),
		item("Add Task"),
	}

	l := list.New(items, itemDelegate{}, 20, 9)
	l.Title = "Hi! Welcome to pomodo, what did you want to do?"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle

	m := homeModel{list: l}
	m.state = s
	s.Navigation.Add(m)
	return m
}

package taskselect

import (
	"context"
	"fmt"
	"log"
	"pomodo/bubbletea/list"
	"pomodo/helpers"
	"pomodo/internal/database"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

type Model struct {
	Options  []database.Task
	Filtered []database.Task
	Selected *database.Task
	Search   string
	Index    int
	Focused  bool
}

func New() Model {
	db := helpers.GetDBQueries()
	tasks, err := db.GetTasks(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	return Model{
		Options:  tasks,
		Filtered: make([]database.Task, 0),
		Selected: nil,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if len(m.Filtered) == 0 {
				m.Selected = nil
			} else {
				m.Selected = &m.Filtered[m.Index]
			}
			m.SetFocused(!m.Focused)
			return m, nil
		case "shift+tab", "up":
			if len(m.Filtered) > 0 {
				m.Index += len(m.Filtered) - 1
				m.Index %= len(m.Filtered)
			}
		case "tab", "down":
			if len(m.Filtered) > 0 {
				m.Index += len(m.Filtered) + 1
				m.Index %= len(m.Filtered)
			}
		case "backspace":
			if len(m.Search) > 0 {
				m.Search = m.Search[:len(m.Search)-1]
				m.filterOptions()
			}
		case "esc":
			m.SetFocused(false)
			return m, nil
		default: // TODO this should check if it is a valid input key
			if msg.Type == tea.KeyRunes {
				m.Search += string(msg.Runes)
				m.filterOptions()
			}
		}
	}
	return m, nil
}

func (m *Model) filterOptions() {
	m.Index = 0
	search := strings.ToLower(m.Search)
	m.Filtered = m.Filtered[:0]
	if len(m.Search) == 0 {
		m.Filtered = m.Options
		return
	}
	for _, option := range m.Options {
		if strings.Contains(strings.ToLower(option.Name), search) {
			m.Filtered = append(m.Filtered, option)
		}
	}
}

func (m Model) View() string {
	b := strings.Builder{}
	b.WriteString("Select Task: [" + m.GetSelected() + " ]")
	if !m.Focused {
		return b.String()
	}
	s := fmt.Sprintf(" (%v of %v)", len(m.Filtered), len(m.Options))
	b.WriteString("\nSearch: " + m.Search + s)
	for i, option := range m.Filtered {
		if i == m.Index {
			b.WriteString("\n" + option.Name)
		} else {
			b.WriteString("\n" + (option.Name))
		}
	}
	return b.String()
}

func (m *Model) SetFocused(f bool) {
	m.Focused = f
	m.Search = ""
	m.Filtered = m.Options
}

func (m *Model) GetSelected() string {
	if m.Selected == nil {
		return "None"
	}
	return m.Selected.Name
}

func (m *Model) GetTaskID() *uuid.UUID {
	if m.Selected == nil {
		return nil
	}
	switch m.Selected.ID.(type) {
	case uuid.UUID:
		id := m.Selected.ID.(uuid.UUID)
		return &id
	case string:
		id, err := uuid.Parse(m.Selected.ID.(string))
		if err != nil {
			log.Fatal(err)
		}
		return &id
	}
	return nil
}

func (m Model) OnSelect() (list.Item, tea.Cmd) {
	m.SetFocused(true)
	return m, nil
}

func (m Model) OnSubmit() (list.Item, tea.Cmd) {
	m.SetFocused(false)
	return m, nil
}

func (m Model) OnCancel() (list.Item, tea.Cmd) {
	m.SetFocused(false)
	return m, nil
}

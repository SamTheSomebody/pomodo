package pages

import (
	"context"
	"log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/google/uuid"

	"pomodo/helpers"
	"pomodo/internal/database"
)

type selectTaskModel struct {
	Options    []database.Task
	Filtered   []database.Task
	Selected   *database.Task
	Search     string
	Index      int
	Open       bool
	Focused    bool
	ShowCursor bool
}

func InitialSelectTaskModel() selectTaskModel {
	db := helpers.GetDBQueries()
	tasks, err := db.GetTasks(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	return selectTaskModel{
		Options:  tasks,
		Filtered: tasks,
		Selected: nil,
	}
}

func (m selectTaskModel) Init() tea.Cmd {
	return nil
}

func (m selectTaskModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if !m.Focused {
		return m, nil
	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if m.Open {
				if len(m.Filtered) == 0 {
					m.Selected = nil
				}
				m.Selected = &m.Filtered[m.Index]
				m.Open = false
				m.Search = ""
			} else {
				m.Open = true
			}
			return m, nil
		case "shift+tab", "up":
			if m.Open && m.Index > 0 {
				m.Index--
			}
		case "tab", "down":
			if m.Open && m.Index < len(m.Filtered)-1 {
				m.Index++
			}
		case "backspace":
			if len(m.Search) > 0 {
				m.Search = m.Search[:len(m.Search)-1]
				m.filterOptions()
			}
		default: // TODO this should check if it is a valid input key
			if m.Open && msg.Type == tea.KeyRunes {
				m.Search += string(msg.Runes)
				m.filterOptions()
			}
		}
	}
	return m, nil
}

func (m *selectTaskModel) filterOptions() {
	search := strings.ToLower(m.Search)
	m.Filtered = nil
	for _, opt := range m.Options {
		if strings.Contains(strings.ToLower(opt.Name), search) {
			m.Filtered = append(m.Filtered, opt)
		}
	}
	m.Index = 0
	if len(m.Filtered) == 0 {
		m.Filtered = []database.Task{{Name: "<no match>"}}
	}
}

func (m selectTaskModel) View() string {
	if len(m.Options) == 0 {
		return ""
	}
	b := strings.Builder{}
	b.WriteString("Select Task: [" + m.GetSelected() + " ]")
	if !m.Focused {
		return b.String()
	}
	s := lipgloss.NewStyle().Background(lipgloss.Color("8"))
	b.WriteString("\n" + s.Render("Search: "+m.Search))
	if m.Open {
		for i, opt := range m.Filtered {
			cursor := "  "
			if i == m.Index {
				cursor = "> "
			}
			b.WriteString(s.Render(cursor + opt.Name + "\n"))
		}
	}
	return b.String()
}

func (m *selectTaskModel) SetFocused(f bool) {
	m.Focused = f
	m.Search = ""
	m.Open = f
}

func (m *selectTaskModel) GetSelected() string {
	if m.Selected == nil {
		return "None"
	}
	return m.Selected.Name
}

func (m *selectTaskModel) GetTaskID() *uuid.UUID {
	if m.Selected == nil {
		return nil
	}
	id := m.Selected.ID.(uuid.UUID)
	return &id
}

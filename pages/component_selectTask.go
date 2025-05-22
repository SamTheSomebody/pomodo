package pages

import (
	"context"
	"log"
	"pomodo/helpers"
	"pomodo/internal/database"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/google/uuid"
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
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if !m.Focused {
			return m, nil
		}

		switch msg.Type {
		case tea.KeyEnter:
			if m.Open && len(m.Filtered) > 0 {
				m.Selected = &m.Filtered[m.Index]
				m.Open = false
				m.Search = ""
			} else {
				m.Open = true
			}

		case tea.KeyUp:
			if m.Open && m.Index > 0 {
				m.Index--
			}

		case tea.KeyDown:
			if m.Open && m.Index < len(m.Filtered)-1 {
				m.Index++
			}

		case tea.KeyBackspace:
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
	selectedText := "None"
	if m.Selected != nil {
		selectedText = m.Selected.Name
	}

	header := lipgloss.NewStyle().Bold(true).Render("Select Task: ")
	input := "[" + selectedText + " ]"

	if !m.Focused {
		return header + input
	}
	input += "\n"
	input += lipgloss.NewStyle().Italic(true).Render("Search:" + m.Search)
	view := header + input + "\n"
	if m.Open {
		for i, opt := range m.Filtered {
			cursor := "  "
			if i == m.Index {
				cursor = "> "
			}
			view += cursor + opt.Name + "\n"
		}
	}
	return view
}

// Public API
func (m *selectTaskModel) SetFocused(f bool) {
	m.Focused = f
	m.Search = ""
	m.Open = f
}

func (m *selectTaskModel) GetSelected() string {
	return m.Selected.Name
}

func (m *selectTaskModel) GetTaskID() *uuid.UUID {
	id := m.Selected.ID.(uuid.UUID)
	return &id
}

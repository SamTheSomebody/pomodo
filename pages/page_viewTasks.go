package pages

import (
	"context"
	"fmt"
	"os"
	"pomodo/helpers"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
	"golang.org/x/term"
)

/*
Name |
*/

type viewTasksModel struct {
	state *State
	table table.Model
}

func InitialViewTasksModel(s *State) tea.Model {
	m := viewTasksModel{
		state: s,
	}
	db := helpers.GetDBQueries()
	tasks, err := db.GetTasks(context.Background())
	if err != nil {
		s.Err = err
		// return m
	}
	rows := make([]table.Row, len(tasks))
	for i, task := range tasks {
		raw := helpers.Raw(task)
		rows[i] = []string{
			raw.ID,
			raw.Name,
			raw.Summary,
			raw.DueAt,
			raw.TimeEstimate,
		}
	}
	t := table.New(
		table.WithFocused(true),
		table.WithWidth(50),
		table.WithRows(rows),
		table.WithColumns([]table.Column{
			{Title: "ID", Width: 0},
			{Title: "Name", Width: 20},
			{Title: "Summary", Width: 30},
			{Title: "Due At", Width: 15},
			{Title: "Time Estimate", Width: 15},
		}),
	)
	width, height, _ := term.GetSize(int(os.Stdout.Fd()))
	t.SetWidth(width)
	height = min(height, len(tasks)+1)
	t.SetHeight(height)
	m.table = t
	s.Log = fmt.Sprintf("Adding %T to %T, State: %p, Navigation State: %p", m, s, &s, &s.Navigation.State)
	s.Navigation.Add(m)
	return m
}

func OnViewTasksButtonClick(s *State) func() (tea.Model, tea.Cmd) {
	return func() (tea.Model, tea.Cmd) {
		return InitialViewTasksModel(s), nil
	}
}

func (m viewTasksModel) Init() tea.Cmd {
	return nil
}

func (m viewTasksModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.state.Message = msg
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			id, err := uuid.Parse(m.table.SelectedRow()[0])
			if err != nil {
				m.state.Err = err
				return m, nil
			}
			return InitialEditTaskModel(m.state, &id), nil
		case "delete":
			id, err := uuid.Parse(m.table.SelectedRow()[0])
			m.state.Log = "Deleting: " + m.table.SelectedRow()[0]
			if err != nil {
				m.state.Err = err
				return m, nil
			}
			db := helpers.GetDBQueries()
			err = db.DeleteTask(context.TODO(), id)
			if err != nil {
				m.state.Err = err
				return m, nil
			}
			return InitialViewTasksModel(m.state), nil
		case "b", "esc", "q":
			return m.state.Navigation.Back()
		case "ctrl+c":
			return nil, tea.Quit
		}
	}
	m.table, _ = m.table.Update(msg)
	var cmd tea.Cmd
	return m, cmd
}

func (m viewTasksModel) View() string {
	b := strings.Builder{}
	b.WriteString(fmt.Sprint("TABLE with ", len(m.table.Rows()), " entries!\n"))
	b.WriteString(m.table.View())
	return m.state.View(b.String())
}

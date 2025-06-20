package pages

import (
	"context"
	"fmt"
	"log"
	"os"
	"pomodo/bubbletea"
	"pomodo/helpers"
	"slices"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
	"golang.org/x/term"
)

/*
Name |
*/

type ViewTasksPage struct {
	Table table.Model
}

func NewViewTasksPage() tea.Model {
	db := helpers.GetDBQueries()
	tasks, err := db.GetTasks(context.Background())
	if err != nil {
		log.Fatal(err) // TODO return a cmd
	}
	rows := make([]table.Row, len(tasks))
	for i, task := range tasks {
		rows[i] = []string{
			task.ID.(string),
			task.Name,
			task.Summary,
			helpers.ParseTime(task.DueAt),
			helpers.ParseDuration(task.TimeEstimateSeconds),
		}
	}
	t := table.New(
		table.WithKeyMap(table.DefaultKeyMap()),
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
	height = min(height, len(tasks)+1)
	t.SetWidth(width)
	t.SetHeight(height)
	return ViewTasksPage{Table: t}
}

func (m ViewTasksPage) Init() tea.Cmd {
	return nil
}

func (m ViewTasksPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			id, err := uuid.Parse(m.Table.SelectedRow()[0])
			if err != nil {
				return m, bubbletea.ErrCmd(err)
			}
			return m, bubbletea.NewPageCmd(func() (tea.Model, tea.Cmd) { return NewEditTaskPage(&id), nil })
		case "delete":
			id, err := uuid.Parse(m.Table.SelectedRow()[0])
			if err != nil {
				return m, bubbletea.ErrCmd(err)
			}
			db := helpers.GetDBQueries()
			err = db.DeleteTask(context.TODO(), id)
			if err != nil {
				return m, bubbletea.ErrCmd(err)
			}
			m.Table.SetRows(slices.Delete(m.Table.Rows(), m.Table.Cursor(), 1))
			return m, bubbletea.LogCmd("Deleted: " + m.Table.SelectedRow()[0])
		}
	}
	m.Table, _ = m.Table.Update(msg)
	var cmd tea.Cmd
	return m, cmd
}

func (m ViewTasksPage) View() string {
	b := strings.Builder{}
	b.WriteString(fmt.Sprint("Table with ", len(m.Table.Rows()), " entries!\n"))
	b.WriteString(m.Table.View())
	return b.String()
}

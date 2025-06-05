package pages

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
	"golang.org/x/term"

	"pomodo/bubbletea"
	"pomodo/helpers"
)

/*
Name |
*/

type ViewTasksPage struct {
	Table  table.Model
	KeyMap *bubbletea.Keymap
}

func NewViewTasksPage(keymap *bubbletea.Keymap) tea.Model {
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
	return ViewTasksPage{Table: t, KeyMap: keymap}
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
			return NewEditTaskPage(&id, m.KeyMap), nil // TODO goto page command
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
			return NewViewTasksPage(m.KeyMap), bubbletea.LogCmd("Deleted: " + m.Table.SelectedRow()[0]) // TODO goto page command
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

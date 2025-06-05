package pages

import (
	"context"
	"fmt"
	"pomodo/bubbletea"
	"pomodo/bubbletea/button"
	"pomodo/bubbletea/list"
	"pomodo/bubbletea/slider"
	"pomodo/helpers"
	"pomodo/internal/database"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

/* Visual
Editing Task: (ID)
>   Name: [text]
    Summary: [text area when editing] [shortend when not]
    Due At: [time date]
    Time Estimate: [XXh XXm]
    Time Spent: [XXh XXm XXs]
    Priority: [int 1-10, 0 is not set]
    Enthusiasm: [int 1-10, 0 is not set]
*/

type EditTaskPage struct {
	Task    database.Task
	List    list.Model
	Focus   int
	HasTask bool
}

func NewEditTaskPage(taskID *uuid.UUID, keymap *bubbletea.Keymap) EditTaskPage {
	hasTask := taskID != nil
	var task database.Task
	if !hasTask {
		task = database.Task{
			ID: uuid.New(),
		}
	} else {
		var err error
		task, err = helpers.GetDBQueries().GetTaskByID(context.Background(), taskID)
		if err != nil {
			// s.Err = err // TODO send a cmd up the chain
			task = database.Task{
				ID: uuid.New(),
			}
		}
	}
	m := EditTaskPage{
		Task:    task,
		HasTask: hasTask,
	}

	values := []struct {
		prompt, placeholder, value string
		validate                   func(string) error
	}{
		{"Name:          ", "Task Name", m.Task.Name, func(string) error { return nil }},
		{"Summary:       ", "Write a summary... ", m.Task.Summary, func(string) error { return nil }},
		{"Due At:        ", "YYYY-MM-DD HH:MM", helpers.ParseTime(m.Task.DueAt), helpers.ValidateTime},
		{"Time Estimate: ", "e.g. 1h30m", helpers.ParseDuration(m.Task.TimeEstimateSeconds), helpers.ValidateDuration},
		{"Time Spent:    ", "e.g. 2h15m30s", helpers.ParseDuration(m.Task.TimeSpentSeconds), helpers.ValidateDuration},
	}

	items := make([]list.Item, len(values))

	for i, v := range values {
		input := textinput.New()
		input.SetValue(v.value)
		input.Placeholder = v.placeholder
		input.Prompt = v.prompt
		input.Validate = v.validate
		input.Width = 50 // TODO make this a sane automated size
		items[i] = list.TextInputItem{Input: input}
	}

	items = append(items,
		slider.New("Priority:      ", int(m.Task.Priority)),
		slider.New("Enthusiasm:    ", int(m.Task.Enthusiasm)),
		button.New("Confirm", func() (tea.Model, tea.Cmd) { return m.Submit() }))
	m.List = list.New(items, keymap)
	return m
}

func (m EditTaskPage) Init() tea.Cmd {
	return textinput.Blink
}

func (m EditTaskPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m.List.Update(msg)
}

func (m EditTaskPage) Submit() (tea.Model, tea.Cmd) {
	name := m.List.Items[0].(list.TextInputItem).Input.Value()
	summary := m.List.Items[1].(list.TextInputItem).Input.Value()
	dueAt, err := time.Parse("DD/MM/YY", m.List.Items[2].(list.TextInputItem).Input.Value())
	timeEstimate, err := time.ParseDuration(m.List.Items[3].(list.TextInputItem).Input.Value())
	// TODO timeSpent, err := time.ParseDuration(m.List.Items[4].(list.TextInputItem).Input.Value())
	priority := m.List.Items[5].(slider.Model).Value
	enthusiasm := m.List.Items[6].(slider.Model).Value

	db := helpers.GetDBQueries()
	if m.HasTask {
		_, err = db.SetTask(context.TODO(), database.SetTaskParams{
			ID:                  m.Task.ID,
			Name:                name,
			Summary:             summary,
			DueAt:               dueAt,
			TimeEstimateSeconds: int64(timeEstimate.Seconds()),
			// TODO TimeSpentSeconds:    int64(timeSpent.Seconds()),
			Priority:   int64(priority),
			Enthusiasm: int64(enthusiasm),
		})
	} else {
		_, err = db.CreateTask(context.TODO(), database.CreateTaskParams{
			ID:                  m.Task.ID,
			Name:                name,
			Summary:             summary,
			DueAt:               dueAt,
			TimeEstimateSeconds: int64(timeEstimate.Seconds()),
			// TODO TimeSpentSeconds:    int64(timeSpent.Seconds()),
			Priority:   int64(priority),
			Enthusiasm: int64(enthusiasm),
		})
	}
	if err != nil {
		return m, bubbletea.ErrCmd(err)
	}
	return m, bubbletea.NewPageCmd(func() (tea.Model, tea.Cmd) { return NewHomePage(m.List.Keys), nil })
}

func (m EditTaskPage) View() string {
	b := strings.Builder{}
	if m.HasTask {
		b.WriteString("Editing")
	} else {
		b.WriteString("Adding")
	}
	b.WriteString(fmt.Sprintf(" Task (%v)\n\n", m.Task.ID))
	// b.WriteString(fmt.Sprintf("List - w: %v, h: %v, items: %v\n", m.list.Width(), m.list.Height(), len(m.list.Items())))
	b.WriteString(m.List.View())
	return b.String()
}

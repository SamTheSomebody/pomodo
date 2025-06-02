package pages

import (
	"context"
	"fmt"
	"pomodo/bubbletea"
	"pomodo/bubbletea/button"
	"pomodo/bubbletea/list"
	"pomodo/helpers"
	"pomodo/internal/database"
	"strings"

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
	Task    helpers.RawTask
	List    list.Model
	Focus   int
	HasTask bool
}

func NewEditTaskPage(taskID *uuid.UUID) EditTaskPage {
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
		Task:    helpers.Raw(task),
		HasTask: hasTask,
	}

	values := []struct{ prompt, placeholder, value string }{
		{"Name:          ", "Task Name", m.Task.Name},
		{"Summary:       ", "Write a summary... ", m.Task.Summary},
		{"Due At:        ", "YYYY-MM-DD HH:MM", m.Task.DueAt},
		{"Time Estimate: ", "e.g. 1h30m", m.Task.TimeEstimate},
		{"Time Spent:    ", "e.g. 2h15m30s", m.Task.TimeSpent},
		{"Priority:      ", "1-10", m.Task.Priority},
		{"Enthusiasm:    ", "1-10", m.Task.Enthusiasm},
	}

	items := make([]list.Item, len(values)+1)

	for i, v := range values {
		input := textinput.New()
		input.SetValue(v.value)
		input.Placeholder = v.placeholder
		input.Prompt = v.prompt
		input.Width = 50 // TODO make this a sane automated size
		item := list.TextInputItem{Input: input}
		items[i] = list.NewItem(item)
	}

	confirmButton := button.New("Confirm", confirmButton(&m))
	items[len(values)] = list.NewItem(confirmButton)
	m.List = list.New(items)
	return m
}

func OnEditTaskButtonClick(taskID *uuid.UUID) func() (tea.Model, tea.Cmd) {
	return func() (tea.Model, tea.Cmd) {
		return NewEditTaskPage(taskID), nil
	}
}

func (m EditTaskPage) Init() tea.Cmd {
	return textinput.Blink
}

func (m EditTaskPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	model, cmd := m.List.Update(msg)
	switch t := model.(type) {
	case list.Model:
		m.List = t
	default:
		return model, cmd
	}
	return m, nil
}

func (m EditTaskPage) Submit() (tea.Model, tea.Cmd) {
	m.Task.Name = m.List.Items[0].Model.(list.TextInputItem).Input.Value()
	m.Task.Summary = m.List.Items[1].Model.(list.TextInputItem).Input.Value()
	m.Task.DueAt = m.List.Items[2].Model.(list.TextInputItem).Input.Value()
	m.Task.TimeEstimate = m.List.Items[3].Model.(list.TextInputItem).Input.Value()
	m.Task.TimeSpent = m.List.Items[4].Model.(list.TextInputItem).Input.Value()
	m.Task.Priority = m.List.Items[5].Model.(list.TextInputItem).Input.Value()
	m.Task.Enthusiasm = m.List.Items[6].Model.(list.TextInputItem).Input.Value()
	var err error
	if m.HasTask {
		err = helpers.EditTask(m.Task)
	} else {
		err = helpers.AddTask(m.Task)
	}
	if err != nil {
		return m, bubbletea.ErrCmd(err)
	}
	return m, nil // TODO close page command
}

func confirmButton(m *EditTaskPage) func() (tea.Model, tea.Cmd) {
	return func() (tea.Model, tea.Cmd) {
		return m.Submit()
	}
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

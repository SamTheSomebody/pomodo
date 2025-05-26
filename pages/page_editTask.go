package pages

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"

	"pomodo/helpers"
	"pomodo/internal/database"
)

/* Visual
Editing task: (ID)
>   Name: [text]
    Summary: [text area when editing] [shortend when not]
    Due At: [time date]
    Time Estimate: [XXh XXm]
    Time Spent: [XXh XXm XXs]
    Priority: [int 1-10, 0 is not set]
    Enthusiasm: [int 1-10, 0 is not set]
*/

type editTaskModel struct {
	state   *State
	task    helpers.RawTask
	inputs  []textinput.Model
	focus   int
	hasTask bool
}

func InitialEditTaskModel(s *State, task database.Task) editTaskModel {
	hasTask := task.ID != nil
	if !hasTask {
		task.ID = uuid.New()
	}

	m := editTaskModel{
		task:   helpers.Raw(task),
		inputs: make([]textinput.Model, 7),
	}

	placeholders := []string{
		"Task Name", "Write a summary... ", "YYYY-MM-DD HH:MM", "e.g. 1h30m", "e.g. 2h15m30s", "1-10", "1-10",
	}

	values := []string{
		m.task.Name, m.task.Summary, m.task.DueAt,
		m.task.TimeEstimate, m.task.TimeSpent, m.task.Priority, m.task.Enthusiasm,
	}

	for i := range m.inputs {
		m.inputs[i] = textinput.New()
		m.inputs[i].Placeholder = placeholders[i]
		m.inputs[i].SetValue(values[i])
		m.inputs[i].Prompt = ""
		m.inputs[i].Width = 50
	}

	m.inputs[0].Focus()
	m.state = s
	m.state.Navigation.Add(m)
	return m
}

func (m editTaskModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m editTaskModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "shift+tab", "up":
			return m.AdjustFocus(-1)
		case "tab", "down":
			return m.AdjustFocus(1)
		case "enter":
			if m.focus == len(m.inputs)-1 {
				return m.Submit()
			}
			m.AdjustFocus(1)
			return m, nil
		}
	}

	if mod, cmd := m.state.ProcessUniversalKeys(msg); mod != nil || cmd != nil {
		return mod, cmd
	}

	m.inputs[m.focus], _ = m.inputs[m.focus].Update(msg)
	return m, nil
}

func (m editTaskModel) Submit() (tea.Model, tea.Cmd) {
	m.task.Name = m.inputs[0].Value()
	m.task.Summary = m.inputs[1].Value()
	m.task.DueAt = m.inputs[2].Value()
	m.task.TimeEstimate = m.inputs[3].Value()
	m.task.TimeSpent = m.inputs[4].Value()
	m.task.Priority = m.inputs[5].Value()
	m.task.Enthusiasm = m.inputs[6].Value()
	var err error
	if m.hasTask {
		err = helpers.EditTask(m.task)
	} else {
		err = helpers.AddTask(m.task)
	}
	if err != nil {
		m.state.Err = err
		return m, nil
	}

	return m.state.Navigation.Back()
}

func (m editTaskModel) AdjustFocus(amount int) (tea.Model, tea.Cmd) {
	m.inputs[m.focus].Blur()
	m.focus += amount + len(m.inputs)
	m.focus %= len(m.inputs)
	m.inputs[m.focus].Focus()
	return m, nil
}

func (m editTaskModel) View() string {
	b := strings.Builder{}
	if m.hasTask {
		b.WriteString("Editing")
	} else {
		b.WriteString("Adding")
	}
	b.WriteString(fmt.Sprintf(" task (%v)\n\n", m.task.ID))
	b.WriteString("Name:          " + m.inputs[0].View() + "\n")
	b.WriteString("Summary:       " + m.inputs[1].View() + "\n")
	b.WriteString("Due At:        " + m.inputs[2].View() + "\n")
	b.WriteString("Time Estimate: " + m.inputs[3].View() + "\n")
	b.WriteString("Time Spent:    " + m.inputs[4].View() + "\n")
	b.WriteString("Priority:      " + m.inputs[5].View() + "\n")
	b.WriteString("Enthusiasm:    " + m.inputs[6].View() + "\n")
	return m.state.View(b.String())
}

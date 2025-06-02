package pages

import (
	"pomodo/bubbletea/button"
	"pomodo/bubbletea/taskselect"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

const focusCount = 3

type ConfigureTimerPage struct {
	TaskID     *uuid.UUID
	TimerInput textinput.Model
	TaskSelect taskselect.Model
	Button     button.Model
	Focus      int
	Duration   time.Duration
}

func NewConfigureTimerPage(t *uuid.UUID) ConfigureTimerPage {
	m := ConfigureTimerPage{
		TaskID:     t,
		TimerInput: textinput.New(),
		TaskSelect: taskselect.New(),
		Focus:      0,
	}
	m.TimerInput.Prompt = "Duration: "
	m.TimerInput.Placeholder = "XXh XXm XXs"
	m.TimerInput.SetValue("")
	m.TimerInput.Width = 50
	m.TimerInput.Focus()
	m.Button = button.New("Confirm", OnTimerButtonClick(m.Duration, m.TaskID))
	return m
}

func OnConfigureTimerModelButtonClick(TaskID *uuid.UUID) func() (tea.Model, tea.Cmd) {
	return func() (tea.Model, tea.Cmd) {
		return NewConfigureTimerPage(TaskID), nil
	}
}

func (m ConfigureTimerPage) Init() tea.Cmd {
	return textinput.Blink
}

func (m ConfigureTimerPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.TaskSelect.Focused {
		mod, _ := m.TaskSelect.Update(msg)
		m.TaskSelect = mod.(taskselect.Model)
		if m.TaskSelect.Focused {
			return m, nil
		}
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if m.Focus == focusCount-1 {
				return m.Button.Update(msg)
			}
			return m.setFocus(1)
		case "tab", "down":
			return m.setFocus(1)
		case "shift+tab", "up":
			return m.setFocus(-1)
		}
	}

	switch m.Focus {
	case 0:
		m.TimerInput, _ = m.TimerInput.Update(msg)
	case 1:
		mod, _ := m.TaskSelect.Update(msg)
		m.TaskSelect = mod.(taskselect.Model)
	}
	return m, nil
}

func (m ConfigureTimerPage) setFocus(increment int) (tea.Model, tea.Cmd) {
	switch m.Focus {
	case 0:
		m.Duration, _ = time.ParseDuration(m.TimerInput.Value())
		m.TimerInput.Blur()
	case 1:
		m.TaskID = m.TaskSelect.GetTaskID()
		m.TaskSelect.SetFocused(false)
	case 2:
		m.Button.SetFocused(false)
	}

	m.Focus += increment + focusCount
	m.Focus %= focusCount
	m.Button.SetFocused(false)
	switch m.Focus {
	case 0:
		m.TimerInput.Focus()
	case 1:
		m.TaskSelect.SetFocused(true)
	case 2:
		m.Button.SetFocused(true)
	}
	return m, nil
}

func (m ConfigureTimerPage) View() string {
	b := strings.Builder{}
	b.WriteString("Timer Configuation:" + "\n")
	b.WriteString(m.TimerInput.View() + "\n")
	b.WriteString(m.TaskSelect.View() + "\n")
	b.WriteString(m.Button.View() + "\n")
	return b.String()
}

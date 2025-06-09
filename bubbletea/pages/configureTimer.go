package pages

import (
	"pomodo/bubbletea/button"
	"pomodo/bubbletea/list"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

type ConfigureTimerPage struct {
	TaskID   *uuid.UUID
	List     list.Model
	Duration time.Duration
}

func NewConfigureTimerPage(t *uuid.UUID) ConfigureTimerPage {
	m := ConfigureTimerPage{
		TaskID: t,
	}
	timerInput := textinput.New()
	timerInput.Prompt = "Duration: "
	timerInput.Placeholder = "XXh XXm XXs"
	timerInput.Width = 50
	timerInput.Validate = func(s string) error {
		_, err := time.ParseDuration(s)
		return err
	}
	confirmButton := button.New("Confirm", OnTimerButtonClick(m.Duration, m.TaskID))
	// TODO Task select
	list := list.New([]list.Item{list.NewTextInput(timerInput), confirmButton})
	m.List = list
	return m
}

func (m ConfigureTimerPage) Init() tea.Cmd {
	return textinput.Blink
}

func (m ConfigureTimerPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	model, cmd := m.List.Update(msg)
	m.List = model.(list.Model)
	return m, cmd
}

func (m ConfigureTimerPage) View() string {
	b := strings.Builder{}
	b.WriteString("Timer Configuation:" + "\n")
	b.WriteString(m.List.View() + "\n")
	return b.String()
}

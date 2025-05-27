package pages

import (
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

const focusCount = 3

type configureTimerModel struct {
	state      *State
	taskID     *uuid.UUID
	timerInput textinput.Model
	selectTask selectTaskModel
	button     buttonModel
	focus      int
	duration   time.Duration
}

func InitialConfigureTimerModel(s *State, t *uuid.UUID) configureTimerModel {
	m := configureTimerModel{
		state:      s,
		taskID:     t,
		timerInput: textinput.New(),
		selectTask: InitialSelectTaskModel(s),
		focus:      0,
	}
	m.timerInput.Prompt = "Duration: "
	m.timerInput.Placeholder = "XXh XXm XXs"
	m.timerInput.SetValue("")
	m.timerInput.Width = 50
	m.timerInput.Focus()
	m.button = InitialButtonModel("Confirm", InitialTimerModel(m.state, m.duration, m.taskID), nil) // Does this return the correct data?
	s.Navigation.Add(m)
	return m
}

func (m configureTimerModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m configureTimerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.state.Message = msg
	if m.selectTask.Focused {
		mod, _ := m.selectTask.Update(msg)
		m.selectTask = mod.(selectTaskModel)
		if m.selectTask.Focused {
			return m, nil
		}
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if m.focus == focusCount-1 {
				return m.button.Update(msg)
			}
			return m.setFocus(1)
		case "tab", "down":
			return m.setFocus(1)
		case "shift+tab", "up":
			return m.setFocus(-1)
		case "esc":
			return m.state.Navigation.Back()
		case "ctrl+c":
			return nil, tea.Quit
		}
	}

	switch m.focus {
	case 0:
		m.timerInput, _ = m.timerInput.Update(msg)
	case 1:
		mod, _ := m.selectTask.Update(msg)
		m.selectTask = mod.(selectTaskModel)
	}
	return m, nil
}

func (m configureTimerModel) setFocus(increment int) (tea.Model, tea.Cmd) {
	switch m.focus {
	case 0:
		m.duration, m.state.Err = time.ParseDuration(m.timerInput.Value())
		m.timerInput.Blur()
	case 1:
		m.taskID = m.selectTask.GetTaskID()
		m.selectTask.SetFocused(false)
	case 2:
		m.button.SetFocused(false)
	}

	m.focus += increment + focusCount
	m.focus %= focusCount
	m.button.SetFocused(false)
	switch m.focus {
	case 0:
		m.timerInput.Focus()
	case 1:
		m.selectTask.SetFocused(true)
	case 2:
		m.button.SetFocused(true)
	}
	return m, nil
}

func (m configureTimerModel) View() string {
	b := strings.Builder{}
	b.WriteString("Timer Configuation:" + "\n")
	b.WriteString(m.timerInput.View() + "\n")
	b.WriteString(m.selectTask.View() + "\n")
	b.WriteString(m.button.View() + "\n")
	return m.state.View(b.String())
}

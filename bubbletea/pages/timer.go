package pages

import (
	"context"
	"log"
	"pomodo/helpers"
	"pomodo/internal/database"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

type TimerPage struct {
	Task            *database.Task
	Timer           timer.Model
	Progress        progress.Model
	CurrentProgress float64
	Timeout         time.Duration
}

// TODO For some reason this doesn't start when initialized, and runs twice as fast when first manually started
func NewTimerPage(duration time.Duration, taskID *uuid.UUID) TimerPage {
	var task database.Task
	if taskID != nil {
		var err error
		task, err = helpers.GetDBQueries().GetTaskByID(context.TODO(), taskID)
		if err != nil {
			log.Fatal(err)
		}
	}
	return TimerPage{
		Task:            &task,
		Timer:           timer.New(duration),
		Progress:        progress.New(),
		CurrentProgress: 0,
		Timeout:         duration,
	}
}

func OnTimerButtonClick(duration time.Duration, TaskID *uuid.UUID) func() (tea.Model, tea.Cmd) {
	return func() (tea.Model, tea.Cmd) {
		return NewTimerPage(duration, TaskID), nil
	}
}

func (m TimerPage) Init() tea.Cmd {
	return m.Timer.Init()
}

func (m TimerPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case timer.TickMsg:
		if !m.Timer.Timedout() {
			m.CurrentProgress = 1 - float64(m.Timer.Timeout)/float64(m.Timeout)
		} else {
			m.CurrentProgress = 1
		}
		var cmd tea.Cmd
		m.Timer, cmd = m.Timer.Update(msg)
		return m, cmd
	case tea.WindowSizeMsg:
		m.Progress.Width = min(msg.Width-padding*2-4, maxWidth)
		return m, nil
	case timer.StartStopMsg:
		var cmd tea.Cmd
		m.Timer, cmd = m.Timer.Update(msg)
		return m, cmd
	case timer.TimeoutMsg:
		// TODO m.State.Keys.Enter.SetEnabled(true)
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "r":
			m.Timer.Timeout = m.Timeout
			m.Timer.Start()
		case "s":
			return m, m.Timer.Toggle()
		case "enter":
			if m.Timer.Timedout() {
				return NewHomePage(), nil // TODO add time to Task // TODO new page command
			}
		}
	}
	return m, nil
}

func (m TimerPage) View() string {
	b := strings.Builder{}
	if m.Task != nil {
		b.WriteString(m.Task.Name + ", " + m.Task.Summary)
	}
	b.WriteString(m.Timer.Timeout.String())
	b.WriteString(m.Progress.ViewAs(m.CurrentProgress) + "\n")

	if m.Timer.Timedout() {
		b.Reset()
		b.WriteString("Finished!")
	}

	return b.String()
}

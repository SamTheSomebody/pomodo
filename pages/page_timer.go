package pages

import (
	"context"
	"log"
	"pomodo/helpers"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

type timerModel struct {
	task          tea.Model
	state         *State
	timerModel    timer.Model
	progressModel progress.Model
	progress      float64
	timeout       time.Duration
}

func (m timerModel) Init() tea.Cmd {
	return m.timerModel.Init()
}

func (m timerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.state.Message = msg
	switch msg := msg.(type) {
	case timer.TickMsg:
		if !m.timerModel.Timedout() {
			m.progress = 1 - float64(m.timerModel.Timeout)/float64(m.timeout)
		} else {
			m.progress = 1
		}
		var cmd tea.Cmd
		m.timerModel, cmd = m.timerModel.Update(msg)
		return m, cmd
	case tea.WindowSizeMsg:
		m.progressModel.Width = min(msg.Width-padding*2-4, maxWidth)
		return m, nil
	case timer.StartStopMsg:
		var cmd tea.Cmd
		m.timerModel, cmd = m.timerModel.Update(msg)
		return m, cmd
	case timer.TimeoutMsg:
		m.state.Keys.Enter.SetEnabled(true)
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return nil, tea.Quit
		case "esc":
			return m.state.Navigation.Back()
		case "r":
			m.timerModel.Timeout = m.timeout
			m.timerModel.Start()
		case "s":
			return m, m.timerModel.Toggle()
		case "enter":
			if m.timerModel.Timedout() {
				return InitialHomeModel(m.state), nil // TODO add time to task
			}
		}
	}
	return m, nil
}

func (m timerModel) View() string {
	b := strings.Builder{}
	if m.task != nil {
		b.WriteString(m.task.View())
	}
	b.WriteString(m.timerModel.Timeout.String())
	b.WriteString(m.progressModel.ViewAs(m.progress) + "\n")

	if m.timerModel.Timedout() {
		b.Reset()
		b.WriteString("Finished!")
	}

	return m.state.View(b.String())
}

// TODO For some reason this doesn't start when initialized, and runs twice as fast when first manually started
func InitialTimerModel(s *State, duration time.Duration, taskID *uuid.UUID) timerModel {
	var task tea.Model
	if taskID != nil {
		t, err := helpers.GetDBQueries().GetTaskByID(context.TODO(), taskID)
		if err != nil {
			log.Fatal(err)
		}
		rawTask := helpers.Raw(t)
		task = InitialTaskModel(&rawTask)
	}

	m := timerModel{
		task:          task,
		timerModel:    timer.New(duration),
		progressModel: progress.New(),
		progress:      0,
		timeout:       duration,
		state:         s,
	}
	s.Navigation.Add(m)
	return m
}

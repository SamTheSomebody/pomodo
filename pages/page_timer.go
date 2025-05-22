package pages

import (
	"context"
	"fmt"
	"log"
	"pomodo/helpers"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

type timerModel struct {
	task          tea.Model
	state         *State
	keymap        timerKeyMap
	timerModel    timer.Model
	progressModel progress.Model
	progress      float64
	timeout       time.Duration
}

type timerKeyMap struct {
	start key.Binding
	stop  key.Binding
	reset key.Binding
}

func (m timerModel) Init() tea.Cmd {
	return m.timerModel.Init()
}

func (m timerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if mod, cmd := m.state.ProcessUniversalKeys(msg); mod != nil || cmd != nil {
		return mod, cmd
	}

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
		m.progressModel.Width = min(msg.Width-len(padding)*2-4, maxWidth)
		return m, nil
	case timer.StartStopMsg:
		var cmd tea.Cmd
		m.timerModel, cmd = m.timerModel.Update(msg)
		m.keymap.stop.SetEnabled(m.timerModel.Running())
		m.keymap.start.SetEnabled(!m.timerModel.Running())
		return m, cmd
	case timer.TimeoutMsg:
		m.state.Keys.Enter.SetEnabled(true)
		return m, nil
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.state.Keys.Back):
			return m.state.Navigation.Back()
		case key.Matches(msg, m.keymap.reset):
			m.timerModel.Timeout = m.timeout
			m.timerModel.Start()
		case key.Matches(msg, m.keymap.start, m.keymap.stop):
			return m, m.timerModel.Toggle()
		case key.Matches(msg, m.state.Keys.Enter):
			return InitialHomeModel(m.state), nil // TODO add time to task
		}
	}
	return m, nil
}

func (m timerModel) View() string {
	s := ""
	if m.task != nil {
		s += m.task.View()
	}
	s += fmt.Sprint("\n", padding, m.timerModel.Timeout)
	s += fmt.Sprint(padding, m.progressModel.ViewAs(m.progress), "\n")

	if m.timerModel.Timedout() {
		s = "\n" + padding + "Finished!\n"
	}
	b := []key.Binding{m.keymap.reset, m.keymap.start, m.keymap.stop}
	s += m.state.HelpView(b...)
	return s
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
		task: task,
		keymap: timerKeyMap{
			start: key.NewBinding(
				key.WithKeys("s"),
				key.WithHelp("s", "start"),
				key.WithDisabled(),
			),
			stop: key.NewBinding(
				key.WithKeys("s"),
				key.WithHelp("s", "stop"),
			),
			reset: key.NewBinding(
				key.WithKeys("r"),
				key.WithHelp("r", "reset"),
			),
		},
		timerModel:    timer.New(duration),
		progressModel: progress.New(),
		progress:      0,
		timeout:       duration,
		state:         s,
	}
	s.Navigation.Add(m)
	return m
}

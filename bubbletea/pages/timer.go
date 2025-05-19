package pages

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	padding  = 2
	maxWidth = 80
)

type timerModel struct {
	timerModel    timer.Model
	progressModel progress.Model
	progress      float64
	timeout       time.Duration
	keymap        keymap
	help          help.Model
	quitting      bool
}

type keymap struct {
	start key.Binding
	stop  key.Binding
	reset key.Binding
	quit  key.Binding
	enter key.Binding
}

func (m timerModel) Init() tea.Cmd {
	return m.timerModel.Init()
}

func (m timerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		m.keymap.stop.SetEnabled(m.timerModel.Running())
		m.keymap.start.SetEnabled(!m.timerModel.Running())
		return m, cmd
	case timer.TimeoutMsg:
		m.keymap.enter.SetEnabled(true)
		return m, nil
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.quit):
			return InitialHomeModel(), nil
		case key.Matches(msg, m.keymap.reset):
			m.timerModel.Timeout = m.timeout
			m.timerModel.Start()
		case key.Matches(msg, m.keymap.start, m.keymap.stop):
			return m, m.timerModel.Toggle()
		case key.Matches(msg, m.keymap.enter):
			return InitialHomeModel(), nil // TODO add time to task
		}
	}
	return m, nil
}

func (m timerModel) helpView() string {
	return "\n" + m.help.ShortHelpView([]key.Binding{
		m.keymap.start,
		m.keymap.stop,
		m.keymap.reset,
		m.keymap.quit,
		m.keymap.enter,
	})
}

func (m timerModel) View() string {
	pad := strings.Repeat(" ", padding)
	s := fmt.Sprint("\n", pad, m.timerModel.Timeout, "\n")
	s += fmt.Sprint("\n", pad, m.progressModel.ViewAs(m.progress), "\n")

	if m.timerModel.Timedout() {
		s = "Finished!"
	}
	s += "\n"
	s += m.helpView()
	return s
}

// TODO For some reason this doesn't start when initialized, and runs twice as fast when first manually started
func InitialTimerModel(duration time.Duration) timerModel {
	return timerModel{
		timerModel:    timer.New(duration),
		progressModel: progress.New(),
		progress:      0,
		timeout:       duration,
		keymap: keymap{
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
			quit: key.NewBinding(
				key.WithKeys("q"),
				key.WithHelp("q", "quit"),
			),
			enter: key.NewBinding(
				key.WithKeys("enter"),
				key.WithHelp("enter", "continue"),
				key.WithDisabled(),
			),
		},
		help: help.New(),
	}
}

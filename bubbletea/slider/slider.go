package slider

import (
	"pomodo/bubbletea"
	"pomodo/bubbletea/list"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Min      int
	Max      int
	Value    int
	OldValue int
	Label    string
	Keymap   Keymap
}

func New(label string, value int) Model {
	if value < 1 || value > 10 { // TODO add an 'unset' value
		value = 5
	}
	return Model{
		Min:      1,
		Max:      10,
		Label:    label,
		Value:    value,
		OldValue: value,
		Keymap:   DefaultKeymap(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, m.Keymap.CursorDown, m.Keymap.NextPage) {
			m.Value++
			if m.Value > m.Max {
				m.Value = m.Max
			}
			return m, cmd
		}
		if key.Matches(msg, m.Keymap.CursorDown, m.Keymap.PrevPage) {
			m.Value--
			if m.Value < m.Min {
				m.Value = m.Min
			}
			return m, cmd
		}
		if key.Matches(msg, m.Keymap.GoToEnd) {
			m.Value = m.Max
			return m, cmd
		}
		if key.Matches(msg, m.Keymap.GoToStart) {
			m.Value = m.Min
			return m, cmd
		}
		if n, err := strconv.Atoi(msg.String()); err != nil {
			if n >= m.Min && n <= m.Max {
				m.Value = n
			}
			return m, cmd
		}
	}
	return m, cmd
}

func (m Model) View() string {
	b := strings.Builder{}
	b.WriteString(m.Label)
	b.WriteRune('|')
	if m.Min < 0 && m.Value >= 0 {
		b.WriteRune(' ')
	}
	b.WriteString(strings.Repeat("-", m.Value-m.Min))
	b.WriteString(strconv.Itoa(m.Value))
	b.WriteString(strings.Repeat("-", m.Max-m.Value))
	if m.Max > 9 && m.Value <= 9 {
		b.WriteRune(' ')
	}
	b.WriteRune('|')
	return b.String()
}

func (m Model) OnSelect() (list.Item, tea.Cmd) {
	m.OldValue = m.Value
	return m, bubbletea.ItemSelectCmd(true)
}

func (m Model) OnSubmit() (list.Item, tea.Cmd) {
	return m, bubbletea.ItemSelectCmd(false)
}

func (m Model) OnCancel() (list.Item, tea.Cmd) {
	m.Value = m.OldValue
	return m, bubbletea.ItemSelectCmd(false)
}

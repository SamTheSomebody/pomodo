package list

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"

	"pomodo/bubbletea"
)

const buffer = 0

type Model struct {
	Index          int
	Items          []Item
	IsItemSelected bool
	Keys           *bubbletea.KeyMap
	Styles         Styles
}

func New(items []Item, keymap *bubbletea.KeyMap) Model {
	m := Model{
		Items:  items,
		Keys:   keymap,
		Styles: DefaultStyle(),
	}
	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.IsItemSelected {
			if key.Matches(msg, m.Keys.Cancel) {
				m.IsItemSelected = false
				m.Items[m.Index], _ = m.Items[m.Index].OnCancel()
				return m, bubbletea.EnableNavigationCmd(true)
			}
			if key.Matches(msg, m.Keys.Submit) {
				m.IsItemSelected = false
				m.Items[m.Index], _ = m.Items[m.Index].OnSubmit()
				return m, bubbletea.EnableNavigationCmd(true)
			}
			model, _ := m.Items[m.Index].Update(msg)
			m.Items[m.Index] = model.(Item)
			return m, bubbletea.EnableNavigationCmd(false)
		}
		if key.Matches(msg, m.Keys.CursorUp) {
			m.Index += len(m.Items) - 1
			m.Index %= len(m.Items)
			return m, nil
		}
		if key.Matches(msg, m.Keys.CursorDown) {
			m.Index++
			m.Index %= len(m.Items)
			return m, nil
		}
		if key.Matches(msg, m.Keys.GoToEnd) {
			m.Index = len(m.Items) - 1
			return m, nil
		}
		if key.Matches(msg, m.Keys.GoToStart) {
			m.Index = 0
			return m, nil
		}
		if key.Matches(msg, m.Keys.Select) {
			m.IsItemSelected = true
			var cmd tea.Cmd
			m.Items[m.Index], cmd = m.Items[m.Index].OnSelect()
			if cmd != nil {
				return m, cmd
			}
			return m, bubbletea.EnableNavigationCmd(false)
		}
	}
	return m, nil
}

func (m Model) View() string {
	b := strings.Builder{}
	b.WriteString(fmt.Sprintf("List with %v items, IsItemSelected: %v, Index: %v, Select Enabled: %v, Submit Enabled: %v\n",
		len(m.Items), m.IsItemSelected, m.Index, m.Keys.Select.Enabled(), m.Keys.Submit.Enabled()))
	for i, item := range m.Items {
		b.WriteString(m.Styles.Render(item, m.Index == i) + "\n")
	}
	return b.String()
}

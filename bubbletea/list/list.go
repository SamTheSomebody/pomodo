package list

import (
	"pomodo/bubbletea"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

const buffer = 0

type Model struct {
	Index          int
	Items          []Item
	IsItemSelected bool
	Keys           *bubbletea.Keymap
	Styles         Styles
}

func New(items []Item, keymap *bubbletea.Keymap) Model {
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
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case bubbletea.ItemSelectMsg:
		m.IsItemSelected = msg.IsSelected
		model, cmd := m.Items[m.Index].Update(msg)
		m.Items[m.Index] = model.(Item)
		return m, cmd
	case tea.KeyMsg:
		if key.Matches(msg, m.Keys.Cancel) {
			var cmd tea.Cmd
			m.Items[m.Index], cmd = m.Items[m.Index].OnCancel()
			return m, cmd
		}
		if key.Matches(msg, m.Keys.Submit) {
			var cmd tea.Cmd
			m.Items[m.Index], cmd = m.Items[m.Index].OnSubmit()
			return m, cmd
		}
		if key.Matches(msg, m.Keys.CursorUp) {
			m.Index += len(m.Items) - 1
			m.Index %= len(m.Items)
			return m, cmd
		}
		if key.Matches(msg, m.Keys.CursorDown) {
			m.Index++
			m.Index %= len(m.Items)
			return m, cmd
		}
		if key.Matches(msg, m.Keys.GoToEnd) {
			m.Index = len(m.Items) - 1
			return m, cmd
		}
		if key.Matches(msg, m.Keys.GoToStart) {
			m.Index = 0
			return m, cmd
		}
		if key.Matches(msg, m.Keys.Select) {
			m.Items[m.Index], cmd = m.Items[m.Index].OnSelect()
			return m, cmd
		}
	}
	if m.IsItemSelected {
		var model tea.Model
		model, cmd = m.Items[m.Index].Update(msg)
		m.Items[m.Index] = model.(Item)
	}
	return m, cmd
}

func (m Model) View() string {
	b := strings.Builder{}
	for i, item := range m.Items {
		b.WriteString(m.Styles.Render(item, m.Index == i) + "\n")
	}
	return b.String()
}

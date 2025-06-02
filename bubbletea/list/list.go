package list

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/term"
)

const buffer = 0

type Model struct {
	Index  int
	Items  []Item
	Keys   *KeyMap
	Styles Styles
	width  int
	height int
}

func New(models []tea.Model) Model {
	w, _, _ := term.GetSize(int(os.Stdout.Fd()))
	h := len(models) + buffer
	items := make([]Item, len(models))
	for i, model := range models {
		items[i] = NewItem(model)
	}
	m := Model{
		Items:  items,
		Keys:   *DefaultKeyMap(),
		width:  w,
		height: h,
	}
	m.SetFocus(0)
	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	selected := m.Items[m.Index]
	if selected.IsFocused {
		model, cmd, isNewModel := selected.Update(msg)
		if isNewModel {
			return model, cmd
		}
		m.Items[m.Index].Model = model
		return m, cmd
	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, m.Keys.CursorUp) {
			m.AdjustFocus(-1)
			return m, nil
		}
		if key.Matches(msg, m.Keys.CursorDown) {
			m.AdjustFocus(1)
			return m, nil
		}
		if key.Matches(msg, m.Keys.GoToEnd) {
			m.SetFocus(len(m.Items) - 1)
			return m, nil
		}
		if key.Matches(msg, m.Keys.GoToStart) {
			m.SetFocus(0)
			return m, nil
		}
		if key.Matches(msg, m.Keys.Cancel) {
			selected.Blur()
			selected.Cancel()
			m.Keys.SetItemFocus(false)
			return m, nil
		}
		if key.Matches(msg, m.Keys.Confirm) {
			selected.Blur()
			m.Keys.SetItemFocus(false)
			return m, nil
		}
		if key.Matches(msg, m.Keys.Select) {
			selected.Focus()
			m.Keys.SetItemFocus(true)
			return m, nil
		}
		if key.Matches(msg, m.Keys.ForceQuit) {
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) View() string {
	b := strings.Builder{}
	b.WriteString(fmt.Sprintf("List with %v items, Key Confirm Enabled: %v, Key Select Enabled: %v\n",
		len(m.Items), m.Keys.Confirm.Enabled(), m.Keys.Select.Enabled()))
	for i, item := range m.Items {
		if m.Index == i {
			b.WriteString(" ")
		}
		b.WriteString(m.Styles.Render(item) + "\n")
	}
	return b.String()
}

func (m *Model) AdjustFocus(increment int) {
	m.Items[m.Index].Blur()
	m.Index += increment + len(m.Items)
	m.Index %= len(m.Items)
	m.Items[m.Index].Focus()
}

func (m *Model) SetFocus(index int) {
	if index < 0 || index >= len(m.Items) {
		return
	}
	m.Items[m.Index].Blur()
	m.Index = index
	m.Items[m.Index].Focus()
}

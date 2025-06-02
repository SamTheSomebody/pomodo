package list

import (
	"reflect"

	tea "github.com/charmbracelet/bubbletea"
)

// Item fits the abstraction for text input but not for button
// Item needs focus, blur and select (and cancel?)

type Item struct {
	OldModel  tea.Model
	Model     tea.Model
	IsFocused bool
}

// This is dumb
func (m *Item) Update(msg tea.Msg) (tea.Model, tea.Cmd, bool) {
	model, cmd := m.Model.Update(msg)
	if reflect.TypeOf(model) != reflect.TypeOf(m.Model) {
		return model, cmd, true
	}
	return m.Model, cmd, false
}

func NewItem(m tea.Model) Item {
	return Item{
		Model:    m,
		OldModel: m,
	}
}

func (i *Item) SetFocus(b bool) {
	i.IsFocused = b
	if b {
		i.OldModel = i.Model
	}
}

func (i *Item) Focus() {
	i.IsFocused = true
	i.OldModel = i.Model
}

func (i *Item) Blur() {
	i.IsFocused = false
}

func (i *Item) Cancel() {
	i.Model = i.OldModel
}

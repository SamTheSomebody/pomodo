package pages

import (
	"context"
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"pomodo/bubbletea"
	"pomodo/bubbletea/button"
	"pomodo/bubbletea/list"
	"pomodo/helpers"
	"pomodo/internal/database"
)

type AllocateTime struct {
	List list.Model
}

func NewAllocateTimePage(keymap *bubbletea.Keymap) AllocateTime {
	m := AllocateTime{}
	t := textinput.New()
	t.Prompt = "Allocate Time: "
	t.Validate = func(s string) error {
		_, err := time.ParseDuration(s)
		return err
	}
	m.List = list.New([]list.Item{
		list.NewTextInput(t),
		button.New("Confirm", func() (tea.Model, tea.Cmd) {
			return NewHomePage(keymap), m.AllocateTime()
		}),
	}, keymap)
	return m
}

func (m AllocateTime) Init() tea.Cmd {
	return m.List.Init()
}

func (m AllocateTime) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	l, cmd := m.List.Update(msg)
	m.List = l.(list.Model)
	return m, cmd
}

func (m AllocateTime) View() string {
	return m.List.View()
}

func (m AllocateTime) AllocateTime() tea.Cmd {
	durationString := m.List.Items[0].(list.TextInputItem).Input.Value()
	allocatedTime, err := time.ParseDuration(durationString)
	if err != nil {
		return bubbletea.ErrCmd(err)
	}
	db := helpers.GetDBQueries()
	user, err := db.GetFirstUser(context.TODO())
	if err != nil {
		return bubbletea.ErrCmd(err)
	}
	params := database.SetUserAllocatedTimeParams{
		ID:                   user.ID,
		AllocatedTimeSeconds: int64(allocatedTime.Seconds()),
	}
	db.SetUserAllocatedTime(context.TODO(), params)
	return bubbletea.LogCmd(fmt.Sprintf("Updated allocated time (%v sec)", params.AllocatedTimeSeconds))
}

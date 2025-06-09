package pages

import (
	"context"
	"pomodo/bubbletea"
	"pomodo/bubbletea/button"
	"pomodo/bubbletea/list"
	"pomodo/bubbletea/slider"
	"pomodo/helpers"
	"pomodo/internal/database"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

type AddUserPage struct {
	List   tea.Model
	UserID uuid.UUID
}

func NewAddUserPage() tea.Model {
	m := AddUserPage{}
	input := textinput.New()
	input.Prompt = "Allocated time:          "
	input.Validate = helpers.ValidateDuration
	input.Placeholder = "e.g. 7h45m"
	input.Width = 50
	items := []list.Item{
		slider.New("Priority Weight:         ", 3),
		slider.New("Enthusiasm Weight:       ", 2),
		slider.New("Daily First task Weight: ", 4),
		slider.New("Due Date Daily Weight:   ", 2),
		list.NewTextInput(input),
		button.New("Confirm", m.Submit),
	}
	m.List = list.New(items)
	return m
}

func (m AddUserPage) Init() tea.Cmd {
	return nil
}

func (m AddUserPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}

func (m AddUserPage) View() string {
	return m.List.View()
}

func (m *AddUserPage) Submit() (tea.Model, tea.Cmd) {
	id := uuid.New()
	db := helpers.GetDBQueries()
	_, err := db.CreateUser(context.TODO(), id)
	if err != nil {
		return m, bubbletea.ErrCmd(err)
	}
	items := m.List.(list.Model).Items
	params := database.SetUserWeightsParams{
		ID:                 id,
		PriorityWeight:     float64(items[0].(slider.Model).Value / 10),
		EnthusiasmWeight:   float64(items[1].(slider.Model).Value / 10),
		DueDateDailyWeight: float64(items[2].(slider.Model).Value / 10),
		FirstTaskWeight:    float64(items[3].(slider.Model).Value / 10),
	}
	_, err = db.SetUserWeights(context.TODO(), params)
	if err != nil {
		return m, bubbletea.ErrCmd(err)
	}
	allocatedTime, err := time.ParseDuration(items[4].(list.TextInputItem).Input.Value())
	if err != nil {
		return m, bubbletea.ErrCmd(err)
	}
	allocatedTimeParams := database.SetUserAllocatedTimeParams{
		ID:                   id,
		AllocatedTimeSeconds: int64(allocatedTime.Seconds()),
	}
	_, err = db.SetUserAllocatedTime(context.TODO(), allocatedTimeParams)
	if err != nil {
		return m, bubbletea.ErrCmd(err)
	}
	return m, bubbletea.NewPageCmd(func() (tea.Model, tea.Cmd) { return NewHomePage(), nil })
}

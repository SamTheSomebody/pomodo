package pages

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

/*


Name |
*/

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type viewTasksModel struct {
	state *State
	table table.Model
}

func (m viewTasksModel) Init() tea.Cmd {
	return nil
}

func (m viewTasksModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m.state.ProcessUniversalKeys(msg)
}

func (m viewTasksModel) View() string {
	return ""
}

func InitialViewTasksModel(s *State) tea.Model {
	m := viewTasksModel{}
	return m
}

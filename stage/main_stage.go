package stage

import (
	"github.com/PandaPy/pginer/component"
	tea "github.com/charmbracelet/bubbletea"
)

type MainStage struct {
	selector *component.Selector
}

func (s MainStage) View() string {
	return s.selector.View()
}

func (s *MainStage) Update(msg tea.Msg) tea.Cmd {
	s.selector.Update(msg)
	if s.selector.Confirmed() {
		return tea.Cmd(func() tea.Msg {
			return InsertStageMsg{NewStage: NewInitProjectStage()}
		})
	}
	return nil
}

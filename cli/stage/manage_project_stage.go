package stage

import (
	"github.com/PandaPy/pginer/component/form"
	tea "github.com/charmbracelet/bubbletea"
)

type ManageProjectStage struct {
	selector form.Form
}

func (s ManageProjectStage) View() string {
	return s.selector.View()
}

func (s *ManageProjectStage) Update(msg tea.Msg) tea.Cmd {
	s.selector.Update(msg)

	if submitted, content := s.selector.Submit(); submitted {
		return tea.Cmd(func() tea.Msg {
			switch s.selector.GetContent() {
			case "创建项目":
				return ToggleStage{
					Content: content,
					Stage:   NewCreateProjectStage(),
				}
			case "管理项目":
				return ToggleStage{
					Content: content,
					Stage:   NewManageProjectStage(),
				}
			default:
				return nil
			}
		})
	}
	return nil
}

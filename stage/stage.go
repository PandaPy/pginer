package stage

import (
	"github.com/PandaPy/pginer/component"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

// Stage 接口定义
type Stage interface {
	View() string
	Update(msg tea.Msg) tea.Cmd
}

type InsertStageMsg struct {
	NewStage Stage
}

// 主入口
func NewMainStage() *MainStage {
	options := []component.ChoiceItem{
		{Title: "创建项目", Label: "Create New Project"},
		{Title: "管理项目", Label: "Manage Existing Project"},
	}
	return &MainStage{
		selector: component.NewSelector("请选择你需要进行的操作", options),
	}
}

// 初始化项目
func NewInitProjectStage() *InitProjectStage {
	return &InitProjectStage{
		input:   component.NewInput("请输入项目名称"),
		spinner: spinner.New(),
	}
}

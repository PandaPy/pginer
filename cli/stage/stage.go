package stage

import (
	"github.com/PandaPy/pginer/component/form"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

// Stage 接口定义
type Stage interface {
	View() string
	Update(msg tea.Msg) tea.Cmd
}

type ToggleStage struct {
	Content string
	Stage   Stage
}

type QuitCli struct {
}

type InsertStageMsg struct {
	NewStage Stage
}

// 主入口
func NewMainStage() *MainStage {
	options := []form.SelectorItem{
		{Title: "创建项目", Label: "Create Project"},
		{Title: "管理项目", Label: "Manage Project"},
	}
	return &MainStage{
		selector: form.NewSelector("请选择你需要进行的操作", options),
	}
}

// 创建项目
func NewCreateProjectStage() *CreateProjectStage {
	return &CreateProjectStage{
		input:   form.NewInput("请输入项目名称"),
		spinner: spinner.New(),
		mode:    "input",
	}
}

// 管理项目
func NewManageProjectStage() *ManageProjectStage {
	options := []form.SelectorItem{
		{Title: "新增 API 模块", Label: "Create API Group"},
		{Title: "同步数据库", Label: "Migration Database"},
		{Title: "上传代码至仓库", Label: "Push to Repository"},
	}
	return &ManageProjectStage{
		selector: form.NewSelector("请选择你需要进行的操作", options),
	}
}

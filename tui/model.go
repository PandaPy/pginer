package tui

import (
	"github.com/PandaPy/pginer/config"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	list      list.Model
	textInput textinput.Model
	state     string
}

// item 定义了列表项的结构体
type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

// InitialModel 初始化并返回 Model
func InitialModel() Model {
	items := createItems()
	delegate := createDelegate()
	l := createList(items, delegate)

	return Model{list: l}
}

// createItems 创建并返回列表项
func createItems() []list.Item {
	return []list.Item{
		item{title: "Init Project", desc: "初始化项目模板"},
		item{title: "Manager Project", desc: "管理项目"},
		item{title: "Commit Project", desc: "提交项目至代码仓库"},
		item{title: "Build Project", desc: "构建项目"},
	}
}

// createDelegate 创建并返回定制的项代理
func createDelegate() list.DefaultDelegate {
	delegate := list.NewDefaultDelegate()

	// 定义选中状态下的样式
	selectedStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderForeground(lipgloss.AdaptiveColor{Light: "#165DFF", Dark: "#165DFF"}).
		Foreground(lipgloss.AdaptiveColor{Light: "#165DFF", Dark: "#165DFF"}).
		Padding(0, 0, 0, 1)

	delegate.Styles.SelectedTitle = selectedStyle
	delegate.Styles.SelectedDesc = selectedStyle

	return delegate
}

// createList 创建并返回带有定制样式的列表
func createList(items []list.Item, delegate list.DefaultDelegate) list.Model {
	l := list.New(items, delegate, 0, 0)
	l.Title = config.ProjectName + " " + config.Version + " - " + config.Description

	l.Styles.Title = lipgloss.NewStyle()
	l.SetFilteringEnabled(false)
	l.SetShowStatusBar(false)

	// 禁用 "ShowFullHelp" 按键以移除 “more” 提示
	l.KeyMap.ShowFullHelp.SetEnabled(false)

	return l
}

// Init 初始化模型
func (m Model) Init() tea.Cmd {
	return nil
}

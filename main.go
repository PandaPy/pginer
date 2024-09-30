package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var choices = []string{"创建项目", "管理项目", "退出"}

type model struct {
	cursor   int
	selected string
	choices  []string
}

// 实现 Bubble Tea 的 `Init` 方法
func (m model) Init() tea.Cmd {
	return nil
}

// 更新模型状态
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		// 向上移动选择
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// 向下移动选择
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// 选择项目
		case "enter":
			m.selected = m.choices[m.cursor]
			return m, tea.Quit

		// 退出
		case "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

// 显示视图
func (m model) View() string {
	s := "请选择一个操作:\n\n"

	// 显示菜单选项
	for i, choice := range m.choices {
		cursor := " " // 没有选择项时的光标
		if m.cursor == i {
			cursor = ">" // 当前选择项的光标
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\n按 q 退出.\n"

	return lipgloss.NewStyle().Bold(true).Render(s)
}

// 创建项目的逻辑
func initProject() {
	fmt.Println("初始化项目...")
}

// 管理项目的逻辑
func manageProject() {
	fmt.Println("管理项目...")
}

func main() {
	// Root 命令
	var rootCmd = &cobra.Command{
		Use:   "pginer",
		Short: "pginer 是一个简单的命令行工具",
		Run: func(cmd *cobra.Command, args []string) {
			m := model{
				choices: choices,
			}
			p := tea.NewProgram(m) // 使用 tea.NewProgram 来启动程序

			// 启动 Bubble Tea 程序
			if err := p.Start(); err != nil {
				fmt.Printf("错误: %v\n", err)
				os.Exit(1)
			}

			switch m.selected {
			case "创建项目":
				initProject()
			case "管理项目":
				manageProject()
			default:
				fmt.Println("退出应用...")
				os.Exit(0)
			}
		},
	}

	// 定义 `createproject` 子命令
	var createProjectCmd = &cobra.Command{
		Use:   "createproject",
		Short: "初始化项目",
		Run: func(cmd *cobra.Command, args []string) {
			initProject()
		},
	}

	// 定义 `manage` 子命令
	var manageCmd = &cobra.Command{
		Use:   "manage",
		Short: "管理项目",
		Run: func(cmd *cobra.Command, args []string) {
			manageProject()
		},
	}

	// 添加子命令到 rootCmd
	rootCmd.AddCommand(createProjectCmd)
	rootCmd.AddCommand(manageCmd)

	// 执行 rootCmd
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

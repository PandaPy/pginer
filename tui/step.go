package tui

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	git "github.com/go-git/go-git/v5"
)

// Step 接口
type Step interface {
	View() string
	Update(msg tea.Msg, m *model) (Step, tea.Cmd)
}

// 主页面
type MainStep struct {
	cursor     int // 游标
	promptText string
	options    []string
	info       string
}

func (s MainStep) View() string {
	sv := renderPromptText(s.promptText, s.info, 1)
	sv += renderChoices(s.options, s.cursor)
	return sv
}

func (s *MainStep) Update(msg tea.Msg, m *model) (Step, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if s.cursor > 0 {
				s.cursor--
			}
		case "down", "j":
			if s.cursor < len(s.options)-1 {
				s.cursor++
			}
		case "enter":
			m.history = append(m.history, renderPromptText(s.promptText, s.options[s.cursor], 2))

			switch s.cursor {
			case 0:
				return InitProjectStep{
					promptText: "请输入项目名称",
					spinner:    spinner.New(),
				}, nil
			case 1:
				return InitProjectStep{
					promptText: "请选择你需要进行的操作",
					spinner:    spinner.New(),
				}, nil
			}

		}
	}
	return s, nil
}

// 初始化项目
type InitProjectStep struct {
	promptText string
	input      string
	infos      []string
	loading    bool
	spinner    spinner.Model
}

func (s InitProjectStep) View() string {
	sv := renderPromptText(s.promptText, s.input, 1)
	if s.loading {
		sv += fmt.Sprintf("%s 正在获取项目模板\n", loadingTextColor.Sprint(s.spinner.View()))
	} else {
		sv += strings.Join(s.infos, "\n")
	}
	return sv
}

func (s InitProjectStep) Update(msg tea.Msg, m *model) (Step, tea.Cmd) {
	validInput := regexp.MustCompile(`^[a-zA-Z0-9._-]$`)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "backspace":
			if len(s.input) > 0 {
				s.input = s.input[:len(s.input)-1]
			}
		case "enter":
			// 更新状态以显示加载信息
			s.loading = true
			s.spinner = spinner.New()
			s.spinner.Spinner = spinner.MiniDot
			return s, tea.Batch(
				func() tea.Msg {
					time.Sleep(10 * time.Second) // 模拟延迟
					// err := cloneTemplate(s.input, "https://github.com/PandaPy/pginer-template.git")
					// if err != nil {
					// 	return cloneErrorMsg{err}
					// }
					return cloneSuccessMsg{}
				},
				s.spinner.Tick,
			)
		default:
			if validInput.MatchString(msg.String()) {
				s.input += msg.String()
			}
		}
	case spinner.TickMsg:
		// 更新 spinner 动画
		var cmd tea.Cmd
		s.spinner, cmd = s.spinner.Update(msg)
		return s, cmd
	case cloneErrorMsg:
		// 错误信息显示，并退出程序
		s.loading = false
		s.infos = append(s.infos, errorTextColor.Sprint("Error cloning template: ", msg.err))
		return s, tea.Quit
	case cloneSuccessMsg:
		// 成功拉取模板后显示成功信息
		s.loading = false

		s.infos = append(s.infos, renderStatusText(s.spinner, "获取项目模板成功", 2))
		return s, tea.Quit
	}
	return s, nil
}

type cloneSuccessMsg struct{}
type cloneErrorMsg struct{ err error }

// cloneTemplate 函数：从 GitHub 仓库拉取模板到指定文件夹
func cloneTemplate(folderName, repoURL string) error {
	// 检查文件夹是否已存在
	if _, err := os.Stat(folderName); !os.IsNotExist(err) {
		return fmt.Errorf("directory %s already exists", folderName)
	}

	// 使用 go-git 克隆仓库
	_, err := git.PlainClone(folderName, false, &git.CloneOptions{
		URL:      repoURL,
		Progress: nil,
		Depth:    1, // 浅克隆，只拉取最新的提交
	})
	if err != nil {
		return fmt.Errorf("failed to clone repository: %w", err)
	}

	return nil
}

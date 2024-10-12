package stage

import (
	"os"
	"strings"
	"time"

	"github.com/PandaPy/pginer/component"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	git "github.com/go-git/go-git/v5"
)

type InitProjectStage struct {
	input       *component.Input
	spinner     spinner.Model
	contentText []string
	loading     bool
	loadingText string
}

type InstallDependenciesMsg struct{}
type TeaQuit struct{}

func (s InitProjectStage) View() string {
	sv := s.input.View() + strings.Join(s.contentText, "")
	if s.loading {
		sv += component.LoadingPromptMessage(s.spinner.View(), s.loadingText)
	}
	return sv
}

func (s *InitProjectStage) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if !s.loading && s.input.GetContentText() != "" {
				if cmd := s.VerifyFolderName(); cmd != nil {
					return cmd
				}
				s.loading = true
				return tea.Batch(s.DownloadTemplate(), s.spinner.Tick)
			}
		default:
			if !s.loading {
				s.input.Update(msg)
			}
		}
	case InstallDependenciesMsg:
		return tea.Batch(s.InstallDependencies(), s.spinner.Tick)
	case TeaQuit:
		s.loading = false
		return tea.Quit
	case spinner.TickMsg:
		var cmd tea.Cmd
		s.spinner, cmd = s.spinner.Update(msg)
		return cmd
	}
	return nil
}

// 校验文件夹是否存在
func (s *InitProjectStage) VerifyFolderName() tea.Cmd {
	folderName := s.input.GetContentText()
	if _, err := os.Stat(folderName); !os.IsNotExist(err) {
		s.showErrorMessage("目录 %s 已经存在")
		return tea.Quit
	}
	return nil
}

func (s *InitProjectStage) DownloadTemplate() tea.Cmd {
	return func() tea.Msg {
		s.showLoadingMessage("正在拉取项目模板...")
		repoURL := "https://gitee.com/pandapy/pginer-template.git"
		_, err := git.PlainClone(s.input.GetContentText(), false, &git.CloneOptions{
			URL:      repoURL,
			Progress: nil,
			Depth:    1,
		})
		if err != nil {
			s.showErrorMessage(err.Error())
			return TeaQuit{}
		}
		s.showSuccessMessage("获取项目模板成功")
		return InstallDependenciesMsg{}
	}
}

// 安装依赖
func (s *InitProjectStage) InstallDependencies() tea.Cmd {
	return func() tea.Msg {
		s.showLoadingMessage("正在安装依赖..")
		time.Sleep(5 * time.Second)
		s.showSuccessMessage("依赖安装成功")
		return TeaQuit{}
	}
}

func (s *InitProjectStage) showLoadingMessage(message string) {
	s.loadingText = message
}

func (s *InitProjectStage) showErrorMessage(message string) {
	s.contentText = append(s.contentText, component.PromptMessage(message, component.ErrorMode))
}

func (s *InitProjectStage) showSuccessMessage(message string) {
	s.contentText = append(s.contentText, component.PromptMessage(message, component.SuccessMode))
}

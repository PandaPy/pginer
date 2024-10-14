package stage

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/PandaPy/pginer/component"
	"github.com/PandaPy/pginer/component/form"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	git "github.com/go-git/go-git/v5"
)

type CreateProjectStage struct {
	input       form.Form
	spinner     spinner.Model
	contentText []string
	promptText  string
	mode        string
}

func (s CreateProjectStage) View() string {
	sv := s.input.View() + "\n" + strings.Join(s.contentText, "")
	if s.mode == "run" {
		sv += component.LoadingPromptMessage(s.spinner.View(), s.promptText)
	}
	return sv
}

func (s *CreateProjectStage) Update(msg tea.Msg) tea.Cmd {
	switch s.mode {
	case "input":
		s.input.Update(msg)
		if submitted, _ := s.input.Submit(); submitted {
			if cmd := s.VerifyFolderName(); cmd != nil {
				return cmd
			}
			return tea.Batch(s.DownloadTemplate(), s.spinner.Tick)
		}
	case "run":
		switch msg.(type) {
		case spinner.TickMsg:
			var cmd tea.Cmd
			s.spinner, cmd = s.spinner.Update(msg)
			return cmd
		}
	}
	return nil
}

// 校验文件夹是否存在
func (s *CreateProjectStage) VerifyFolderName() tea.Cmd {
	folderPath := s.input.GetContent()
	absPath, err := filepath.Abs(folderPath)
	if err != nil {
		s.showErrorMessage(fmt.Sprintf("无法解析路径 %s: %v", folderPath, err))
		return tea.Quit
	}

	if _, err := os.Stat(absPath); !os.IsNotExist(err) {
		s.showErrorMessage(fmt.Sprintf("路径已经存在: %s", absPath))
		return tea.Quit
	}
	return nil
}

func (s *CreateProjectStage) DownloadTemplate() tea.Cmd {
	folderPath := s.input.GetContent()
	return func() tea.Msg {
		s.mode = "run"
		s.showLoadingMessage("正在拉取项目模板...")
		repoURL := "https://gitee.com/pandapy/pginer-template.git"
		_, err := git.PlainClone(folderPath, false, &git.CloneOptions{
			URL:      repoURL,
			Progress: nil,
			Depth:    1,
		})
		if err != nil {
			s.showErrorMessage(fmt.Sprintf("获取项目模板失败: %s", err.Error()))
			return QuitCli{}
		}
		s.showSuccessMessage("获取项目模板成功")
		return QuitCli{}
	}
}

func (s *CreateProjectStage) showLoadingMessage(message string) {
	s.promptText = message
}

func (s *CreateProjectStage) showErrorMessage(message string) {
	s.mode = "finish"
	s.contentText = append(s.contentText, component.PromptMessage(message, component.ErrorMode))
}

func (s *CreateProjectStage) showSuccessMessage(message string) {
	s.mode = "finish"
	s.contentText = append(s.contentText, component.PromptMessage(message, component.SuccessMode))
}

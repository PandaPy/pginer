package stage

import (
	"fmt"
	"os"

	"github.com/PandaPy/pginer/component"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fatih/color"
	git "github.com/go-git/go-git/v5"
)

type InitProjectStage struct {
	input       *component.Input
	spinner     spinner.Model
	contentText string
	loading     bool
}

func (s InitProjectStage) View() string {
	sv := s.input.View() + s.contentText
	return sv
}

func (s *InitProjectStage) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if !s.loading {
				return tea.Batch(s.spinner.Tick, s.InitProject())
			}
		default:
			if !s.loading {
				s.input.Update(msg)
			}
		}
	case spinner.TickMsg:
		if s.loading {
			var cmd tea.Cmd
			s.contentText = component.Text(s.spinner.View()).Color(color.FgCyan).String() + " 正在拉取项目模板"
			s.spinner, cmd = s.spinner.Update(msg)
			return cmd
		} else {
			s.contentText = component.Text("✔").Color(color.FgGreen).String() + " 获取项目模板成功"
			return tea.Quit
		}
	}
	return nil
}

func (s *InitProjectStage) InitProject() tea.Cmd {
	return func() tea.Msg {
		s.loading = true

		s.CloneTemplate(s.input.GetContentText(), "https://github.com/PandaPy/pginer-template.git")
		s.contentText = component.Text("✔").Color(color.FgGreen).String() + " 获取项目模板成功"
		s.loading = false
		return nil
	}
}

func (s *InitProjectStage) CloneTemplate(folderName, repoURL string) error {
	if _, err := os.Stat(folderName); !os.IsNotExist(err) {
		return fmt.Errorf("directory %s already exists", folderName)
	}
	_, err := git.PlainClone(folderName, false, &git.CloneOptions{
		URL:      repoURL,
		Progress: nil,
		Depth:    1,
	})
	if err != nil {
		return fmt.Errorf("failed to clone repository: %w", err)
	}
	return nil
}

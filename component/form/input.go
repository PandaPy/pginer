package form

import (
	"regexp"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/PandaPy/pginer/component"
)

type Input struct {
	promptText  string
	ContentText string
	submitted   bool
}

// 初始化
func NewInput(promptText string) *Input {
	return &Input{
		promptText: promptText,
		submitted:  false,
	}
}

// 视图
func (s Input) View() string {
	return s.renderPromptText()
}

// 更新
func (s *Input) Update(msg tea.Msg) tea.Cmd {
	validInput := regexp.MustCompile(`^[a-zA-Z0-9._-]$`)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "backspace":
			if len(s.ContentText) > 0 {
				s.ContentText = s.ContentText[:len(s.ContentText)-1]
			}
		case "enter":
			s.submitted = true
		default:
			if validInput.MatchString(msg.String()) {
				s.ContentText += msg.String()
			}
		}
	}
	return nil
}

// 提交
func (s *Input) Submit() (bool, string) {
	var promptText = s.renderPromptText()
	return s.submitted, promptText
}

// 渲染提示
func (s *Input) renderPromptText() string {
	questionMark := component.FgGreen.Bold(true).Render("? ")
	prompt := component.FgWhite.Bold(true).Render(s.promptText)
	result := component.FgCyan.Render(s.ContentText)
	return questionMark + prompt + " " + result
}

// 获取输入内容
func (s *Input) GetContent() string {
	return s.ContentText
}

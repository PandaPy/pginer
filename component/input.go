package component

import (
	"regexp"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fatih/color"
)

type Input struct {
	promptText  string
	contentText string
}

func NewInput(promptText string) *Input {
	return &Input{
		promptText: promptText,
	}
}

func (s *Input) Update(keyMsg tea.KeyMsg) tea.Cmd {
	validInput := regexp.MustCompile(`^[a-zA-Z0-9._-]$`)
	switch keyMsg.String() {
	case "backspace":
		if len(s.contentText) > 0 {
			s.contentText = s.contentText[:len(s.contentText)-1]
		}
	default:
		if validInput.MatchString(keyMsg.String()) {
			s.contentText += keyMsg.String()
		}
	}
	return nil
}

// 前缀显示
func (s Input) Prefix() string {
	return Text("? ").Color(color.FgGreen).Bold().String()
}

// 提示文本
func (s Input) PromptText() string {
	return Text(s.promptText).Color(color.FgWhite).Bold().String()
}

// 后缀显示
func (s Input) Suffix() string {
	return Text(s.contentText).Color(color.FgHiBlack).String()
}

// 获取输入内容
func (s Input) GetContentText() string {
	return s.contentText
}

// 视图
func (s Input) View() string {
	return s.Prefix() + s.PromptText() + " " + s.Suffix() + "\n"
}

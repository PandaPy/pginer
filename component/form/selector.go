package form

import (
	"github.com/PandaPy/pginer/component"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Selector struct {
	focus      int
	promptText string
	options    []SelectorItem
	submitted  bool
}

type SelectorItem struct {
	Title string
	Label string
}

// 初始化
func NewSelector(promptText string, options []SelectorItem) *Selector {
	return &Selector{
		focus:      0,
		promptText: promptText,
		options:    options,
		submitted:  false,
	}
}

// 视图
func (s Selector) View() string {
	return s.renderPromptText("(Use arrow keys)") + "\n" + s.renderOptions()
}

// 更新
func (s *Selector) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if s.focus > 0 {
				s.focus--
			}
		case "down", "j":
			if s.focus < len(s.options)-1 {
				s.focus++
			}
		case "enter":
			s.submitted = true
		}
	}
	return nil
}

// 提交
func (s *Selector) Submit() (bool, string) {
	var promptText = s.renderPromptText(s.options[s.focus].Title)
	return s.submitted, promptText
}

// 获取结果
func (s *Selector) GetContent() string {
	return s.options[s.focus].Title
}

// 渲染提示
func (s *Selector) renderPromptText(resultText string) string {
	questionMark := component.FgGreen.Bold(true).Render("? ")
	prompt := component.FgWhite.Bold(true).Render(s.promptText)
	result := component.FgHiBlack.Render(resultText)
	if s.submitted {
		result = component.FgCyan.Render(resultText)
	}
	return questionMark + prompt + " " + result
}

// 渲染选项
func (s *Selector) renderOptions() string {
	var lines []string
	for i, choice := range s.options {
		focusSymbol := " "
		titleColor := component.FgWhite
		labelColor := component.FgHiBlack

		if i == s.focus {
			focusSymbol = "❯"
			titleColor = component.FgCyan
			labelColor = component.FgCyan
		}

		line := titleColor.Render(focusSymbol+" "+choice.Title+" - ") + labelColor.Render(choice.Label)
		lines = append(lines, line)
	}

	return lipgloss.JoinVertical(lipgloss.Left, lines...)
}

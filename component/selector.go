package component

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fatih/color"
)

type ChoiceItem struct {
	Title string
	Label string
}

type Selector struct {
	cursor     int
	promptText string
	options    []ChoiceItem
	confirmed  bool
}

func NewSelector(promptText string, options []ChoiceItem) *Selector {
	return &Selector{
		cursor:     0,
		promptText: promptText,
		options:    options,
		confirmed:  false,
	}
}

func (s *Selector) Update(msg tea.Msg) tea.Cmd {
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
			// 确认选择
			s.confirmed = true
		}
	}
	return nil
}

func (s Selector) View() string {
	var resultText string
	if s.confirmed {
		// 如果确认选择，只显示已选项的标题
		resultText = s.options[s.cursor].Title
		return s.RenderPromptText(s.promptText, resultText, true)
	} else {
		// 未确认时显示提示文本和所有选项
		resultText = "(Use arrow keys)"
		return s.RenderPromptText(s.promptText, resultText, false) + s.RenderOptions()
	}
}

func (s Selector) RenderPromptText(promptText, resultText string, finish bool) string {
	questionMark := Text("? ").Color(color.FgGreen).Bold().String()
	prompt := Text(promptText).Color(color.FgWhite).Bold().String()
	result := Text(resultText).Color(getColor(finish)).String()
	return questionMark + prompt + " " + result + "\n"
}

func (s Selector) RenderOptions() string {
	output := ""
	for i, choice := range s.options {
		cursorSymbol := " "
		titleColor := color.FgWhite
		labelColor := color.FgHiBlack

		if i == s.cursor {
			cursorSymbol = "❯"
			titleColor = color.FgCyan
		}

		output += Text(cursorSymbol + " " + choice.Title + " - ").Color(titleColor).String()
		output += Text(choice.Label + "\n").Color(labelColor).String()
	}
	return output
}

func (s Selector) Confirmed() bool {
	return s.confirmed
}

func getColor(finish bool) color.Attribute {
	if finish {
		return color.FgCyan
	}
	return color.FgHiBlack
}

func (s *Selector) Reset() {
	s.cursor = 0
	s.confirmed = false
}

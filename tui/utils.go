package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/fatih/color"
)

var (
	titleColor           = color.New(color.FgBlue).Add(color.Bold)
	questionMarkColor    = color.New(color.FgGreen)
	promptTextColor      = color.New(color.FgWhite).Add(color.Bold)
	optionTextColor      = color.New(color.FgWhite)
	optionTextLightColor = color.New(color.FgCyan).Add(color.Bold)
	successTextColor     = color.New(color.FgGreen)
	errorTextColor       = color.New(color.FgRed)
	loadingTextColor     = color.New(color.FgCyan)
)

// 渲染提示文本
func renderPromptText(promptText, info string, showMode int) string {
	var infoColor *color.Color
	switch showMode {
	case 1:
		infoColor = color.New(color.FgHiBlack)
	case 2:
		infoColor = color.New(color.FgCyan)
	default:
		infoColor = color.New(color.Reset)
	}

	return questionMarkColor.Sprint("? ") + promptTextColor.Sprint(promptText) + infoColor.Sprintln(" "+info)
}

// 渲染状态文本
func renderStatusText(sp spinner.Model, info string, showMode int) string {
	switch showMode {
	case 1:
		return fmt.Sprintf("%s %s\n", loadingTextColor.Sprint(sp.View()), info)
	case 2:
		return fmt.Sprintf("%s %s\n", successTextColor.Sprint("✔"), info)
	default:
		return ""
	}
}

// 渲染选择项列表
func renderChoices(choices []string, cursor int) string {
	var result string
	for i, choice := range choices {
		cursorSymbol := " "
		if cursor == i {
			cursorSymbol = "❯"
			result += optionTextLightColor.Sprintf("%s %s\n", cursorSymbol, choice)
		} else {
			result += optionTextColor.Sprintf("%s %s\n", cursorSymbol, choice)
		}
	}
	return result
}

package component

import "github.com/fatih/color"

type PromptMessageMode int

const (
	SuccessMode PromptMessageMode = iota // 成功
	ErrorMode                            // 错误
	DefaultMode                          // 默认
)

func PromptMessage(contentText string, mode PromptMessageMode) string {
	switch mode {
	case SuccessMode:
		return Text("✔ ").Color(color.FgGreen).String() + Text(contentText+"\n").Color(color.FgWhite).String()
	case ErrorMode:
		return Text("✖ ").Color(color.FgRed).String() + Text(contentText+"\n").Color(color.FgWhite).String()
	case DefaultMode:
		return Text(contentText + "\n").Color(color.FgWhite).String()
	}
	return contentText
}

func LoadingPromptMessage(spinner string, contentText string) string {
	return Text(spinner+" ").Color(color.FgGreen).String() + Text(contentText+"\n").Color(color.FgWhite).String()
}

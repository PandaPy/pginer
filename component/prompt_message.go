package component

type PromptMessageMode int

const (
	SuccessMode PromptMessageMode = iota // 成功
	ErrorMode                            // 错误
	DefaultMode                          // 默认
)

func PromptMessage(contentText string, mode PromptMessageMode) string {
	switch mode {
	case SuccessMode:
		return FgGreen.Render("✔ ") + FgWhite.Render(contentText+"\n")
	case ErrorMode:
		return FgRed.Render("✘ ") + FgWhite.Render(contentText+"\n")
	case DefaultMode:
		return FgWhite.Render(contentText + "\n")
	}
	return contentText
}

func LoadingPromptMessage(spinner string, contentText string) string {
	return FgGreen.Render(spinner+" ") + FgWhite.Render(contentText+"\n")
}

package form

import tea "github.com/charmbracelet/bubbletea"

type Form interface {
	View() string
	Update(msg tea.Msg) tea.Cmd
	Submit() (bool, string)
	GetContent() string
}

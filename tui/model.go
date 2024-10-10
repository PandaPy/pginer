package tui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	step    Step
	history []string
}

func InitialModel() *model {
	return &model{
		step: &MainStep{
			promptText: "请选择你需要进行的操作",
			options:    []string{"初始化项目", "管理项目"},
			info:       "(Use arrow keys)",
		},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	newStep, cmd := m.step.Update(msg, m)
	m.step = newStep

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, cmd
}

func (m model) View() string {
	s := titleColor.Sprintln(asciiArt)
	s += strings.Join(m.history, "\n")
	s += m.step.View()
	s += "\n\n\n"
	return s
}

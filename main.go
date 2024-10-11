package main

import (
	"fmt"
	"os"

	"github.com/PandaPy/pginer/component"
	"github.com/PandaPy/pginer/stage"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	stages []stage.Stage
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		default:
			cmd := m.stages[len(m.stages)-1].Update(msg)
			return m, cmd
		}
	case stage.InsertStageMsg:
		m.stages = append(m.stages, msg.NewStage)
	default:
		cmd := m.stages[len(m.stages)-1].Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m model) View() string {
	s := component.Banner()
	// 获取当前阶段并显示
	for _, stage := range m.stages {
		s += stage.View()
	}
	return s + "\n\n"
}

func main() {
	m := model{
		stages: []stage.Stage{
			stage.NewMainStage(),
		},
	}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("could not run program:", err)
		os.Exit(1)
	}
}

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

func InitialModel() model {
	return model{
		stages: []stage.Stage{
			stage.NewMainStage(),
		},
	}
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
	case tea.QuitMsg:
		return m, tea.Quit
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
	if _, err := tea.NewProgram(InitialModel()).Run(); err != nil {
		fmt.Println("could not run program:", err)
		os.Exit(1)
	}
}

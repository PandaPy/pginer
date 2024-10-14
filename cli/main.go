package cli

import (
	"fmt"
	"os"

	"github.com/PandaPy/pginer/cli/stage"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	version string
	infos   []string
	stage   stage.Stage
}

func (m model) Init() tea.Cmd {
	return nil
}

func InitialModel(version string) model {
	return model{
		version: version,
		infos:   []string{},
		stage:   stage.NewMainStage(),
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		default:
			cmd := m.stage.Update(msg)
			return m, cmd
		}
	case stage.ToggleStage:
		m.stage = msg.Stage
		m.infos = append(m.infos, msg.Content)
	case stage.QuitCli:
		return m, tea.Quit
	default:
		cmd := m.stage.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m model) View() string {
	lines := append([]string{NewVersionText(m.version)}, m.infos...)
	lines = append(lines, m.stage.View())
	return lipgloss.JoinVertical(lipgloss.Left, lines...) + "\n"
}

func Main(version string) {
	if _, err := tea.NewProgram(InitialModel(version)).Run(); err != nil {
		fmt.Println("could not run program:", err)
		os.Exit(1)
	}
}

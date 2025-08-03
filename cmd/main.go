package main

import (
	"fmt"
	"os"

	"github.com/Soheil7799/go-server-tools/internal/commands"
	"github.com/Soheil7799/go-server-tools/internal/config"
	ui "github.com/Soheil7799/go-server-tools/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
)

type SelectionMessage struct {
	Choice string
}

type model struct {
	screen     int
	menuModel  ui.MenuModel
	sshModel   ui.SshModel
	rsyncModel ui.RsyncModel
}

func initializeModel() model {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to load config file: %v", err))
	}
	return model{
		screen:     0,
		menuModel:  ui.NewMenuModel(),
		sshModel:   ui.NewSshModel(cfg),
		rsyncModel: ui.NewRsyncModel(cfg),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case ui.SelectionMessage:
		switch msg.Choice {
		case "SSH":
			m.screen = 1
		case "rsync":
			m.screen = 2
		case "Exit":
			return m, tea.Quit
		}
	case ui.SSHReadyMsg:
		err := commands.ExecuteSSH(msg.Server, msg.Key)
		if err != nil {
			fmt.Print(err)
			return m, nil
		}
		return m, tea.Quit
	case ui.RsyncReadyMsg:
		err := commands.ExecuteRsync(msg.Server.Host, msg.Key.Path, msg.LocalPath, msg.RemotePath, msg.Direction)
		if err != nil {
			fmt.Print(err)
			return m, nil
		}
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		switch m.screen {
		case 0:
			updatedModel, cmd := m.menuModel.Update(msg)
			m.menuModel = updatedModel.(ui.MenuModel)
			return m, cmd
		case 1:
			updatedModel, cmd := m.sshModel.Update(msg)
			m.sshModel = updatedModel.(ui.SshModel)
			return m, cmd
		case 2:
			updatedModel, cmd := m.rsyncModel.Update(msg)
			m.rsyncModel = updatedModel.(ui.RsyncModel)
			return m, cmd
		}

	}
	return m, nil
}
func (m model) View() string {
	switch m.screen {
	case 0:
		return m.menuModel.View()
	case 1:
		return m.sshModel.View()
	case 2:
		return m.rsyncModel.View()
	default:
		return "Unknown screen"
	}
}
func main() {
	m := initializeModel()
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

}

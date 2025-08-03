package main

import (
	"fmt"
	"os"

	ui "github.com/Soheil7799/go-server-tools/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
)

type SelectionMessage struct {
	Choice string
}

type model struct {
	screen    int
	menuModel ui.MenuModel
	sshMode   ui.SshModel
}

func initializeModel() model {
	return model{
		screen:    0,
		menuModel: ui.NewMenuModel(),
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

		}

	}
	return m, nil
}
func (m model) View() string {
	switch m.screen {
	case 0:
		return m.menuModel.View()
	case 1:
		return "SSH Screen (coming soon)\nPress q to quit"
	case 2:
		return "rsync Screen (coming soon)\nPress q to quit"
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

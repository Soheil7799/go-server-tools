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
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return nil, nil }
func (m model) View() string                            { return "" }

func main() {
	m := initializeModel()
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

}

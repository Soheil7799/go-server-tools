package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	message string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {

			return m, tea.Quit
		}

	}
	return m, nil
}
func (m model) View() string {
	return "Welcome to Go Server Tools!\n\nPress 'q' to quit.\n"
}

func main() {
	m := model{
		message: "Starting up...",
	}
	p := tea.NewProgram(m)
	_, err := p.Run()
	if err != nil {
		fmt.Printf("Encountred Error : %s", err)
		os.Exit(1)

	}
}

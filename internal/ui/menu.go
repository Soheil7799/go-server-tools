package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices  []string
	cursor   int
	selected string
}

func initializeModel() model {
	return model{
		choices: []string{"SSH", "rsync", "Exit"},
		cursor:  0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "enter", " ":
			m.selected = m.choices[m.cursor]
			return m, tea.Quit
		case "ctrl+c", "q":
			return m, tea.Quit

		}
	}
	return m, nil
}
func (m model) View() string {
	s := "which command do you want to run ?\n\n"
	for i, c := range m.choices {
		selectedCursor := "[ ]"

		if m.cursor == i {
			selectedCursor = "[*]"
		}
		s += fmt.Sprintf("%s %s\n", selectedCursor, c)
	}
	s += "\n\nYou can quit with ctrl+c or typing \"q\""
	return s
}

func main() {
	m := initializeModel()
	p := tea.NewProgram(m)
	_, err := p.Run()
	if err != nil {
		fmt.Printf("Encountred Error : %s", err)
		os.Exit(1)

	}
}

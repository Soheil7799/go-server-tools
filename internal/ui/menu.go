package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type MenuModel struct {
	choices  []string
	cursor   int
	selected string
}
type SelectionMessage struct {
	Choice string
}

func NewMenuModel() MenuModel {
	return MenuModel{
		choices: []string{"SSH", "rsync", "Exit"},
		cursor:  0,
	}
}

func (m MenuModel) Init() tea.Cmd {
	return nil
}

func (m MenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			return m, func() tea.Msg {
				return SelectionMessage{Choice: m.selected}
			}
		case "ctrl+c", "q":
			return m, tea.Quit

		}
	}
	return m, nil
}
func (m MenuModel) View() string {
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

// func main() {
// 	m := NewMenuModel()
// 	p := tea.NewProgram(m)
// 	_, err := p.Run()
// 	if err != nil {
// 		fmt.Printf("Encountred Error : %s", err)

// 	}
// }

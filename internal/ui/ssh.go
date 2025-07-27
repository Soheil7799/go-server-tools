package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type SshModel struct {
	Keys           []string
	SelectedKey    string
	Servers        []string
	SelectedServer string
	Cursor         int
	Step           int
}

func NewSshModel() SshModel {
	return SshModel{
		Cursor:  0,
		Keys:    []string{"rsa_key_1", "rsa_key_2", "rsa_key_3"},
		Servers: []string{"backend_1", "frontend_1", "logic"},
		Step:    0,
	}
}

type SSHReadyMsg struct {
	Server string
	Key    string
}

func (m SshModel) Init() tea.Cmd {
	return nil
}

func (m SshModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "down", "j":
			var leng int
			if m.Step == 0 {
				leng = len(m.Servers)
			} else {
				leng = len(m.Keys)
			}
			if m.Cursor < leng-1 {
				m.Cursor++
			}
		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "enter", " ":
			switch m.Step {
			case 0:
				m.SelectedServer = m.Servers[m.Cursor]
				m.Step = 1
				m.Cursor = 0
			case 1:
				m.SelectedKey = m.Keys[m.Cursor]
				return m, func() tea.Msg {
					return SSHReadyMsg{
						Server: m.SelectedServer,
						Key:    m.SelectedKey,
					}
				}
			}
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m SshModel) View() string {
	var s string
	if m.SelectedServer == "" {
		s = "Choose the server you want to connect:\n\n"
		for i, r := range m.Servers {
			cursor := "[ ]"
			if m.Cursor == i {
				cursor = "[*]"
			}
			s += fmt.Sprintf("%s %s\n", cursor, r)
		}
	} else if m.SelectedServer != "" {
		s = "Choose the key you want to use:\n\n"
		for i, k := range m.Keys {
			cursor := "[ ]"
			if m.Cursor == i {
				cursor = "[*]"
			}
			s += fmt.Sprintf("%s %s\n", cursor, k)
		}
	}
	return s
}

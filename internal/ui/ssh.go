package ui

import (
	"fmt"

	"github.com/Soheil7799/go-server-tools/internal/config"
	tea "github.com/charmbracelet/bubbletea"
)

type SshModel struct {
	Keys           []config.SSHKey
	SelectedKey    config.SSHKey
	Servers        []config.Server
	SelectedServer config.Server
	Cursor         int
	Step           int
}

func NewSshModel(cfg *config.Config) SshModel {
	return SshModel{
		Cursor:  0,
		Keys:    cfg.SSHKeys,
		Servers: cfg.Servers,
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
						Server: m.SelectedServer.Host,
						Key:    m.SelectedKey.Path,
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
	if m.SelectedServer.Name == "" {
		s = "Choose the server you want to connect:\n\n"
		for i, r := range m.Servers {
			cursor := "[ ]"
			if m.Cursor == i {
				cursor = "[*]"
			}
			s += fmt.Sprintf("%s %s - %s\n", cursor, r.Name, r.Description)
		}
	} else if m.SelectedServer.Name != "" {
		s = "Choose the key you want to use:\n\n"
		for i, k := range m.Keys {
			cursor := "[ ]"
			if m.Cursor == i {
				cursor = "[*]"
			}
			s += fmt.Sprintf("%s %s - %s\n", cursor, k.Name, k.Description)
		}
	}
	return s
}

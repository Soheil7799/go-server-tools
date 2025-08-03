package ui

import (
	"fmt"

	"github.com/Soheil7799/go-server-tools/internal/config"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type RsyncModel struct {
	Keys           []config.SSHKey
	SelectedKey    config.SSHKey
	Servers        []config.Server
	SelectedServer config.Server
	LocalPath      textinput.Model
	RemotePath     textinput.Model
	Cursor         int
	Direction      int // 0 = to server  1 = from server
	Step           int // 0=server_selection 1=key_selection 2=to/from 3=localpath 4=remotepath
}

func NewRsyncModel(cfg *config.Config) RsyncModel {
	localInput := textinput.New()
	localInput.Placeholder = "/path/to/local/file(s)"
	remoteInput := textinput.New()
	remoteInput.Placeholder = "/path/to/remote/file(s)"

	return RsyncModel{
		Cursor:     0,
		Keys:       cfg.SSHKeys,
		Servers:    cfg.Servers,
		Step:       0,
		LocalPath:  localInput,
		RemotePath: remoteInput,
	}
}

type RsyncReadyMsg struct {
	Server     config.Server
	Key        config.SSHKey
	LocalPath  string
	RemotePath string
	Direction  int
}

func (m RsyncModel) Init() tea.Cmd {
	return nil
}

func (m RsyncModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	if m.Step == 3 {
		m.LocalPath, cmd = m.LocalPath.Update(msg)
	} else if m.Step == 4 {
		m.RemotePath, cmd = m.RemotePath.Update(msg)
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "down", "j":
			var leng int
			switch m.Step {
			case 0:
				leng = len(m.Servers)
			case 1:
				leng = len(m.Keys)
			case 2:
				leng = 2
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
				m.Cursor = 0
				m.Step = 2
			case 2:
				m.Direction = m.Cursor
				m.Cursor = 0
				if m.Direction == 0 {
					// to Server
					// get the local then copy to server
					m.Step = 3
					m.LocalPath.Focus()
				} else {
					// to Local
					// get the server then copy to local
					m.Step = 4
					m.RemotePath.Focus()
				}
			case 3:
				if m.Direction == 0 {
					// to server -> should get server next
					m.Step = 4
					m.RemotePath.Focus()
				} else {
					// to local -> server already taken
					m.Step = 5
				}
			case 4:
				if m.Direction == 0 {
					// to server -> local already taken
					m.Step = 5
				} else {
					// to local -> should get local
					m.Step = 3
					m.LocalPath.Focus()
				}
			case 5:
				// all pathes already taken, return the ready message
				return m, func() tea.Msg {
					return RsyncReadyMsg{
						Server:     m.SelectedServer,
						Key:        m.SelectedKey,
						LocalPath:  m.LocalPath.Value(),
						RemotePath: m.RemotePath.Value(),
						Direction:  m.Direction,
					}
				}
			}
		case "ctrl+c", "q":
			return m, tea.Quit
		}

	}
	return m, cmd
}

func (m RsyncModel) View() string {
	var s string
	switch m.Step {
	case 0: // server selection
		s = "Choose the server you want to connect:\n\n"
		for i, r := range m.Servers {
			cursor := "[ ]"
			if m.Cursor == i {
				cursor = "[*]"
			}
			s += fmt.Sprintf("%s %s - %s\n", cursor, r.Name, r.Description)
		}
	case 1: // key selection
		s = "Choose the key you want to use:\n\n"
		for i, k := range m.Keys {
			cursor := "[ ]"
			if m.Cursor == i {
				cursor = "[*]"
			}
			s += fmt.Sprintf("%s %s - %s\n", cursor, k.Name, k.Description)
		}
	case 2: // Direction selection
		s = "Choose the direction you want to transfer\n\n"
		dir := []string{"Copy \"TO\" server", "Copy \"FROM\" server"}
		for i, d := range dir {
			cursor := "[ ]"
			if m.Cursor == i {
				cursor = "[*]"
			}
			s += fmt.Sprintf("%s %s\n", cursor, d)
		}
	case 3:
		s = fmt.Sprintf("Enter local path:\n\n%s", m.LocalPath.View())
	case 4:
		s = fmt.Sprintf("Enter remote path:\n\n%s", m.RemotePath.View())

	}
	return s
}

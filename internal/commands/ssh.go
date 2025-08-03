package commands

import (
	"os"
	"os/exec"
)

func ExecuteSSH(server, key string) error {
	cmd := exec.Command("ssh", "-i", key, server)

	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

package commands

import (
	"fmt"
	"os"
	"os/exec"
)

func ExecuteRsync(server, key, localPath, remotePath string, direction int) error {
	remoteServerPath := fmt.Sprintf("%s:%s", server, remotePath)
	remoteServerSSH := fmt.Sprintf("\"ssh -i %s\"", key)
	var cmd *exec.Cmd
	if direction == 0 {
		cmd = exec.Command("rsync", "-az", "-e", remoteServerSSH, localPath, remoteServerPath)
	} else {
		cmd = exec.Command("rsync", "-az", "-e", remoteServerSSH, remoteServerPath, localPath)
	}

	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

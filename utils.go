package pam

import (
	"fmt"
	"os/exec"
)

func execCommand(cmd *exec.Cmd) error {

	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("args : %s \tstderr : %s \tstdout: %s", cmd.Args, err.Error(), string(output))
	}

	return nil
}

package mcvm

import (
	"os/exec"
	"strings"
)

// Reads MCVM command output as an array of strings
func ReadMCVMOutput(args ...string) []string {
	cmd, err := exec.Command("mcvm", args...).Output()
	if err != nil {
		panic(err)
	}

	return strings.Split(strings.TrimSuffix(string(cmd), "\n"), "\n")
}

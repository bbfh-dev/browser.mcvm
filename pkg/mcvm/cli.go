package mcvm

import (
	"os/exec"
	"strings"
)

// Reads MCVM command output as an array of strings
func ExecMCVMRaw(args ...string) (string, error) {
	cmd, err := exec.Command("mcvm", args...).Output()
	if err != nil {
		return "", err
	}

	return string(cmd), nil
}

// Reads MCVM command output as an array of strings
func ExecMCVM(args ...string) ([]string, error) {
	cmd, err := ExecMCVMRaw(args...)
	if err != nil {
		return nil, nil
	}

	return strings.Split(strings.TrimSuffix(cmd, "\n"), "\n"), nil
}

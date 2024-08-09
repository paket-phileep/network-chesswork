package utilities

import (
	"fmt"
	"os/exec"
	"strings"
)

func RunCommand(command string) (bool, error) {
	// Split the command into the command name and arguments
	parts := strings.Fields(command)
	if len(parts) == 0 {
		return false, fmt.Errorf("no command provided")
	}

	// The first part is the command, the rest are the arguments
	cmdName := parts[0]
	cmdArgs := parts[1:]

	// Define the command and arguments
	cmd := exec.Command(cmdName, cmdArgs...)

	// Execute the command and capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error executing command: %s\n", err)
		return false, err
	}

	// Print the output
	fmt.Printf("Command output:\n%s", output)
	return true, nil
}

package utilities

import (
	"fmt"
	"os"
	"os/exec"
)

// / RunScript executes a bash script located at the given path and returns whether it was successful and any error encountered.
func RunScript(scriptPath string, props map[string]string) (bool, error) {
	// Define the command to execute the script

	err := FilePathValidity(scriptPath)
	if err != nil {
		fmt.Printf("problem with scriptPath with path: %s\n", err)
		return false, err
	}

	cmd := exec.Command("bash", scriptPath)

	// Set environment variables if any
	envVars := os.Environ()
	for key, value := range props {
		envVars = append(envVars, fmt.Sprintf("%s=%s", key, value))
	}
	cmd.Env = envVars

	// Capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error executing script: %s\n", err)
		return false, err
	}

	// Print the output
	fmt.Printf("Script output:\n%s", output)

	return true, nil
}

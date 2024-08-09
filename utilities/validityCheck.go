package utilities

import (
	"fmt"
	"os"
)

func FilePathValidity(path string) error {
	// Check if the script path is valid
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return fmt.Errorf("script file does not exist at path: %s", path)
	}
	if err != nil {
		return fmt.Errorf("error checking script file: %w", err)
	}
	if !info.Mode().IsRegular() {
		return fmt.Errorf("path is not a regular file: %s", path)
	}

	return nil

}

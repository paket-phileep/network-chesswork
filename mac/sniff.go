package mac

import (
	"fmt"
	"network-chesswork/utilities"
)

// Sniff performs a sniffing operation based on the provided range.
func Sniff(interfaceRange string) error {
	if interfaceRange == "" {
		return fmt.Errorf("invalid range provided")
	}

	// Define the path to your script
	scriptPath := "./scripts/sniff-network.bash"

	// Define environment variables to pass to the script
	props := map[string]string{
		"range": interfaceRange,
	}

	// Run the script with the environment variables
	success, err := utilities.RunScript(scriptPath, props)

	if !success {
		fmt.Printf("Script execution failed: %v\n", err)
	} else {
		fmt.Println("Script executed successfully")
	}

	// Assuming utils.RunScript takes a string and an interface{}
	if err != nil {
		return err
	}

	return nil
}

package cronJob

import (
	"fmt"
	"network-chesswork/mac"
)

func FindConnectedMacAddresses(interfaceRange string) error {

	err := mac.Sniff(interfaceRange)

	// Assuming utils.RunScript takes a string and an interface{}
	if err != nil {
		return err
	} else {
		fmt.Println("Script executed successfully")
	}
	return nil
}

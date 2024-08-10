package cronJob

import (
	"fmt"
	"network-chesswork/mac"
	"time"
)

// FindConnectedMacAddresses starts a process to find connected MAC addresses using the specified network interface.
// If the interface is empty, it defaults to "en0". The process runs for up to 5 minutes.
func FindConnectedMacAddresses(iface string) error {
	if iface == "" {
		fmt.Println("No INTERFACE provided, using existing en0.")
		iface = "en0"
	}

	// Start a timer for 5 minutes
	timer := time.NewTimer(5 * time.Minute)
	defer timer.Stop()

	// Use a channel to signal completion
	done := make(chan error, 1)

	go func() {
		// Run the MAC address finding process and signal completion
		done <- findConnectedMacAddresses(iface, timer.C)
	}()

	// Wait for either the completion of the process or the timer expiration
	select {
	case err := <-done:
		return err
	case <-timer.C:
		fmt.Println("Operation timed out")
		return nil
	}
}

// findConnectedMacAddresses continuously finds connected MAC addresses
// and stops when the timeout channel signals.
func findConnectedMacAddresses(iface string, timeoutChan <-chan time.Time) error {
	ticker := time.NewTicker(10 * time.Second) // Adjust the interval as needed
	defer ticker.Stop()

	for {
		select {
		case <-timeoutChan:
			return nil
		case <-ticker.C:
			// Simulate work
			err := mac.Sniff(iface)
			if err != nil {
				return err
			}
		}
	}
}

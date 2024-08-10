package cronJob

import (
	"context"
	"fmt"
	"log"
	"network-chesswork/mac"
	"time"
)

func FindConnectedMacAddresses(iface string) error {
	if iface == "" {
		iface = "en0"
		log.Printf("No INTERFACE provided, using default: %s", iface)
	}

	timeout := 10 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	done := make(chan error, 1)
	go func() {
		if err := mac.Sniff(ctx, iface); err != nil {
			done <- fmt.Errorf("error sniffing MAC addresses: %w", err)
		} else {
			done <- nil
		}
	}()

	select {
	case err := <-done:
		// Sniff function completed within the timeout
		return err
	case <-ctx.Done():
		// Timeout occurred or context was cancelled
		return fmt.Errorf("operation timed out")
	}
}

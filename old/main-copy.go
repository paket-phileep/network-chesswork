package mainasdd

import (
	"flag"
	"fmt"
	"log"
	"network-chesswork/cronJob"
	"network-chesswork/info" // Assuming this is the correct package path

	"github.com/robfig/cron/v3" // Updated import path for the robfig/cron package
)

var (
	ssidFlag = flag.String("ssid", "", "SSID to connect to or use existing connection")
)

func main() {
	c := cron.New()
	flag.Parse()

	if *ssidFlag == "" {
		// Implement logic to use an existing connection
		fmt.Println("No SSID provided, using existing connection.")

		// Get the information from the existing connection about the IP ranges, etc.
		networkInfo, err := info.Network()
		if err != nil {
			log.Fatalf("Error getting network information: %v", err)
		}

		fmt.Println("Successfully gathered network information")
		for key, value := range networkInfo {
			fmt.Printf("%s: %s\n", key, value)
		}

		// Extract the CIDR Notation
		cidrNotation, ok := networkInfo["CIDR Notation"]
		if !ok {
			log.Fatalf("CIDR Notation not found in network information")
		}

		// initial run as this runs every 2 hours
		go cronJob.FindConnectedMacAddresses(cidrNotation)
		c.AddFunc("@every 2h30m", func() {
			fmt.Println("Every 2 hours 30 minutes")
			go cronJob.FindConnectedMacAddresses(cidrNotation)
		})

	} else {
		// Implement the logic to handle the SSID parameter
		fmt.Println("SSID flag provided:", *ssidFlag)
	}

	c.AddFunc("@every 1s", func() {
		// fmt.Println("Every 1 second")
		go cronJob.HealthCheck("www.google.com")
	})

	c.Start()

	// Keep the main function running to allow cron jobs to work
	select {} // This blocks indefinitely
}

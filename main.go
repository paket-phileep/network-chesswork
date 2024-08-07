package main

import (
	"flag"
	"fmt"
	"log"
	"network-chesswork/cronJob"
	"network-chesswork/info"

	"github.com/robfig/cron"
)

var (
	ssidFlag = flag.String("ssid", "", "SSID to connect to or use existing connection")
)

func main() {
	flag.Parse()
	// parse the cli flags
	// Access the value using the pointer
	if *ssidFlag != "" {
		// TODO: Implement the logic to handle the SSID parameter
		fmt.Println("SSID flag provided:", *ssidFlag)
	} else {
		// TODO: Implement logic to use an existing connection
		fmt.Println("No SSID provided, using existing connection.")

		// get the information from the existing connection about the ip ranges ect

		networkInfo, err := info.Network()
			return
		}
		fmt.Println("Successfully gathered network information")
		for key, value := range networkInfo {
			fmt.Printf("%s: %s\n", key, value)
		}
	}

	c := cron.New()

	// Schedule the health check to run every minute
	_, err := c.AddFunc("* * * * *", cronJob.HealthCheck)
	if err != nil {
		log.Fatalf("Error adding health check cron job: %v", err)
	}
	fmt.Printf("Added health check cron job with entry ID: %v\n", entryID)

	// Schedule the MAC gathering to run every three hours
	_, err = c.AddFunc("0 */3 * * *", cronJob.FindConnectedMacAddresses)
	if err != nil {
		log.Fatalf("Error adding MAC gathering cron job: %v", err)
	}
	fmt.Printf("Added MAC gathering cron job with entry ID: %v\n", entryID)

	c.Start()

	// Keep the program running
	select {}

}

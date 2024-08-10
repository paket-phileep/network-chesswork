package main

import (
	"flag"
	"fmt"
	"network-chesswork/cronJob"

	"github.com/robfig/cron/v3"
)

var (
	interfaceFlag = flag.String("interface", "", "INTERFACE to connect to or en0")
)

func main() {
	c := cron.New()
	flag.Parse()

	if *interfaceFlag == "" {
		// Implement logic to use an existing connection
		fmt.Println("No INTERFACE provided")

		go cronJob.FindConnectedMacAddresses("")

		c.AddFunc("@every 2h30m", func() {
			fmt.Println("Every 2 hours 30 minutes")

			go cronJob.FindConnectedMacAddresses("")
		})

	} else {
		// Implement the logic to handle the INTERFACE parameter
		fmt.Println("INTERFACE flag provided:", *interfaceFlag)
	}

	// c.AddFunc("@every 1s", func() {
	// 	// fmt.Println("Every 1 second")
	// 	go cronJob.HealthCheck("www.google.com")
	// })

	c.Start()

	// Keep the main function running to allow cron jobs to work
	select {} // This blocks indefinitely
}

package net

import (
	"fmt"

	"github.com/go-ping/ping"
)

func Ping(host string) error {
	pinger, err := ping.NewPinger(host)
	if err != nil {
		return err
	}
	pinger.Count = 3
	err = pinger.Run() // Blocks until finished.
	if err != nil {
		panic(err)
		
	}
	stats := pinger.Statistics() // get send/receive/duplicate/rtt stats

	fmt.Printf("Type of stats: %T\n", stats)

	return nil

}

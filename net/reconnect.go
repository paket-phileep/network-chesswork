package net

import (
	"log"
	"network-chesswork/mac"

	"network-chesswork/utilities"
)

func Reconnect(iface string) error {
	if iface == "" {
		iface = "en0"
		log.Printf("No INTERFACE provided, using default: %s", iface)
	}

	keys, err := utilities.ReadKeysJSON("../temp/source-mac.json")
	if err != nil {
		return err
	}
	for i := 0; i < len(keys); i++ {
		mac.ChangeMAC(iface, keys[i])
		err := Ping("www.google.com")
		if err != nil {
			continue
		} else {
			return nil
		}
	}

	return nil

}

package main

import (
	"fmt"
	"log"
	"network-chesswork/utilities"
	"os"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <network_interface>")
		os.Exit(1)
	}

	iface := os.Args[1]

	// Open the network interface for packet capture
	handle, err := pcap.OpenLive(iface, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	macAddresses := make(map[string]bool)

	for packet := range packetSource.Packets() {
		// Extract Ethernet layer
		ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
		if ethernetLayer != nil {
			ethernetPacket, ok := ethernetLayer.(*layers.Ethernet)
			if !ok {
				continue
			}

			srcMAC := ethernetPacket.SrcMAC.String()
			dstMAC := ethernetPacket.DstMAC.String()

			if _, found := macAddresses[srcMAC]; !found {
				macAddresses[srcMAC] = true
				fmt.Println("Source MAC Address:", srcMAC)
				// Specify the file name
				path := "./temp/source-mac.json"

				data := map[string]interface{}{
					srcMAC: map[string]interface{}{
						"time": time.Now().UTC().Format(time.RFC3339),
					},
				}
				utilities.AppendJSON(path, data)
			}

			if _, found := macAddresses[dstMAC]; !found {
				macAddresses[dstMAC] = true
				fmt.Println("Destination MAC Address:", dstMAC)
				path := "./temp/dest-mac.json"

				data := map[string]interface{}{
					srcMAC: map[string]interface{}{
						"time": time.Now().UTC().Format(time.RFC3339),
					},
				}

				utilities.AppendJSON(path, data)
			}
		}
	}
}

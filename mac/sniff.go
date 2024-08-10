package mac

import (
	"context"
	"fmt"
	"network-chesswork/utilities"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func Sniff(ctx context.Context, iface string) error {
	if iface == "" {
		return fmt.Errorf("network interface is required")
	} else {
		fmt.Println("Sniffing package received interface", iface)
	}

	// Open the network interface for packet capture
	handle, err := pcap.OpenLive(iface, 1600, true, pcap.BlockForever)
	if err != nil {
		return fmt.Errorf("failed to open live capture: %w", err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	macAddresses := make(map[string]bool)

	for {
		select {
		case <-ctx.Done():
			// Handle context cancellation
			fmt.Println("Sniffing stopped due to context cancellation")
			return ctx.Err()
		case packet := <-packetSource.Packets():
			fmt.Println("Packet received:", packet)
			fmt.Println("Packet layers:", packet.Layers())

			// Extract Ethernet layer
			ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
			if ethernetLayer == nil {
				fmt.Println("No Ethernet layer found in packet")
				continue
			}

			ethernetPacket, ok := ethernetLayer.(*layers.Ethernet)
			if !ok {
				fmt.Println("Failed to assert Ethernet layer")
				continue
			}

			srcMAC := ethernetPacket.SrcMAC.String()
			dstMAC := ethernetPacket.DstMAC.String()

			if !macAddresses[srcMAC] {
				macAddresses[srcMAC] = true
				fmt.Println("Source MAC Address:", srcMAC)
				path := "./temp/source-mac.json"

				data := map[string]interface{}{
					srcMAC: map[string]interface{}{
						"time": time.Now().UTC().Format(time.RFC3339),
					},
				}
				if err := utilities.AppendJSON(path, data); err != nil {
					fmt.Printf("Error appending source MAC address to JSON: %v\n", err)
					return fmt.Errorf("failed to append source MAC address to JSON: %w", err)
				}
			}

			if !macAddresses[dstMAC] {
				macAddresses[dstMAC] = true
				fmt.Println("Destination MAC Address:", dstMAC)
				path := "./temp/dest-mac.json"

				data := map[string]interface{}{
					dstMAC: map[string]interface{}{
						"time": time.Now().UTC().Format(time.RFC3339),
					},
				}
				if err := utilities.AppendJSON(path, data); err != nil {
					fmt.Printf("Error appending destination MAC address to JSON: %v\n", err)
					return fmt.Errorf("failed to append destination MAC address to JSON: %w", err)
				}
			}
		}
	}
}

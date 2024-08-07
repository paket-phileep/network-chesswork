package info

import (
	"fmt"
	"net"
)

// getLocalNetworkInfo retrieves the local machine's IP address and subnet mask
func getLocalNetworkInfo() (string, string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", "", err
	}

	for _, iface := range interfaces {
		if iface.Flags&net.FlagUp == 0 || (iface.Flags&net.FlagLoopback != 0) {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			return "", "", err
		}

		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && ipnet.IP.To4() != nil {
				return ipnet.IP.String(), net.IP(ipnet.Mask).String(), nil
			}
		}
	}

	return "", "", fmt.Errorf("no active network interface found")
}

// convertNetmaskToCIDR converts a netmask to its CIDR prefix length
func convertNetmaskToCIDR(mask string) (int, error) {
	ipMask := net.IPMask(net.ParseIP(mask).To4())
	ones, _ := ipMask.Size()
	return ones, nil
}

// calculateIPRange calculates the IP range for the given IP address and subnet mask
func calculateIPRange(ip, mask string) (string, string, string, string, error) {
	ipAddr := net.ParseIP(ip).To4()
	if ipAddr == nil {
		return "", "", "", "", fmt.Errorf("invalid IP address")
	}

	maskIP := net.IPMask(net.ParseIP(mask).To4())
	ipNet := &net.IPNet{IP: ipAddr, Mask: maskIP}

	network := ipNet.IP.Mask(ipNet.Mask)
	broadcast := make(net.IP, len(ipAddr))
	for i := range ipAddr {
		broadcast[i] = ipAddr[i] | ^maskIP[i]
	}

	// Calculating the first and last host addresses
	firstHost := make(net.IP, len(ipAddr))
	copy(firstHost, network)
	firstHost[len(firstHost)-1]++

	lastHost := make(net.IP, len(broadcast))
	copy(lastHost, broadcast)
	lastHost[len(lastHost)-1]--

	return network.String(), broadcast.String(), firstHost.String(), lastHost.String(), nil
}

func Network() (map[string]string, error) {
	ip, netmask, err := getLocalNetworkInfo()
	if err != nil {
		fmt.Println("Error getting local network info:", err)
		return nil, err
	}

	cidr, err := convertNetmaskToCIDR(netmask)
	if err != nil {
		fmt.Println("Error converting netmask to CIDR:", err)
		return nil, err
	}

	network, broadcast, firstHost, lastHost, err := calculateIPRange(ip, netmask)
	if err != nil {
		fmt.Println("Error calculating IP range:", err)
		return nil, err
	}

	// Return a map with IP range details
	ipRangeDetails := map[string]string{
		"IP Address":        ip,
		"Subnet Mask":       netmask,
		"CIDR Notation":     fmt.Sprintf("%s/%d", ip, cidr),
		"Network Address":   network,
		"Broadcast Address": broadcast,
		"First Host":        firstHost,
		"Last Host":         lastHost,
	}

	// for key, value := range ipRangeDetails {
	// 	fmt.Printf("%s: %s\n", key, value)
	// }

	return ipRangeDetails, nil

}

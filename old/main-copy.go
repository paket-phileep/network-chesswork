package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"runtime"
)

// const pingTarget = "8.8.8.8" // Hardcoded ping target

// // fetchWiFiInterfaceName retrieves the Wi-Fi interface name based on the OS.
// func fetchWiFiInterfaceName() (string, error) {
// 	var cmd *exec.Cmd
// 	var output []byte
// 	var err error

// 	switch runtime.GOOS {
// 	case "linux":
// 		cmd = exec.Command("nmcli", "-t", "-f", "DEVICE,TYPE", "device")
// 	case "darwin":
// 		cmd = exec.Command("networksetup", "-listallhardwareports")
// 	case "windows":
// 		cmd = exec.Command("netsh", "wlan", "show", "interfaces")
// 	default:
// 		return "", fmt.Errorf("unsupported OS")
// 	}

// 	output, err = cmd.CombinedOutput()
// 	if err != nil {
// 		return "", err
// 	}

// 	var re *regexp.Regexp
// 	var match []string
// 	outputStr := string(output)
// 	fmt.Print(outputStr)

// 	switch runtime.GOOS {
// 	case "linux":
// 		re = regexp.MustCompile(`^(\S+):wifi`)
// 		match = re.FindStringSubmatch(outputStr)
// 	case "darwin":
// 		re = regexp.MustCompile(`Hardware Port: Wi-Fi\nDevice: (\w+)`)
// 		match = re.FindStringSubmatch(outputStr)
// 	case "windows":
// 		re = regexp.MustCompile(`^\s*Name\s*:\s*(\S+)`)
// 		match = re.FindStringSubmatch(outputStr)
// 	}

// 	if len(match) > 1 {
// 		return match[1], nil
// 	}
// 	return "", fmt.Errorf("Wi-Fi interface not found")
// }

// // fetchCurrentSSID retrieves the SSID of the currently connected Wi-Fi network based on the OS.
// func fetchCurrentSSID() (string, error) {
// 	var cmd *exec.Cmd
// 	switch runtime.GOOS {
// 	case "linux":
// 		cmd = exec.Command("nmcli", "-t", "-f", "SSID", "connection", "show", "--active")
// 	case "darwin":
// 		cmd = exec.Command("networksetup", "-getairportnetwork", "en0") // Replace 'en0' with your Wi-Fi interface
// 	case "windows":
// 		cmd = exec.Command("netsh", "wlan", "show", "interfaces")
// 	default:
// 		return "", fmt.Errorf("unsupported OS")
// 	}

// 	output, err := cmd.CombinedOutput()

// 	if err != nil {
// 		return "", err
// 	}

// 	var re *regexp.Regexp
// 	var match []string
// 	outputStr := string(output)
// 	fmt.Print(outputStr)

// 	switch runtime.GOOS {
// 	case "linux":
// 		return strings.TrimSpace(outputStr), nil
// 	case "darwin":
// 		re = regexp.MustCompile(`Current Network : (.*)`)
// 		match = re.FindStringSubmatch(outputStr)
// 	case "windows":
// 		re = regexp.MustCompile(`^\s*SSID\s*:\s*(.*)`)
// 		match = re.FindStringSubmatch(outputStr)
// 	}

// 	if len(match) > 1 {
// 		return strings.TrimSpace(match[1]), nil
// 	}
// 	return "", fmt.Errorf("SSID not found")
// }

// // fetchBSSID retrieves the BSSID of the currently connected Wi-Fi network based on the OS.
// func fetchBSSID() (string, error) {
// 	var cmd *exec.Cmd
// 	switch runtime.GOOS {
// 	case "linux":
// 		cmd = exec.Command("iwconfig")
// 	case "darwin":
// 		cmd = exec.Command("/System/Library/PrivateFrameworks/Apple80211.framework/Versions/Current/Resources/airport", "-I")
// 	case "windows":
// 		cmd = exec.Command("netsh", "wlan", "show", "interfaces")
// 	default:
// 		return "", fmt.Errorf("unsupported OS")
// 	}

// 	output, err := cmd.CombinedOutput()
// 	if err != nil {
// 		return "", err
// 	}

// 	var re *regexp.Regexp
// 	var match []string

// 	switch runtime.GOOS {
// 	case "linux":
// 		re = regexp.MustCompile(`Access Point: ([0-9A-Fa-f:]{17})`)
// 		match = re.FindStringSubmatch(string(output))
// 	case "darwin":
// 		re = regexp.MustCompile(`BSSID:\s+(.*)`)
// 		match = re.FindStringSubmatch(string(output))
// 	case "windows":
// 		re = regexp.MustCompile(`^\s*BSSID\s*:\s*(.*)`)
// 		match = re.FindStringSubmatch(string(output))
// 	}

// 	if len(match) > 1 {
// 		return strings.TrimSpace(match[1]), nil
// 	}
// 	return "", fmt.Errorf("BSSID not found")
// }

// // scanNetworkForMACs scans the local network for connected devices and returns their MAC addresses.

// // scanNetworkForMACs scans the network for MAC addresses based on the operating system.
// func scanNetworkForMACs() ([]string, error) {
// 	var cmd *exec.Cmd
// 	var output []byte
// 	var err error

// 	switch runtime.GOOS {
// 	case "linux":
// 		cmd = exec.Command("sudo", "arp-scan", "-l")
// 		output, err = cmd.CombinedOutput()
// 		if err != nil {
// 			return nil, fmt.Errorf("command execution error: %w", err)
// 		}

// 	case "darwin":
// 		cmd = exec.Command("arp", "-a")
// 		output, err = cmd.CombinedOutput()
// 		if err != nil {
// 			return nil, fmt.Errorf("command execution error: %w", err)
// 		}

// 	case "windows":
// 		cmd = exec.Command("arp", "-a")
// 		output, err = cmd.CombinedOutput()
// 		if err != nil {
// 			return nil, fmt.Errorf("command execution error: %w", err)
// 		}

// 	default:
// 		return nil, fmt.Errorf("unsupported OS")
// 	}

// 	// Define a regex pattern for MAC addresses (6 pairs of hex digits separated by colons or hyphens)
// 	re := regexp.MustCompile(`(?:[0-9A-Fa-f]{2}[:-]){5}[0-9A-Fa-f]{2}`)

// 	// Convert the output to a string and find all matches
// 	outputStr := string(output)
// 	matches := re.FindAllString(outputStr, -1)

// 	// Clean up matches if necessary (e.g., remove duplicates)
// 	matches = removeDuplicates(matches)

// 	return matches, nil
// }

// // removeDuplicates removes duplicate strings from a slice.
// func removeDuplicates(stringsSlice []string) []string {
// 	keys := make(map[string]bool)
// 	list := []string{}
// 	for _, entry := range stringsSlice {
// 		if _, value := keys[entry]; !value {
// 			keys[entry] = true
// 			list = append(list, entry)
// 		}
// 	}
// 	return list
// }

// // ping checks the connectivity to the given target.
// func ping(target string) bool {
// 	cmd := exec.Command("ping", "-c", "1", target)
// 	err := cmd.Run()
// 	return err == nil
// }

// // executeCommand runs a shell command.
// func executeCommand(command string) error {
// 	fmt.Println("Executing:", command)
// 	cmd := exec.Command("bash", "-c", command)
// 	return cmd.Run()
// }

// func main() {
// 	// Define command-line flags
// 	macToSpoof := flag.String("mac", "", "MAC address to spoof (optional)")
// 	bssid := flag.String("bssid", "", "BSSID of the access point (optional)")
// 	clientMac := flag.String("client", "", "MAC address of the client to deauthenticate (optional)")
// 	interval := flag.Int("interval", 1, "Ping interval in minutes")

// 	// Parse command-line arguments
// 	flag.Parse()

// 	// Fetch the Wi-Fi interface name
// 	wifiInterface, err := fetchWiFiInterfaceName()
// 	if err != nil {
// 		fmt.Println("Error fetching Wi-Fi interface:", err)
// 		return
// 	}

// 	// Fetch the current SSID
// 	ssid, err := fetchCurrentSSID()
// 	if err != nil {
// 		fmt.Println("Error fetching current SSID:", err)
// 		return
// 	}

// 	fmt.Printf("Currently connected to SSID: %s\n", ssid)

// 	// Fetch BSSID if not provided
// 	if *bssid == "" {
// 		bssidVal, err := fetchBSSID()
// 		if err != nil {
// 			fmt.Println("Error fetching BSSID:", err)
// 			return
// 		}
// 		*bssid = bssidVal
// 		fmt.Printf("Using BSSID: %s\n", *bssid)
// 	}

// 	// Use the provided MAC address or fetch a default one
// 	if *macToSpoof == "" {
// 		fmt.Println("MAC address to spoof is not provided. Fetching a MAC address from the network.")
// 		macAddresses, err := scanNetworkForMACs()
// 		if err != nil {
// 			fmt.Println("Error scanning network for MAC addresses:", err)
// 			return
// 		}

// 		if len(macAddresses) > 0 {
// 			*macToSpoof = macAddresses[0] // Choose the first MAC address found
// 			fmt.Printf("Using MAC address from network: %s\n", *macToSpoof)
// 		} else {
// 			fmt.Println("No MAC addresses found on the network. Exiting.")
// 			return
// 		}
// 	} else {
// 		fmt.Printf("Using provided MAC address: %s\n", *macToSpoof)
// 	}

// 	for {
// 		if !ping(pingTarget) {
// 			fmt.Println("Ping failed. Taking corrective actions.")

// 			// Bring the interface down
// 			if err := executeCommand(fmt.Sprintf("sudo ip link set %s down", wifiInterface)); err != nil {
// 				fmt.Println("Error bringing interface down:", err)
// 				continue
// 			}

// 			// Spoof the MAC address
// 			if err := executeCommand(fmt.Sprintf("sudo macchanger --mac %s %s", *macToSpoof, wifiInterface)); err != nil {
// 				fmt.Println("Error spoofing MAC address:", err)
// 				continue
// 			}

// 			// Bring the interface back up
// 			if err := executeCommand(fmt.Sprintf("sudo ip link set %s up", wifiInterface)); err != nil {
// 				fmt.Println("Error bringing interface up:", err)
// 				continue
// 			}

// 			// Deauth client
// 			if *clientMac != "" {
// 				if err := executeCommand(fmt.Sprintf("sudo aireplay-ng -0 1 -a %s -c %s %s", *bssid, *clientMac, wifiInterface)); err != nil {
// 					fmt.Println("Error deauthenticating client:", err)
// 					continue
// 				}
// 			}

// 			// Release and renew IP address
// 			if err := executeCommand(fmt.Sprintf("sudo dhclient -r && sudo dhclient -v %s", wifiInterface)); err != nil {
// 				fmt.Println("Error releasing and renewing IP address:", err)
// 				continue
// 			}

// 			// Optional: sleep for a short while to ensure changes take effect
// 			time.Sleep(10 * time.Second)
// 		} else {
// 			fmt.Println("Ping successful.")
// 		}

// 		// Wait for the next ping interval
// 		time.Sleep(time.Duration(*interval) * time.Minute)
// 	}
// }

func fetchWiFiInterfaceName() (string, error) {
	var cmd *exec.Cmd
	var output []byte
	var err error

	// Determine the command based on the operating system
	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("nmcli", "-t", "-f", "DEVICE,TYPE", "device")
	case "darwin":
		cmd = exec.Command("networksetup", "-listallhardwareports")
	case "windows":
		cmd = exec.Command("netsh", "wlan", "show", "interfaces")
	default:
		return "", fmt.Errorf("unsupported OS")
	}

	// Execute the command and capture the output
	output, err = cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	// Define the regular expression based on the operating system
	var re *regexp.Regexp
	var match []string
	outputStr := string(output)

	switch runtime.GOOS {
	case "linux":
		re = regexp.MustCompile(`^(\S+):wifi`)
		match = re.FindStringSubmatch(outputStr)
	case "darwin":
		re = regexp.MustCompile(`(?m)^Hardware Port:\s*Wi-Fi\s*Device:\s*(\S+)`)
		match = re.FindStringSubmatch(outputStr)
	case "windows":
		re = regexp.MustCompile(`(?m)^Name\s*:\s*(\S+)`)
		match = re.FindStringSubmatch(outputStr)
	}

	if len(match) > 1 {
		return match[1], nil
	}

	return "", fmt.Errorf("Wi-Fi interface not found")
}

// downInterfaceConn brings down a network interface and returns a message and an error if any.
func downInterfaceConn(interfaceName string) (string, error) {
	// Example command to bring down the interface (this may vary by OS)
	cmd := exec.Command("sudo", "ifconfig", interfaceName, "down")

	// Run the command
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to bring down interface %s: %w", interfaceName, err)
	}

	return "Interface brought down successfully", nil
}

// func downInterfaceConn(interface: string) (string, error) {
// 	var cmd *exec.Cmd
// 	var output []byte
// 	var err error

// 	// Determine the command based on the operating system
// 	switch runtime.GOOS {
// 	case "linux":
// 		cmd = exec.Command("nmcli", "-t", "-f", "DEVICE,TYPE", "device")
// 	case "darwin":
// 		cmd = exec.Command("networksetup", "-listallhardwareports")
// 	case "windows":
// 		cmd = exec.Command("netsh", "wlan", "show", "interfaces")
// 	default:
// 		return "", fmt.Errorf("unsupported OS")
// 	}

// 	// Execute the command and capture the output
// 	output, err = cmd.CombinedOutput()
// 	if err != nil {
// 		return "", err
// 	}

// 	// Define the regular expression based on the operating system
// 	var re *regexp.Regexp
// 	var match []string
// 	outputStr := string(output)

// 	switch runtime.GOOS {
// 	case "linux":
// 		re = regexp.MustCompile(`^(\S+):wifi`)
// 		match = re.FindStringSubmatch(outputStr)
// 	case "darwin":
// 		re = regexp.MustCompile(`(?m)^Hardware Port:\s*Wi-Fi\s*Device:\s*(\S+)`)
// 		match = re.FindStringSubmatch(outputStr)
// 	case "windows":
// 		re = regexp.MustCompile(`(?m)^Name\s*:\s*(\S+)`)
// 		match = re.FindStringSubmatch(outputStr)
// 	}

// 	if len(match) > 1 {
// 		return match[1], nil
// 	}

// 	return "", fmt.Errorf("Wi-Fi interface not found")
// }

func main() {
	var iface string
	var err error

	// Fetch WiFi interface name
	if iface, err = fetchWiFiInterfaceName(); err != nil {
		fmt.Println("Error:", err)
		return // Exit if there's an error
	}

	fmt.Println("Proceeding with interface:", iface)

	// Bring down the interface connection
	if downInterface, err := downInterfaceConn(iface); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Status", downInterface)
	}

	if iface, err := fetchWiFiInterfaceName(); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Proceeding with interface:", iface)
	}
	if iface, err := fetchWiFiInterfaceName(); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Proceeding with interface:", iface)
	}
	if iface, err := fetchWiFiInterfaceName(); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Proceeding with interface:", iface)
	}
	if iface, err := fetchWiFiInterfaceName(); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Proceeding with interface:", iface)
	}
	if iface, err := fetchWiFiInterfaceName(); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Proceeding with interface:", iface)
	}
	if iface, err := fetchWiFiInterfaceName(); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Proceeding with interface:", iface)
	}
}

# network-chesswork

This Go application is designed to monitor and manage Wi-Fi connectivity on various operating systems. Here’s a breakdown of its functionality:

1. **Fetch Wi-Fi Interface Name**: The app determines the name of the Wi-Fi interface based on the operating system (Linux, macOS, Windows).
2. **Fetch Current SSID**: It retrieves the SSID (network name) of the currently connected Wi-Fi network.
3. **Fetch BSSID**: The app retrieves the BSSID (MAC address of the access point) of the currently connected Wi-Fi network. If not provided via command-line arguments, it will be fetched automatically.
4. **Scan Network for MAC Addresses**: It scans the local network to find MAC addresses of connected devices. This is useful for network diagnostics or spoofing.
5. **Ping Check**: The application periodically pings a target IP address (8.8.8.8, a Google DNS server) to check for connectivity.
6. **Corrective Actions on Ping Failure**:

   - **Bring Interface Down**: If the ping fails, the app first brings down the Wi-Fi interface.
   - **MAC Address Spoofing**: It then spoofs (changes) the MAC address of the Wi-Fi interface.
   - **Bring Interface Up**: The interface is brought back up.
   - **Deauthenticate Client**: If a specific client MAC address is provided, the app will deauthenticate it from the network (using `aireplay-ng` for Linux).
   - **Release and Renew IP Address**: Finally, it releases and renews the IP address assigned to the Wi-Fi interface.

7. **Recheck and Sleep**: The application sleeps for a specified interval (default 1 minute) and repeats the ping check.

### Usage Scenario

This app could be used in a scenario where:

- You want to monitor your Wi-Fi connection and perform automated recovery actions if connectivity is lost.
- You are troubleshooting network issues and need to frequently change the MAC address or reset the connection to test different configurations.
- The application automates network management tasks, which can be helpful in environments with unstable Wi-Fi connections.

### Command-Line Flags

- `-mac`: The MAC address to spoof (optional). If not provided, the app will use a MAC address found on the network.
- `-bssid`: The BSSID of the access point (optional). If not provided, it will be fetched automatically.
- `-client`: The MAC address of a client to deauthenticate (optional).
- `-interval`: The interval in minutes between ping checks (default is 1 minute).

Overall, the application aims to keep the Wi-Fi connection stable by automatically performing corrective actions when connectivity issues are detected.

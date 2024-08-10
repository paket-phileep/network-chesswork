#!/bin/bash

# Parse command-line arguments
while [[ "$#" -gt 0 ]]; do
    case $1 in
        --range=*) range="${1#*=}"; shift ;;
        *) echo "Unknown parameter passed: $1"; exit 1 ;;
    esac
done

# Check if the range variable is set
if [ -z "$range" ]; then
    echo "Error: Network range (--range) is not set."
    exit 1
fi

# Ensure nmap is installed
if ! command -v nmap &> /dev/null; then
    echo "Error: nmap is not installed."
    exit 1
fi

# Ensure the output directory exists
mkdir -p "./temp/nmap"

# script -c "nmap -sn \"$range\"" ./temp/nmap/network-clients.log
nmap -sn "$range" > ./temp/nmap/network-clients.log 2> error.log

# # Run nmap scan
# nmap -sn "$range" -oN "./temp/nmap/network-clients.log" 2> error.log

# Check if nmap command was successful
if [ $? -ne 0 ]; then
    echo "Error: nmap scan failed."
    exit 1
fi

echo "Scan completed successfully."

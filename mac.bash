#!/bin/bash

# Check if the range variable is set
if [ -z "$range" ]; then
  echo "Error: Network range (props.range) is not set."
  exit 1
fi

# Ensure nmap is installed
if ! command -v nmap &> /dev/null; then
  echo "Error: nmap is not installed."
  exit 1
fi

# Run nmap scan
echo "Scanning network range: $range"
nmap -sn "$range"

# Check if nmap command was successful
if [ $? -ne 0 ]; then
  echo "Error: nmap scan failed."
  exit 1
fi

echo "Scan completed successfully."

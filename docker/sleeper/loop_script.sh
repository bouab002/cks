#!/bin/bash

# Loop indefinitely
while true; do
    # Print the user ID
    echo "User ID: $(id -u)"
    
    # List the /proc directory
    echo "Contents of /proc directory:"
    ls /proc
    
    # Sleep for 60 seconds
    sleep 60
done
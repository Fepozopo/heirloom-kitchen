#!/bin/bash

# Exit immediately if a command exits with a non-zero status
set -e

# Generate the site
echo "Generating the site..."
go build -o heirloom-kitchen
./heirloom-kitchen

# Check if the build was successful
if [ $? -ne 0 ]; then
    echo "Build failed. Exiting."
    exit 1
fi

echo "Site generated successfully."
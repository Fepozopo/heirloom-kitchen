#!/bin/bash

# Replace "REPO_NAME" with your actual GitHub repository name
REPO_NAME="pantry-of-the-past"

# Compile the Go application
go build -o pantry-of-the-past src/main.go

# Check if the build was successful
if [ $? -ne 0 ]; then
    echo "Build failed. Please check the Go code for errors."
    exit 1
fi
# Check if the compiled application exists
if [ ! -f pantry-of-the-past ]; then
    echo "Compiled application not found."
    exit 1
fi
# Check if the repository name is provided
if [ -z "$REPO_NAME" ]; then
    echo "Repository name is required."
    exit 1
fi

echo "Build successful."
echo "Running the application with repository name: $REPO_NAME"

# Execute the compiled application with the repository name as an argument
./pantry-of-the-past "$REPO_NAME"
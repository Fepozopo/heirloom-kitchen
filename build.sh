#!/bin/bash

REPO_NAME="culinary-keepsakes"

# Compile the Go application
go build -o bin/app src/main.go

# Check if the build was successful
if [ $? -ne 0 ]; then
    echo "Build failed. Please check the Go code for errors."
    exit 1
fi
# Check if the compiled application exists
if [ ! -f bin/app ]; then
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
./bin/app "/$REPO_NAME/"
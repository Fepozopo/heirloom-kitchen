#!/bin/bash

REPO_NAME="culinary-keepsakes"

# Compile the Go application
go build -o bin/app src/main.go

echo "Build successful."

./bin/app
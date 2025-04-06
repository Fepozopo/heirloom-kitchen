package main

import (
	"fmt"

	"github.com/Fepozopo/heirloom-kitchen/src/website"
)

func main() {
	// Copy static files to public directory
	err := website.CopyStaticToPublic()
	if err != nil {
		fmt.Println("Error copying static files:", err)
		return
	}
	fmt.Println("Static files copied successfully.")
}

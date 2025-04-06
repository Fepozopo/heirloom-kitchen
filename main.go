package main

import (
	"fmt"
	"log"
	"net/http"

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

	// Generate the markdown pages
	err = website.GeneratePagesRecursive("content", "template.html", "public")
	if err != nil {
		fmt.Println("Error generating pages:", err)
		return
	}
	fmt.Println("Pages generated successfully.")

	// Start the HTTP server to serve the public directory
	http.Handle("/", http.FileServer(http.Dir("public")))
	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

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

	// Generate the index page
	err = website.GeneratePage("content/index.md", "template.html", "public/index.html")
	if err != nil {
		fmt.Println("Error generating index page:", err)
		return
	}
	fmt.Println("Index page generated successfully.")

	// Start the HTTP server to serve the public directory
	http.Handle("/", http.FileServer(http.Dir("public")))
	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

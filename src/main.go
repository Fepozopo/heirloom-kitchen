package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Fepozopo/pantry-of-the-past/src/website"
)

func main() {
	// Grab the basepath from the first CLI argument or default to "/"
	basepath := "/"
	if len(os.Args) > 1 {
		basepath = os.Args[1]
	}

	// Copy static files to docs directory
	err := website.CopyStaticToPublic()
	if err != nil {
		fmt.Println("Error copying static files:", err)
		return
	}
	fmt.Println("Static files copied successfully.")

	// Generate the markdown pages
	err = website.GeneratePagesRecursive("content", "template.html", "docs", basepath)
	if err != nil {
		fmt.Println("Error generating pages:", err)
		return
	}
	fmt.Println("Pages generated successfully.")

	// Start the HTTP server to serve the docs directory
	http.Handle("/", http.FileServer(http.Dir("docs")))
	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

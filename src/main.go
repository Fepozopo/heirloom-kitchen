package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Fepozopo/culinary-keepsakes/src/website"
)

func main() {
	// Grab the basepath from the first CLI argument or default to "/"
	basepath := "/"
	if len(os.Args) > 1 {
		basepath = os.Args[1]
		if !strings.HasSuffix(basepath, "/") {
			basepath += "/"
		}
		if !strings.HasPrefix(basepath, "/") {
			basepath = "/" + basepath
		}
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

	// Start the HTTP server to serve the docs directory, including static files with basepath
	fs := http.FileServer(http.Dir("docs"))
	http.Handle(basepath, http.StripPrefix(basepath, fs))

	fmt.Println("Server started at http://localhost:8080" + basepath)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

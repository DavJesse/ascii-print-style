package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"web/web"
)

func main() {
	// Kill server when arguments other than the program name are added
	if len(os.Args) != 1 {
		fmt.Println("Too many arguments")
		return
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", Web.SubmitFormHandler)               // Serve home page
	http.HandleFunc("/files/art.txt", Web.DownloadArtHandler) // Route to serve the download
	http.HandleFunc("/ascii-art", Web.SubmitFormHandler)      // Serve /ascii-art at POST requests
	log.Fatal(http.ListenAndServe(":8000", nil))              // Defines host port
}

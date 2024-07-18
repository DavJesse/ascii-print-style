package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"web/Web"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Too many arguments")
		return
	}

	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("styles"))))
	http.HandleFunc("/", Web.SubmitFormHandler)
	http.HandleFunc("/ascii-art", Web.SubmitFormHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

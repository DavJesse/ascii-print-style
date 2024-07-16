package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"web/Web"
)

func main() {
	// End program should user provide more than one argument
	if len(os.Args) > 1 {
		fmt.Println("Too many arguments")
		return
	}

	http.HandleFunc("/", Web.FormHandler)
	http.HandleFunc("/submit-ascii-art", Web.SubmitFormHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Fatal(http.ListenAndServe(":8000", nil))

}

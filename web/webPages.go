package Web

import (
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"web/lib"
)

var output, err string

func SubmitFormHandler(w http.ResponseWriter, r *http.Request) {
	var bnStyle, inputStr string
	var tmpl *template.Template

	// Label any requests other than 'GET' and 'POST' requests as 'invalid requests'
	if !(r.Method == http.MethodGet || r.Method == http.MethodPost) {
		w.WriteHeader(http.StatusMethodNotAllowed)

		tmpl = template.Must(template.ParseFiles("templates/errorPrinter.html"))
		tmpl.Execute(w, struct {
			Code  int
			Issue string
		}{Issue: "Invalid Request Method!", Code: http.StatusMethodNotAllowed})
		return

		// Serve form at initial visit of site
	} else if r.Method == http.MethodGet {
		if r.URL.Path == "/ascii-art" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			tmpl = template.Must(template.ParseFiles("templates/errorPrinter.html"))
			tmpl.Execute(w, struct {
				Issue string
				Code  int
			}{Issue: "405: Method not allowed", Code: http.StatusMethodNotAllowed})
			return
		}
		if r.URL.Path != "/" {
			w.WriteHeader(http.StatusNotFound)
			tmpl = template.Must(template.ParseFiles("templates/errorPrinter.html"))
			tmpl.Execute(w, struct {
				Issue string
				Code  int
			}{Issue: "404: Page not found", Code: http.StatusNotFound})
			return
		}
		tmpl = template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, nil)

		// Serve form and ascii-art/error after form submission
	} else if r.Method == http.MethodPost {

		// Extract banner style selected and text inputted in form
		bnStyle = r.FormValue("style")
		if bnStyle == "" {
			bnStyle = "standard"
		}
		inputStr = r.FormValue("inputStr")

		// Generate the ASCII art
		output, err = lib.AsciiArt(inputStr, bnStyle+".txt")

		if err != "" {
			// Handle errors by rendering an error template
			tmpl = template.Must(template.ParseFiles("templates/errorPrinter.html"))

			if strings.Contains(err, "PRINTABLE ASCII") {
				w.WriteHeader(http.StatusBadRequest)
				tmpl.Execute(w, struct {
					Issue string
					Code  int
				}{Issue: err, Code: http.StatusBadRequest})
			}
			if strings.Contains(err, "Error reading") {
				w.WriteHeader(http.StatusNotFound)
				tmpl.Execute(w, struct {
					Issue string
					Code  int
				}{Issue: err, Code: http.StatusNotFound})
			}
			if strings.Contains(err, "modified") {
				w.WriteHeader(http.StatusInternalServerError)
				tmpl.Execute(w, struct {
					Issue string
					Code  int
				}{Issue: err, Code: http.StatusInternalServerError})
			}

			// If no error print ascii-art below form on submitForm.html
		} else {

			// Render the form with the ASCII art and download button
			tmpl = template.Must(template.ParseFiles("templates/submitForm.html"))
			tmpl.Execute(w, struct {
				AsciiArt string
				Input    string
			}{AsciiArt: output, Input: inputStr})
		}
	}
}

func DownloadArtHandler(w http.ResponseWriter, r *http.Request) {
	contentLength := len([]byte(output)) // Retrieve size of file
	contentType := "plain/text"          // save content type

	// Set the appropriate headers to attachment to force download
	w.Header().Set("Content-Disposition", "attachment; filename=download-art.txt")
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Length", strconv.Itoa(contentLength))
	w.Write([]byte(output))
}

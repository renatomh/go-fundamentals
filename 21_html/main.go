package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Serving HTML pages

// Declaring the template files to be used in the app
var templates *template.Template

type user struct {
	Name  string
	Email string
}

func main() {

	// This will parse all HTML files
	templates = template.Must(template.ParseGlob("*.html"))

	// Route to serve the HTML page
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		// Creating a new user
		u := user{"The Hobbits", "the.hobbits@isengard.com"}

		// Executing the template to be returend to the request
		// The last parameter is used to provide data to populate the HTML template
		templates.ExecuteTemplate(w, "home.html", u)
	})

	// Listening requests on specified port
	fmt.Println("Listening on port 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}

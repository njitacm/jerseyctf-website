package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	L "github.com/njitacm/jerseyctf-website/src/lib"
)

// CONSTANT
const PORT = 9990

func index(w http.ResponseWriter, r *http.Request) {
	// This ensures that all *.html within the templates/ dir and their
	// specific template are received before called
	tmpl := template.Must(template.ParseGlob("templates/*.html"))

	// Header Template in Layout.html
	tmpl.ExecuteTemplate(w, "header", nil)

	// Nav Bar Template in Layout.html
	//navBar := []string{"Registration", "Speakers", "Sponsors","FAQ", "Recognitions", "Resources"}
	navBar := []string{"Registration","Sponsors","FAQ","Resources"}
	tmpl.ExecuteTemplate(w, "navNBody", navBar)

	// Infographic in Infographic.html
	tmpl.ExecuteTemplate(w, "infographic", nil)

	// registration in registration.html
	tmpl.ExecuteTemplate(w, "registration", nil)

	// schedule in schedule.html
	//tmpl.ExecuteTemplate(w, "schedule", nil)

	// speakers in speakers.html
	//L.Speaker(w, tmpl)

	// sponsors in sponsors.html
	sponsors := L.GetSponsors()
	tmpl.ExecuteTemplate(w, "sponsors", sponsors)

	// faq in faq.html
	faq := L.GetFaq()
	tmpl.ExecuteTemplate(w, "faq", faq)

	// recongitions in recognitions.html
	//L.Recognition(w, tmpl)

	// writeups in writeups.html
	resources := L.GetResources()
	tmpl.ExecuteTemplate(w, "resources", resources)

	// Footer Template in Layout.html
	tmpl.ExecuteTemplate(w, "footer", nil)

}

func main() {
	port := strconv.Itoa(PORT)

	fmt.Printf("Starting server on http://localhost:%s\n", port)

	// Web Server starts at this directory
	fs := http.FileServer(http.Dir("."))

	// Puts everything from File Server into a /assets/ directory
	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	http.HandleFunc("/", index)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

package main

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	L "github.com/njitacm/jerseyctf-registration-site/src/lib"
)

// CONSTANT
const (
	PORT = 9990
)

func index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseGlob("templates/*.html"))

	// Header Template in Layout.html
	tmpl.ExecuteTemplate(w, "header", nil)

	// Nav Bar Template in Layout.html
	navBar := []string{"registration", "schedule", "speakers", "sponsors", "faq", "recognitions"}
	tmpl.ExecuteTemplate(w, "navNBody", navBar)

	// Infographic in Infographic.html
	tmpl.ExecuteTemplate(w, "infographic", nil)

	// registration in registration.html
	tmpl.ExecuteTemplate(w, "registration", nil)

	// schedule in schedule.html
	tmpl.ExecuteTemplate(w, "schedule", nil)

	// speakers in speakers.html

	// sponsors in sponsors.html

	// faq in faq.html
	faq := L.GetFaq()

	tmpl.ExecuteTemplate(w, "faq", faq)

	// recongitions in recognitions.html

	// Footer Template in Layout.html
	tmpl.ExecuteTemplate(w, "footer", nil)

}

func main() {
	port := strconv.Itoa(PORT)

	fs := http.FileServer(http.Dir("."))

	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	http.HandleFunc("/", index)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

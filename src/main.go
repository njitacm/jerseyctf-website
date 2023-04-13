package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
	L "github.com/njitacm/jerseyctf-website/src/lib" //Not the current Repo, but everything works, look into if changes should be made
)

// CONSTANT
const PORT = 9990

func index(w http.ResponseWriter, r *http.Request) {
	// This ensures that all *.html within the templates/ dir and their
	// specific template are received before called
	tmpl := template.Must(template.ParseGlob("templates/*.html"))

	tmpl.ExecuteTemplate(w, "header", nil)

	//The sections that load are based off of what strings are in this list
	navBar := []string{ "Registration","Schedule","Speakers","Sponsors","FAQ","Resources"} // "Registration","Schedule","Speakers","Sponsors","Recognitions","FAQ","Resources"
	tmpl.ExecuteTemplate(w, "navNBody", navBar)

	tmpl.ExecuteTemplate(w, "infographic", nil)

	for _, value := range navBar {
		switch value {
			case "Registration":
				tmpl.ExecuteTemplate(w, "registration", nil)

			case "Schedule":
				tmpl.ExecuteTemplate(w, "schedule", nil)

			case "Speakers":
				L.Speaker(w, tmpl)

			case "Sponsors":
				sponsors := L.GetSponsors()
				tmpl.ExecuteTemplate(w, "sponsors", sponsors)

			case "FAQ":
				faq := L.GetFaq()
				tmpl.ExecuteTemplate(w, "faq", faq)

			case "Recognitions":
				L.Recognition(w, tmpl)

			case "Resources":
				resources := L.GetResources()
				tmpl.ExecuteTemplate(w, "resources", resources)
		}
	}
	tmpl.ExecuteTemplate(w, "footer", nil)

}

func ccpp(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseGlob("templates/*.html"))

	tmpl.ExecuteTemplate(w, "header", nil)

	navBar := []string{"Code of Conduct", "Privacy Policy"}
	tmpl.ExecuteTemplate(w, "navNBody", navBar)

	tmpl.ExecuteTemplate(w, "code-of-conduct", nil)

	tmpl.ExecuteTemplate(w, "privacy-policy", nil)

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
	http.HandleFunc("/ccpp", ccpp)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

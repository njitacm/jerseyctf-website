package main

import (
	"log"
	"net/http"
	"strconv"
	"text/template"
)

// CONSTANT
const (
	PORT = 9990
)

func index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseGlob("templates/*.html"))

	tmpl.ExecuteTemplate(w, "header", nil)
	tmpl.ExecuteTemplate(w, "footer", nil)

}

func main() {
	port := strconv.Itoa(PORT)

	fs := http.FileServer(http.Dir("."))

	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	http.HandleFunc("/", index)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

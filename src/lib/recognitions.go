package lib

import (
	"net/http"
	"text/template"
)

type Person struct {
	Name   string
	Sector string
}

// Copy and Paste Add
// {
// 	Name: "",
// 	Sector: "",
// },

func GetRecognizedPeople() []Person {
	return []Person{
		{
			Name:   "Ilan Ponimansky",
			Sector: "Engineering",
		},
		{
			Name:   "Robert Bruder",
			Sector: "Engineering",
		},
		{
			Name:   "Amanda Galante",
			Sector: "Engineering",
		},
		{
			Name:   "George Gindler",
			Sector: "Engineering",
		},
		{
			Name:   "Abdel Sy Fane",
			Sector: "General Organization",
		},
		{
			Name:   "Toni Diaz",
			Sector: "Marketing",
		},
		{
			Name:   "Dina Anello",
			Sector: "Sponsorships",
		},
		{
			Name:   "NJIT ACM E-Board",
			Sector: "General Organization",
		},
	}
}

// Handles the Recognition Tab and Writes to Website
func Recognition(w http.ResponseWriter, tpl *template.Template) {
	recognitions := GetRecognizedPeople()

	tpl.ExecuteTemplate(w, "recognition-start", nil)

	var iter int8

	for idx, person := range recognitions {

		iter++

		j := idx + 1
		if j%4 == 1 {
			tpl.ExecuteTemplate(w, "recognition-card-group-start", nil)
		}

		tpl.ExecuteTemplate(w, "recognition-card", person)

		if iter%4 == 0 {
			tpl.ExecuteTemplate(w, "recognition-div-end", nil)
		}
	}

	tpl.ExecuteTemplate(w, "recognition-div-end", nil)
}

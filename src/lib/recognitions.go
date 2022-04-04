package lib

import (
	"net/http"
	"text/template"
)

type Instit struct {
	Name   string
	People []string
}


func GetRecognizedPeople() []Instit {
	return []Instit{
		{
			Name:"NJIT YWCC Administration and Students",
			People:[]string{"Dina","Andres", "Tensei", "Robert", "Nishaant", "Andrew", "George", "Keith", "Sreya"," SCI Principal Investigators", "Logan", "David", "Christian", "Cade", "Alexander", "Massa", "Joseph", "Ethan", "Anthony" },
		},
		{
			Name:"NJCCIC Engineers and Interns",
			People: []string{"Rob", "Mandy", "Kevin", "Penelope", "Edward"},
		},
		{
			Name:"RUSEC Members",
			People: []string{"Rajat", "Philip", "Yousef"},
		},
		{
			Name:"FRSecure Engineers",
			People: []string{"Eric", "Matt"},
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

		if len(person.People) <10 {
			tpl.ExecuteTemplate(w, "recognition-person-list", person)
		} else{
			per1:=person
			per2:=person
			per1.People=per1.People[:10]
			per2.People=per2.People[10:]
			tpl.ExecuteTemplate(w, "recognition-person-list", per1)
			tpl.ExecuteTemplate(w, "recognition-person-list", per2)
		}
		for i:=0;i<3;i++ {
			tpl.ExecuteTemplate(w, "recognition-div-end", nil)
		}

		if iter%4 == 0 {
			tpl.ExecuteTemplate(w, "recognition-div-end", nil)
		}
	}

	tpl.ExecuteTemplate(w, "recognition-div-end", nil)

}

package lib

import (
	"net/http"
	"text/template"
)

// Segment Number refers to how one wishes to break down
// cards
const (
	SEGMENT_NUM = 1
)

type SpeakerStruct struct {
	Pic      string
	Name     string
	Position string
	Talk     string
	Time     string
}

// Copy and Paste Add
// {
// 	Pic: "",
// 	Name: "",
// 	Position: "",
// 	Role: "",
// 	LinkedIn: "",
// },

func GetSpeakers() []SpeakerStruct {
	return []SpeakerStruct{
		{
			Pic:      "personPlaceholder.png",
			Name:     "James Perry",
			Position: "Security Lead at AWS",
			Talk:     "TBD",
			Time:     "",
		},
		{
			Pic:      "personPlaceholder.png",
			Name:     "TBD",
			Position: "",
			Talk:     "",
			Time:     "",
		},
		{
			Pic:      "personPlaceholder.png",
			Name:     "TBD",
			Position: "",
			Talk:     "",
			Time:     "",
		},
	}
}

// --------------------------------------------

// 1- 5 inclusive
// All on the same line
// 6 - 12 inclusive
// Being equal or bottom line has more than top line
// >12
// More Dynamic Grid like in Recognitions

// Handles Speaker Tab and writes to website
func Speaker(w http.ResponseWriter, tpl *template.Template) {
	speakers := GetSpeakers()

	tpl.ExecuteTemplate(w, "speaker-start", nil)

	tpl.ExecuteTemplate(w, "speaker-card-group-start", nil)

	for _, speaker := range speakers {
		tpl.ExecuteTemplate(w, "speaker-card", speaker)
	}

	tpl.ExecuteTemplate(w, "speaker-div-end", nil)

	tpl.ExecuteTemplate(w, "speaker-div-end", nil)

}

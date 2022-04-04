package lib

import (
	"net/http"
	"text/template"
)

// Segment Number refers to how one wishes to break down
// cards
const (
	SEGMENT_NUM = 5
)

type SpeakerStruct struct {
	Pic      string
	Name     string
	Position string
	Role     string
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
			Pic:      "MichaelGeraghty.jpg",
			Name:     "Michael Geraghty",
			Position: "CISO, State of New Jersey & Director | NJCCIC",//CISO, State of New Jersey & Director, 
			Role:     "Keynote Speaker",
		},
		{
			Pic:      "",
			Name:     "Jon Taylor",
			Position: "Senior Security Manager and Principal Consultant | Accenture",
			Role:     "Keynote Speaker",
		},{
			Pic:      "",
			Name:     "Brian Herron",
			Position: "Supervisory Special Agent | FBI", 
			Role:     "Keynote Speaker",
		},
		{
			Pic:      "",
			Name:     "Kevin McKenzie",
			Position: "Cyber Security Analyst | NJCCIC", 
			Role:     "Keynote Speaker",
		},
		{
			Pic:      "",
			Name:     "Max Saltonstall",
			Position: "(Pre-Recorded) IT Technical Director | Google", 
			Role:     "Keynote Speaker",
		},
		{
			Pic:      "",
			Name:     "Donnie Rodgers",
			Position: "Dynamic Site Security Analyst | PlainDilemma", 
			Role:     "Keynote Speaker",
		},
		{
			Pic:      "",
			Name:     "Debbi Blyth",
			Position: "(Pre-Recorded) Executive Public Sector Strategist | CrowdStrike", 
			Role:     "Keynote Speaker",
		},
		{
			Pic:      "",
			Name:     "Oscar Minks",
			Position: "Chief Technology Officer | FRSecure", 
			Role:     "Keynote Speaker",
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

	length := len(speakers)

	// Creates a single line of speakers
	if length < 6 {
		tpl.ExecuteTemplate(w, "speaker-card-group-start", nil)

		for _, speaker := range speakers {
			tpl.ExecuteTemplate(w, "speaker-card", speaker)
		}

		tpl.ExecuteTemplate(w, "speaker-div-end", nil)
	}

	// Creates 2 card groups, split by using some integer division
	if length <= 12 && length >= 6 {

		// First Group of Speakers (Speakers[0: length/2])
		tpl.ExecuteTemplate(w, "speaker-card-group-start", nil)

		for _, speaker := range speakers[0 : length/2] {
			tpl.ExecuteTemplate(w, "speaker-card", speaker)
		}

		tpl.ExecuteTemplate(w, "speaker-div-end", nil)

		// Second Group of Speakers (Speakers[length/2: length])
		tpl.ExecuteTemplate(w, "speaker-card-group-start", nil)

		for _, speaker := range speakers[length/2 : length] {
			tpl.ExecuteTemplate(w, "speaker-card", speaker)
		}

		tpl.ExecuteTemplate(w, "speaker-div-end", nil)

	}

	// Creates a dynamic grid
	if length > 12 {
		var iter int

		for idx, speaker := range speakers {

			iter++

			j := idx + 1
			if j%SEGMENT_NUM == 1 {
				tpl.ExecuteTemplate(w, "speaker-card-group-start", nil)
			}

			tpl.ExecuteTemplate(w, "speaker-card", speaker)

			if iter%SEGMENT_NUM == 0 {
				tpl.ExecuteTemplate(w, "speaker-div-end", nil)
			}
		}
	}

	tpl.ExecuteTemplate(w, "speaker-div-end", nil)

}

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
	Talk     string
	Time	 string
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
			Pic:      "JohnHammond.png",
			Name:     "John Hammond",
			Position: "Cybersecurity Researcher, Educator & Content Creator", 
			Talk:     "Cybersecurity Shop Talk",
			Time:	  "Saturday 4/15 (1:00 P.M. EDT)",
		},
		{
			Pic:      "personPlaceHolder.png",
			Name:     "NJCCIC Speaker #1",
			Position: "", 
			Talk:     "Behind the Scenes of Cyber Defense: A Look into the SOC",
			Time:	  "Saturday 4/15 (2:00 P.M. EDT)",
		},
		{
			Pic:      "personPlaceHolder.png",
			Name:     "NJCCIC Speaker #2",
			Position: "", 
			Talk:     "Two-Factor Authentication: Not as Secure as You Think",
			Time:	  "Saturday 4/15 (3:00 P.M. EDT)",
		},
		{
			Pic:      "personPlaceHolder.png",
			Name:     "Ilan Ponimansky",
			Position: "Staff Cloud Security Engineer at Block", 
			Talk:     "",
			Time:	  "Saturday 4/15 (4:00 P.M. EDT)",
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


	//TODO:Simplify the code, it now works with rowcol, no need for this nonesense
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

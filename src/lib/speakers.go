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
			Pic:      "CelinePedalino.jpg",
			Name:     "Celine Pedalino",
			Position: "SOC Analyst at NJCCIC", 
			Talk:     "Behind the Scenes of Cyber Defense: A Look into the SOC",
			Time:	  "Saturday 4/15 (2:00 P.M. EDT)",
		},
		{
			Pic:      "SeonukKim.jpg",
			Name:     "Seonuk Kim",
			Position: "SOC Analyst at NJCCIC", 
			Talk:     "Behind the Scenes of Cyber Defense: A Look into the SOC",
			Time:	  "Saturday 4/15 (2:00 P.M. EDT)",
		},
		{
			Pic:      "TrentMeyers.jpg",
			Name:     "Trent Meyers",
			Position: "SOC Analyst at NJCCIC", 
			Talk:     "Behind the Scenes of Cyber Defense: A Look into the SOC",
			Time:	  "Saturday 4/15 (2:00 P.M. EDT)",
		},
		{
			Pic:      "SwathiParthibha.jpg",
			Name:     "Swathi Parthibha",
			Position: "Security Analyst at NJCCIC", 
			Talk:     "Behind the Scenes of Cyber Defense: A Look into the SOC",
			Time:	  "Saturday 4/15 (2:00 P.M. EDT)",
		},
		{
			Pic:      "personPlaceholder.png",
			Name:     "Andrew Garcia",
			Position: "Security Analyst at NJCCIC", 
			Talk:     "Two-Factor Authentication: Not as Secure as You Think",
			Time:	  "Saturday 4/15 (3:00 P.M. EDT)",
		},
		{
			Pic:      "IlanPonimansky.jpg",
			Name:     "Ilan Ponimansky",
			Position: "Staff Cloud Security Engineer at Block", 
			Talk:     "Acing your Cloud Security Interviews AMA",
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



	tpl.ExecuteTemplate(w, "speaker-card-group-start", nil)

	for _, speaker := range speakers {
		tpl.ExecuteTemplate(w, "speaker-card", speaker)
	}

	tpl.ExecuteTemplate(w, "speaker-div-end", nil)


	tpl.ExecuteTemplate(w, "speaker-div-end", nil)

}

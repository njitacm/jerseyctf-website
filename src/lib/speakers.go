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
		{},
		/*
			{
				Pic:      "JohnHammond.png",
				Name:     "John Hammond",
				Position: "Cybersecurity Researcher, Educator & Content Creator",
				Talk:     "Cybersecurity Shop Talk",
				Time:	  "",
			},
			{
				Pic:      "CelinePedalino.jpg",
				Name:     "Celine Pedalino",
				Position: "SOC Analyst at the NJCCIC",
				Talk:     "Behind the Scenes of Cyber Defense: A Look into the SOC",
				Time:	  "",
			},
			{
				Pic:      "SeonukKim.jpg",
				Name:     "Seonuk Kim",
				Position: "SOC Analyst at the NJCCIC",
				Talk:     "Behind the Scenes of Cyber Defense: A Look into the SOC",
				Time:	  "",
			},
			{
				Pic:      "TrentMeyers.jpg",
				Name:     "Trent Meyers",
				Position: "SOC Analyst at the NJCCIC",
				Talk:     "Behind the Scenes of Cyber Defense: A Look into the SOC",
				Time:	  "",
			},
			{
				Pic:      "SwathiParthibha.jpg",
				Name:     "Swathi Parthibha",
				Position: "Security Analyst at the NJCCIC",
				Talk:     "Behind the Scenes of Cyber Defense: A Look into the SOC",
				Time:	  "",
			},
			{
				Pic:      "personPlaceholder.png",
				Name:     "Andrew Garcia",
				Position: "Security Analyst at the NJCCIC",
				Talk:     "Two-Factor Authentication: Not as Secure as You Think",
				Time:	  "",
			},
			{
				Pic:      "IlanPonimansky.jpg",
				Name:     "Ilan Ponimansky",
				Position: "Staff Cloud Security Engineer at Block",
				Talk:     "Acing your Cloud Security Interviews AMA",
				Time:	  "",
			},
			{
				Pic:      "personPlaceholder.png",
				Name:     "Kevin Conklin",
				Position: "Cyber National Security Supervisory Special Agent at FBI Newark",
				Talk:     "Inside the FBI Cyber Program",
				Time:	  "",
			},
		*/
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

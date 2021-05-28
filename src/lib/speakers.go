package lib

import (
	"net/http"
	"text/template"
)

type SpeakerStruct struct {
	Pic      string
	Name     string
	Position string
	Role     string
	LinkedIn string
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
			Position: "CISO, State of New Jersey & Director, NJCCIC",
			Role:     "Keynote Speaker",
			LinkedIn: "michael-geraghty-13139820",
		},
		{
			Pic:      "AbdelSyFane.jpg",
			Name:     "Abdel Sy Fane",
			Position: "President, CSNP",
			Role:     "DevSecOps",
			LinkedIn: "abdelsyfane",
		},
		{
			Pic:      "GabrielleBotbol.jpg",
			Name:     "Gabrielle Botbol",
			Position: "Pentester, CSbyGB",
			Role:     "Solving Web Security Vulnerabilities with Pentesting",
			LinkedIn: "gabriellebotbol",
		},
		{
			Pic:      "IlanPonimansky.jpg",
			Name:     "Ilan Ponimansky",
			Position: "Sr. Cloud Security Consultant, ScaleSec",
			Role:     "Cloud Security",
			LinkedIn: "ilanponimansky",
		},
		{
			Pic:      "JohnJackson.jpg",
			Name:     "John Jackson",
			Position: "Sr. App. Security Engineer, Shutterstock & Founder, Sakura Samurai",
			Role:     "Day Zero to 0day, Offensive Security Theory",
			LinkedIn: "johnjhacking",
		},
		{
			Pic:      "JonHelmus.jpg",
			Name:     "Jon Helmus",
			Position: "Offensive Security Educator and Penetration Tester",
			Role:     "Functional Testing: A New Era of Pentesting",
			LinkedIn: "jon-helmus-474146103",
		},
		{
			Pic:      "SethKirschner.jpg",
			Name:     "Seth Kirschner",
			Position: "Security Engineer & Assistant Vice President, MUFG Securities Americas Inc.",
			Role:     "Consulting within Big4 Cyber",
			LinkedIn: "sethkirschner",
		},
		{
			Pic:      "WilliamPrice.png",
			Name:     "William Price",
			Position: "President, CyberX",
			Role:     "Hacking Humans: Using OSINT to Put Together Social Engineering Scenarios that Work",
			LinkedIn: "cyberxcyecurityspecialistbers",
		},
		{
			Pic:      "JoannaHuisman.jpg",
			Name:     "Joanna Huisman",
			Position: "Senior Vice President of Strategic Insights and Research, KnowBe4",
			Role:     "Driving a Security Culture",
			LinkedIn: "joannahuisman",
		},
		{
			Pic:      "JosephRusso.jpg",
			Name:     "Joseph Russo",
			Position: "Bureau Chief for Cybersecurity Engineering and Operations, NJCCIC",
			Role:     "Security Telemetry Aggregation and Analysis",
			LinkedIn: "joe-russo-79358b117",
		},
		{
			Pic:      "HopeWalker.jpg",
			Name:     "Hope Walker",
			Position: "Consultant, SpecterOps",
			Role:     "Icemoon's Fantastical Intro to Offensive Security",
			LinkedIn: "hope-walker-82398984",
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
	if length <= 12 {

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

		segmentNum := 5
		var iter int

		for idx, speaker := range speakers {

			iter++

			j := idx + 1
			if j%segmentNum == 1 {
				tpl.ExecuteTemplate(w, "speaker-card-group-start", nil)
			}

			tpl.ExecuteTemplate(w, "speaker-card", speaker)

			if iter%segmentNum == 0 {
				tpl.ExecuteTemplate(w, "speaker-div-end", nil)
			}
		}
	}

	tpl.ExecuteTemplate(w, "speaker-div-end", nil)

}

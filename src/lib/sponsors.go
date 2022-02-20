package lib

type Sponsor struct {
	HREF   string
	Source string
	ALT    string
	Width  string
}

type Sponsors struct {
	CISO   []Sponsor
	Manager []Sponsor
	Analyst []Sponsor
}

/*
For the Sponsors You have to segment the the categories between:
- CISO
- Security Manager
- Security Analyst

Each Category Segment belongs to a separate component so unfortunately
they can't combined

*/

func GetSponsors() Sponsors {
	ciso := []Sponsor{

		{
			HREF:   "https://www.gemini.com/",
			Source: "gemini_logo.png",
			ALT:    "Gemini",
			Width:  "350",
		},
	}

	manager := []Sponsor{
		{
			HREF:   "https://www.paloaltonetworks.com/",
			Source: "palo_alto_logo.png",
			ALT:    "Palo Alto Networks",
			Width:  "350",
		},
		{
			HREF:   "https://cloud.google.com/",
			Source: "google_logo.gif",
			ALT:    "Google",
			Width:  "200",
		},
	}

	analyst := []Sponsor{
		{
			HREF:   "https://sci.njit.edu/",
			Source: "njit_sci_logo.png",
			ALT:    "NJIT SCI",
			Width:  "300",
		},
	}

	return Sponsors{
		CISO:   ciso,
		Manager: manager,
		Analyst: analyst,
	}
}

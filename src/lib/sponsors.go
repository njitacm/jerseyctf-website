package lib

type Sponsor struct {
	HREF   string
	Source string
	ALT    string
	Width  string
}

type Sponsors struct {
	Gold   []Sponsor
	Silver []Sponsor
	Bronze []Sponsor
}

/*
For the Sponsors You have to segment the the categories between:
- Gold
- Silver
- Bronze

Each Category Segment belongs to a separate component so unfortunately
they can't combined

*/

func GetSponsors() Sponsors {
	gold := []Sponsor{

		{
			HREF:   "https://www.cyber.nj.gov/",
			Source: "gemini_logo.png",
			ALT:    "Gemini",
			Width:  "350",
		},
	}

	silver := []Sponsor{
		{
			HREF:   "https://www.digitalocean.com/",
			Source: "palo_alto_logo.png",
			ALT:    "Palo Alto Networks",
			Width:  "350",
		},
		{
			HREF:   "https://www.knowbe4.com/",
			Source: "google_logo.gif",
			ALT:    "Google",
			Width:  "200",
		},
	}

	bronze := []Sponsor{
		{
			HREF:   "https://www.securityinnovation.com/",
			Source: "njit_sci_logo.png",
			ALT:    "NJIT SCI",
			Width:  "300",
		},
	}

	return Sponsors{
		Gold:   gold,
		Silver: silver,
		Bronze: bronze,
	}
}

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
			HREF:   "https://specterops.io/",
			Source: "Specter_ops_logo.png",
			ALT:    "Specter Ops",
			Width:  "400",
		},
		{
			HREF:   "https://www.cyber.nj.gov/",
			Source: "NJCCIC_logo.png",
			ALT:    "NJCCIC",
			Width:  "200",
		},
	}

	silver := []Sponsor{
		{
			HREF:   "https://www.digitalocean.com/",
			Source: "DO_Logo.png",
			ALT:    "Digital Ocean",
			Width:  "250",
		},
		{
			HREF:   "https://www.knowbe4.com/",
			Source: "KnowBe4_logo.png",
			ALT:    "KnowBe4",
			Width:  "300",
		},
	}

	bronze := []Sponsor{
		{
			HREF:   "https://www.securityinnovation.com/",
			Source: "security_innovation_logo.png",
			ALT:    "Security Innovation",
			Width:  "300",
		},
		{
			HREF:   "https://contrastsecurity.com/",
			Source: "Contrast_Security_logo.png",
			ALT:    "Contrast Security",
			Width:  "300",
		},
	}

	return Sponsors{
		Gold:   gold,
		Silver: silver,
		Bronze: bronze,
	}
}

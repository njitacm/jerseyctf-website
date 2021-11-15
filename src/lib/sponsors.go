package lib

type Sponsor struct {
	HREF   string
	Source string
	ALT    string
	Type   string
	Width  string
}

func GetSponsors() []Sponsor {
	return []Sponsor{
		{
			HREF:   "https://specterops.io/",
			Source: "Specter_ops_logo.png",
			ALT:    "Specter Ops",
			Type:   "gold",
			Width:  "400",
		},
		{
			HREF:   "https://www.cyber.nj.gov/",
			Source: "NJCCIC_logo.png",
			ALT:    "NJCCIC",
			Type:   "gold",
			Width:  "200",
		},
		{
			HREF:   "https://www.digitalocean.com/",
			Source: "DO_Logo.png",
			ALT:    "Digital Ocean",
			Type:   "silver",
			Width:  "250",
		},
		{
			HREF:   "https://www.knowbe4.com/",
			Source: "KnowBe4_logo.png",
			ALT:    "KnowBe4",
			Type:   "silver",
			Width:  "300",
		},
		{
			HREF:   "https://www.securityinnovation.com/",
			Source: "security_innovation_logo.png",
			ALT:    "Security Innovation",
			Type:   "bronze",
			Width:  "300",
		},
		{
			HREF:   "https://contrastsecurity.com/",
			Source: "Contrast_Security_logo.png",
			ALT:    "Contrast Security",
			Type:   "bronze",
			Width:  "300",
		},
	}
}

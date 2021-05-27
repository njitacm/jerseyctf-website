package lib

type Sponsor struct {
	HREF    string
	SrcFile string
	ALT		string
}


type Sponsors struct {
	Gold   []Sponsor
	Silver []Sponsor
	Bronze []Sponsor
}

func GetSponsors() Sponsors {
	gold := []Sponsor{
		{
			HREF: "https://specterops.io/",
			SrcFile: "Specter_ops_logo.png",
			ALT: "Specter Ops",
		},
		{
			HREF: "https://www.cyber.nj.gov/",
			SrcFile: "NJCCIC_logo.png",
			ALT: "NJCCIC",
		},
	}

	silver := []Sponsor{
		{
			HREF: "https://www.digitalocean.com/",
			SrcFile: "DO_Logo.png",
			ALT: "Digital Ocean",
		},
		{
			HREF: "https://www.knowbe4.com/",
			SrcFile: "KnowBe4_logo.png",
			ALT: "KnowBe4",
		},
	}

	bronze := []Sponsor{
		{
			HREF: "https://www.securityinnovation.com/",
			SrcFile: "security_innovation_logo.png",
			ALT: "Security Innovation",
		},
		{
			HREF: "https://contrastsecurity.com/",
			SrcFile: "Contrast_Security_logo.png",
			ALT: "Contrast Security",
		},
	}


	return Sponsors{
		Gold: gold,
		Silver: silver,
		Bronze: bronze,
	}
}
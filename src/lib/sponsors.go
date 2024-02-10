package lib

type Tier struct {
	Spons    []Sponsor
	TierName string
	CSSName  string
}

type Sponsor struct {
	HREF   string
	Source string
	ALT    string
	Width  string
}

type Sponsors struct {
	AllTiers []Tier
}

func GetSponsors() Sponsors {
	title := []Sponsor{
		{
			HREF:   "https://www.cyber.nj.gov/",
			Source: "NJCCIC_logo.png",
			ALT:    "NJCCIC",
			Width:  "200",
		},
	}

	ciso := []Sponsor{
		{
			HREF:   "https://www.merck.com/",
			Source: "Merck.png",
			ALT:    "Merck",
			Width:  "200",
		},
	}

	manager := []Sponsor{}

	analyst := []Sponsor{
		{
			HREF:   "https://aws.amazon.com/",
			Source: "AWS_Logo.png",
			ALT:    "AWS",
			Width:  "200",
		},
		{
			HREF:   "https://www.crowdstrike.com/en-us/",
			Source: "crowdstrike.png",
			ALT:    "CrowdStrike",
			Width:  "200",
		},
		{
			HREF:   "https://www.isaca.org/",
			Source: "ISACA_logo_RGB.jpg",
			ALT:    "ISACA",
			Width:  "200",
		},
	}

	tiers := []Tier{
		{
			Spons:    title,
			TierName: "Title Sponsor",
			CSSName:  "sponsor-title",
		},

		{
			Spons:    ciso,
			TierName: "CISO Sponsor",
			CSSName:  "sponsor-ciso",
		},
		{
			Spons:    manager,
			TierName: "Security Manager Sponsor",
			CSSName:  "sponsor-manager",
		},
		{
			Spons:    analyst,
			TierName: "Security Analyst Sponsor",
			CSSName:  "sponsor-analyst",
		},
	}

	return Sponsors{
		AllTiers: tiers,
	}
}

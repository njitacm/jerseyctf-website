package lib

type Tier struct {
	Spons []Sponsor
	TierName string
	CSSName string
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
	}
	
	manager := []Sponsor{
	}

	analyst := []Sponsor{
		{
			HREF:   "https://www.offensive-security.com/",
			Source: "offensive_security_logo.png",
			ALT:    "Offensive Security",
			Width:  "400",
		},
	}
		
	tiers := []Tier{
		{
			Spons: title,
			TierName: "Title Sponsor",
			CSSName: "sponsor-title",
		},
		
		{
			Spons: ciso,
			TierName: "CISO Sponsor",
			CSSName: "sponsor-ciso",
		},
		{
			Spons: manager,
			TierName: "Security Manager Sponsor",
			CSSName: "sponsor-manager",
		},
		{
			Spons: analyst,
			TierName: "Security Analyst Sponsor",
			CSSName: "sponsor-analyst",
		},
	}

	return Sponsors{
		AllTiers: tiers,
	}
}

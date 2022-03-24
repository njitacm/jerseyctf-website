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
	/*
	ciso := []Sponsor{
	}
	*/
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
		{
			HREF:   "https://www.crowdstrike.com/",
			Source: "crowdstrike_logo.png",
			ALT:    "CrowdStrike",
			Width:  "350",
		},
	}

	analyst := []Sponsor{
		{
			HREF:   "https://sci.njit.edu/",
			Source: "njit_sci_logo.png",
			ALT:    "NJIT SCI",
			Width:  "300",
		},
		{
			HREF:   "https://www.lutron.com/en-US/pages/default.aspx",
			Source: "lutron_logo.png",
			ALT:    "Lutron",
			Width:  "300",
		},
	}
		
	tiers := []Tier{
		{
			Spons: title,
			TierName: "Title Sponsor",
			CSSName: "sponsor-title",
		},
		/*
		{
			Spons: ciso,
			TierName: "CISO Sponsor",
			CSSName: "sponsor-ciso",
		},*/
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

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
			HREF:   "https://www.paloaltonetworks.com/",
			Source: "palo_alto_logo.png",
			ALT:    "Palo Alto",
			Width:  "400",
		},
		{
			HREF:   "https://cloud.google.com/",
			Source: "google_logo.gif",
			ALT:    "Google",
			Width:  "200",
		},
		{
			HREF:   "https://www.github.com/",
			Source: "github_logo.png",
			ALT:    "GitHub",
			Width:  "150",
		},
		{
			HREF:   "https://www.offsec.com/",
			Source: "offsec_logo.jpg",
			ALT:    "OffSec",
			Width:  "400",
		},
		{
			HREF:   "https://www.specterops.io",
			Source: "specterops_logo.png",
			ALT:    "SpecterOps",
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

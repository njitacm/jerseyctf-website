package lib

type Faq struct {
	Question string
	Answer   string
}

func GetFaq() []Faq {

	return []Faq{
		{
			Question: "What is JerseyCTF?",
			Answer:   "JerseyCTF is a beginner-friendly CTF, meant to get people interested in security. The competition is open to anyone, so we will make sure to provide a range of challenge difficulties.",
		},
		{
			Question: "When and where is JerseyCTF?",
			Answer:   "JerseyCTF will be virtually hosted on Discord. The competition will be on April 10th-11th, 2021!",
		},
		{
			Question: "Who can sign up to compete?",
			Answer:   "Students, beginners, and professionals are all welcome to compete as long as they are over the age of 18.",
		},
		{
			Question: "Will there be prizes?",
			Answer:   "Yes! There are 2 leaderboard divisions: Student and Open. Prizes will be awarded in both divisions for top scoring teams, and anyone who checks-in will receive a complimentary JerseyCTF t-shirt!",
		},
		{
			Question: "How big is JerseyCTF?",
			Answer:   "Space is limited, so we will be capping the competition at a maximum of 125 teams or 500 people.",
		},
		{
			Question: "Do I have to pay to compete in JerseyCTF?",
			Answer:   "Nope, there's no fee to sign up or to attend!",
		},
		{
			Question: "Do I have to register as a team?",
			Answer:   "No! You're welcome to register with your friends and compete with a team, but you can also register by yourself and find a team at the event.",
		},
		{
			Question: "You guys haven't answered my question!",
			Answer:   "Feel free to email us any additional questions and concerns at <a href=\"mailto:jerseyctf@njit.edu\">jerseyctf@njit.edu</a>",
		},
	}
}

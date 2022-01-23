package lib

type Faq struct {
	Question string
	Answer   string
}

func GetFaq() []Faq {

	return []Faq{
		{
			Question: "What is JerseyCTF?",
			Answer:   "JerseyCTF is a beginner-friendly Capture the Flag competition that aims to inspire interest in cybersecurity. Hosted by the <a href=https://njit.acm.org>NJIT ACM</a> organization and <a href=https://sci.njit.edu>NJIT SCI</a> program, it is geared towards students, beginners, and professionals alike. JerseyCTF provides participants with jeopardy-style questions in categories including cryptography, forensics, binary exploitation/reversing, open-source intelligence, and web exploitation.",
		},
		{
			Question: "Who can sign up to compete?",
			Answer:   "Students, beginners, and professionals are all welcome to compete as long as they are over the age of 18.",
		},
		{
			Question: "What do I need to compete?",
			Answer:   "All you need is an Internet-connected device in order to compete. It is completely optional if you want to use ethical hacking/penetration testing Linux distributions like Kali or Parrot OS, but these may help you since there are many pre-installed tools. These distributions are open-source and can be deployed in VMware, VirtualBox, or any other virtualization software.",
		},
		{
			Question: "Will there be prizes?",
			Answer:   "Yes! There are 2 leaderboard divisions: Student and Open. Prizes will be awarded in both divisions for top scoring teams, and anyone who checks-in will receive a complimentary JerseyCTF t-shirt!",
		},
		{
			Question: "How big is JerseyCTF?",
			Answer:   "JerseyCTF had over 600 registrations in 2021, the first year of the competition. We are looking forward to welcoming more registrants and participants in 2022!",
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

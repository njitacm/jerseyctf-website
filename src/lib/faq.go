package lib

type Faq struct {
	Question string
	Answer   string
}

func GetFaq() []Faq {

	return []Faq{
		{
			Question: "What is JerseyCTF?",
			Answer:   "JerseyCTF is a beginner-friendly Capture the Flag competition that aims to inspire interest in cybersecurity. Hosted by the NJIT <a target='_blank' href=https://njit.acm.org>ACM</a> and <a target='_blank' href= https://njiticc.com>NICC</a> organizations and the <a target='_blank' href=https://sci.njit.edu>NJIT SCI</a> program, it is geared towards students, beginners, and professionals alike. JerseyCTF provides participants with jeopardy-style questions in categories including cryptography, forensics, binary exploitation/reversing, open-source intelligence, and web exploitation.  JerseyCTF features a track parallel to the CTF competition, consisting of presentations and panels with renowned speakers from industry and government.",
		},
		{
			Question: "Who can sign up to compete?",
			Answer:   "Students, beginners, and professionals are all welcome to compete. Must be 18+ to attend in-person.",
		},
		{
			Question: "What do I need to compete?",
			Answer:   "All you need is an Internet-connected device in order to compete and watch our tech talks. While optional, it is recommended for you to use penetration testing Linux distributions like Kali or Parrot OS because the many pre-installed tools will help you. These distributions are free, open-source, and can be deployed in VMware, VirtualBox, or any other virtualization software. Check out our <a href=https://ctf.jerseyctf.com/resources target='_blank'> beginner resources page</a> for more recommendations. ",
		},
		{
			Question: "Will there be prizes?",
			Answer:   "Yes! There are 2 leaderboard divisions: Student and Open. Prizes will be awarded in both divisions for top scoring teams and participants residing in the United States will receive a complimentary JerseyCTF t-shirt.",
		},
		{
			Question: "How big is JerseyCTF?",
			Answer:   "JerseyCTF had 1,515 registrations in 2023. We are looking forward to welcoming more registrants and participants in 2024!",
		},
		{
			Question: "Do I have to pay to compete in JerseyCTF?",
			Answer:   "Nope, there's no fee to sign up or to attend!",
		},
		{
			Question: "Do I have to register as a team?",
			Answer:   "No! You're welcome to register with your friends and compete with a team of up to 4 people, but you can also register by yourself or find a team in the Discord server.",
		},
		{
			Question: "You guys haven't answered my question!",
			Answer:   "Feel free to email us any additional questions and concerns at <a target=\"_blank\" href=\"mailto:jerseyctf@njit.edu\">jerseyctf@njit.edu</a>",
		},
	}
}

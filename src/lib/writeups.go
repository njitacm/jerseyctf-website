package lib

type Writeup struct {
	Link string
	Year string
}

// Copy and Paste Add
// {
// 	Link: "",
// 	Year: "",
// },

// Returns an Array of Writeups for usage throughout program
func GetWriteups() []Writeup {
	return []Writeup{
		{
			Link: "https://github.com/njitacm/ctf-challenges/tree/main/writeups",
			Year: "2021",
		},
	}
}

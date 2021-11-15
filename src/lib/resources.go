package lib

type Resource struct {
	Year     string
	Writeups string
	Playlist string
}

// Copy and Paste Add
// {
// 	Year:     "2022",
// 	Writeups: "#",
// 	Playlist: "#",
// },

// Returns an Array of Writeups for usage throughout program
func GetResources() []Resource {
	return []Resource{
		{
			Year:     "2021",
			Writeups: "https://github.com/njitacm/ctf-challenges/tree/main/writeups",
			Playlist: "https://www.youtube.com/playlist?list=PLrcTWWy-esnDYt1niwIETam5s-nljoeD9",
		},
		{
			Year:     "2022",
			Writeups: "#",
			Playlist: "#",
		},
	}
}

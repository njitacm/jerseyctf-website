package lib

type Resource struct {
	Year     string
	Writeups string
	Playlist string
}


// Returns an Array of Writeups for usage throughout program
func GetResources() []Resource {
	return []Resource{
		{
			Year:     "2021",
			Writeups: " https://github.com/njitacm/jerseyctf-2021-challenges/tree/main/writeups",
			Playlist: "https://www.youtube.com/playlist?list=PLrcTWWy-esnDYt1niwIETam5s-nljoeD9",
		},
		{
			Year:     "2022",
			Writeups: "https://github.com/njitacm/jerseyctf-2022-challenges/tree/main/writeups",
			Playlist: "https://youtube.com/playlist?list=PLrcTWWy-esnCuaiEMSj6Bst4phnq-Qg6B",
		},
		{
			Year:"2023",
			Writeups: "https://ctftime.org/event/1908/tasks/",
			Playlist: "https://www.youtube.com/playlist?list=PLrcTWWy-esnDXi3khogdlAgkisL19IM08",
		},
	}
}

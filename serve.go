package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func getAbsLocation() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	ret := path[:index]
	return ret
}

func webserver(port string, mode int8) {
	fmt.Printf("Server on port %s\n", port)

	var fileServer http.Handler

	if mode == 0 {
		fileServer = http.FileServer(http.Dir(".")) // must run within  directory that has html
	}

	if mode == 1 {

		location := getAbsLocation()
		fmt.Printf("Program running at %s\n", location) // must build in order to work

		// change the http.Dir string to the specific place on the machine otherwise it is easy to compromise
		fileServer = http.FileServer(http.Dir(location))
	}

	http.Handle("/", fileServer)

	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	/*
		mode:
			0 --> local  					Must be in directory where html/css is
			1 --> remote/production			Must be an executable to run
	*/

	const (
		PORT = 9990
		mode = 1
	)

	// turns int to string
	port := strconv.Itoa(PORT)

	webserver(port, mode)

}

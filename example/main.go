package main

import (
	"fmt"
	"net/http"

	"github.com/Masterminds/go-fileserver"
)

func main() {

	// Specity a NotFoundHandler to use when no file is found.
	fileserver.NotFoundHandler = func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintln(w, "That page could not be found.")
	}

	// Serve a directory of files.
	dir := http.Dir("./files")
	http.ListenAndServe(":8080", fileserver.FileServer(dir))
}

// func main() {
// 	http.HandleFunc("/", readme)
// 	http.ListenAndServe(":8080", nil)
// }

// func readme(res http.ResponseWriter, req *http.Request) {
// 	fileserver.ServeFile(res, req, "./files/readme.txt")
// }

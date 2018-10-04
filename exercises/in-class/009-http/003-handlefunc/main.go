package main

import (
	"io"
	"net/http"
)

func main() {
	// index
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Index</h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>About</h1>")
}

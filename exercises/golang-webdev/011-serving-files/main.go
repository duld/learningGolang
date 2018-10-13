/*
- Create a five page website with a different image on each page using
	http.ServeFile

- Create a five page website with a different image on each page using
  http.FileServer and http.StripPrefix
*/
package main

import (
	"fmt"
	"net/http"
)

func main() {
	// File server
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	// Routes - serve file
	http.HandleFunc("/", indexHandler) // root - sf
	http.HandleFunc("/sharkdude", sharkHandler)
	http.HandleFunc("/gobeklitepe", gobeklitepeHandler)
	http.HandleFunc("/gobeklitepe/crane-pillar", gobeklitepeCraneHandler)
	http.HandleFunc("/this-is-fine", thisisfineHandler)
	http.HandleFunc("/yosemite", yosemiteHandler)

	// Routes - file server
	http.HandleFunc("/file-server", fileServerHandler) // root - fs
	http.HandleFunc("/file-server/sharkdude", sharkHandlerFS)
	http.HandleFunc("/file-server/gobeklitepe", gobeklitepeHandlerFS)
	http.HandleFunc("/file-server/gobeklitepe/crane-pillar", gobeklitepeCraneHandlerFS)
	http.HandleFunc("/file-server/this-is-fine", thisisfineHandlerFS)
	http.HandleFunc("/file-server/yosemite", yosemiteHandlerFS)

	fmt.Println("Server up and running on port 8000")
	http.ListenAndServe(":8000", nil)
}

// Route Handlers - ServeFile
func indexHandler(w http.ResponseWriter, req *http.Request) {
	page := `
	<h1><a href="/">Serve-file</a></h1>
	<h1><a href="/file-server">File-Server</a></h1>
	<h2><a href="/sharkdude">Sharkdude</a></h2>
	<h2><a href="/gobeklitepe">Gobekli Tepe</a></h2>
	<h2><a href="/gobeklitepe/crane-pillar">Gobekli Tepe: Crane Pillar</a></h2>
	<h2><a href="/this-is-fine">This is fine</a></h2>
	<h2><a href="/yosemite">Yosemite</a></h2>
	`

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, page)
}

func sharkHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./assets/sharkdude.png")
}

func gobeklitepeHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./assets/gobekli_tepe.jpg")
}

func gobeklitepeCraneHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./assets/gobekli_tepe_crane_pillar.jpg")
}

func thisisfineHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./assets/this-is-fine.jpg")
}

func yosemiteHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./assets/yosemite.jpg")
}

// Route Handlers - FileServer
func fileServerHandler(w http.ResponseWriter, req *http.Request) {
	page := `
	<h1><a href="/">Serve-file</a></h1>
	<h1><a href="/file-server">File-Server</a></h1>
	<h2><a href="/file-server/sharkdude">Sharkdude - fs</a></h2>
	<h2><a href="/file-server/gobeklitepe">Gobekli Tepe - fs</a></h2>
	<h2><a href="/file-server/gobeklitepe/crane-pillar">Gobekli Tepe: Crane Pillar - fs</a></h2>
	<h2><a href="/file-server/this-is-fine">This is fine - fs</a></h2>
	<h2><a href="/file-server/yosemite">Yosemite - fs</a></h2>
	`

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, page)
}

func sharkHandlerFS(w http.ResponseWriter, req *http.Request) {
	page := `
		<h1>Suns out fins out</h1>
		<img src="/assets/sharkdude.png">
	`

	// http.ServeFile(w, req, "./assets/sharkdude.png")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, page)
}

func gobeklitepeHandlerFS(w http.ResponseWriter, req *http.Request) {
	page := `
		<h1>Rewriting ancient history one Megalith at a time.</h1>
		<img src="/assets/gobekli_tepe.jpg">
	`

	// http.ServeFile(w, req, "./assets/gobekli_tepe.jpg")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, page)
}

func gobeklitepeCraneHandlerFS(w http.ResponseWriter, req *http.Request) {
	page := `
		<h1>Legs don't work like that.</h1>
		<img src="/assets/gobekli_tepe_crane_pillar.jpg">
	`
	// http.ServeFile(w, req, "./assets/gobekli_tepe_crane_pillar.jpg")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, page)
}

func thisisfineHandlerFS(w http.ResponseWriter, req *http.Request) {
	page := `
		<h1>Keep calm and melt your skin.</h1>
		<img src="/assets/this-is-fine.jpg">
	`
	// http.ServeFile(w, req, "./assets/this-is-fine.jpg")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, page)
}

func yosemiteHandlerFS(w http.ResponseWriter, req *http.Request) {
	page := `
		<h1>When this sucker blows, say goodbye to the West Coast!</h1>
		<img src="/assets/yosemite.jpg">
	`
	// http.ServeFile(w, req, "./assets/yosemite.jpg")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, page)
}

/*
Serve the files in the "starting-files" folder
Use "http.FileServer"
*/
package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)

	// start server
	fmt.Println("Server up and running on port 8000")
	http.ListenAndServe(":8000", nil)
}

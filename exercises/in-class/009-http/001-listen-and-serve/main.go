package main

import (
	"fmt"
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
	fmt.Printf("%v %v %v\n", req.Method, req.URL, req.Proto)
}

func main() {
	http.Handle("/dog", http.HandlerFunc(d))
	http.Handle("/cat", http.HandlerFunc(c))

	http.ListenAndServe(":8080", nil)

}

func serverUp(port string) {
	fmt.Println("Server is up and running on port:", port)
}

package main

import (
	"fmt"
	"log"
	"net/http"
)

type CustomHandler struct{}

func (h CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		h.helloHandler(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (h CustomHandler) helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!") // Send "Hello, World!" as the response
}

func main() {
	handler := CustomHandler{}
	log.Fatal(http.ListenAndServe(":8080", handler)) // Start the server on port 8080
}


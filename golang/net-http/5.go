package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my dynamic website!")
	})
	http.HandleFunc("/my-path-1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Welcome to my path 1")
		// supports only /my-path-2
	})
	http.HandleFunc("/my-path-2/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(401)
		fmt.Fprintf(w, "Welcome to my path 2")
		// supports /my-path-2/abc too
	})

	http.ListenAndServe(":8080", nil)
	// what is this nil?
}

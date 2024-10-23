package main

import (
	"net/http"
)

func myHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", myHandlerFunc)

	mux.HandleFunc("/my-path", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("my path handler is handling"))
	})

	http.ListenAndServe(":8080", mux)
	// previously, we used "DefaultServeMux"
}

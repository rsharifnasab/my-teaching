package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("key-1", "val-1")
		w.WriteHeader(http.StatusAccepted) // or simply 202
		fmt.Fprintf(w, "Welcome to my dynamic website!")
	})

	http.ListenAndServe(":8080", nil)
}

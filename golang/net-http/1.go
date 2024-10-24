package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my dynamic website!\n")
		w.Write([]byte("bye"))
		// what is fprintf? what is w?
		// what is status code? read docs
		// what about content length?
	})

	// listen to addr/port
	http.ListenAndServe(":8080", nil)
}

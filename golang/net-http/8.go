package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Method (GET, POST, etc.) as string
	method := r.Method

	// Request URL
	url := r.URL.String()

	// Protocol (HTTP/1.1, etc.)
	proto := r.Proto

	// Headers
	headers := r.Header
	contentType := headers.Get("Content-Type")

	// Query parameters
	queryParams := r.URL.Query()
	query := queryParams.Get("q")

	// RemoteAddr contains the remote IP address and port of the request
	remoteAddr := r.RemoteAddr

	// Write response to client
	fmt.Fprintf(w, "Request Method: %s\n", method)
	fmt.Fprintf(w, "Request URL: %s\n", url)
	fmt.Fprintf(w, "Request Protocol: %s\n", proto)
	fmt.Fprintf(w, "Request Headers: %v\n", headers)
	fmt.Fprintf(w, "Content type: %v\n", contentType)
	fmt.Fprintf(w, "Remote Address: %s\n", remoteAddr)
	fmt.Fprintf(w, "Query Parameters: %v\n", queryParams)
	fmt.Fprintf(w, "Query: %v\n", query)

	if method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read body", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close() // Always close the body after reading

		fmt.Fprintf(w, "Request Body: %s\n", string(body))
	}
}

func main() {
	// Register the handler function for the root path
	http.HandleFunc("/", handler)

	// Start the HTTP server on port 8080
	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// curl -i "http://localhost:8080/path?queryparam=valueinparam"
// curl -i "http://localhost:8080/path?queryparam=valueinparam" -X POST -d '{"key":"value"}' -H 'content-type: application/json'

//

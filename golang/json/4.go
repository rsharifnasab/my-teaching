package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

type User struct {
	Name       string    `json:"name"`
	Age        int       `json:"age"`
	Email      string    `json:"email"`
	IsVerified bool      `json:"isVerified"`
	DateJoined time.Time `json:"dateJoined"`
}

func main() {
	user := User{
		Name:       "Jane Doe",
		Age:        28,
		Email:      "janedoe@example.com",
		IsVerified: true,
		DateJoined: time.Now(),
	}

	jsonData, err := json.MarshalIndent(user, "", "  ") // MarshalIndent for pretty-printing
	if err != nil {
		log.Fatalf("Error marshalling struct to JSON: %s", err)
	}

	// Print the JSON data to the console
	fmt.Println(string(jsonData))

	// Finally, write the JSON data to a file
	err = ioutil.WriteFile("output.json", jsonData, 0644)
	if err != nil {
		log.Fatalf("Error writing JSON to file: %s", err)
	}
}

/*
// Custom marshalling for the User struct
func (u User) MarshalJSON() ([]byte, error) {
	type Alias User
	return json.Marshal(
		&struct {
			DateJoined string `json:"dateJoined"` // Format DateJoined as a string
			Alias
		}{
			DateJoined: u.DateJoined.Format("2006-01-02"), // Custom date format (YYYY-MM-DD)
			Alias:      (Alias)(u),                        // Embed the original fields
		},
	)
}
*/

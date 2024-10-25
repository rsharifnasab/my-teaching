package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type User struct {
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Email      string `json:"email"`
	IsVerified bool   `json:"isVerified"`
}

func main() {
	user := User{
		Name:       "Jane Doe",
		Age:        28,
		Email:      "janedoe@example.com",
		IsVerified: true,
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

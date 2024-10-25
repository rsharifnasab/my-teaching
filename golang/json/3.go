package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type User struct {
	// private
	key1 string `json:"key1"`

	Arr1 []string `json:"arr1"`

	Nested1 map[string]interface{} `json:"nested1"`
}

func main() {
	file, err := os.Open("a.json")
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	var user User

	err = json.Unmarshal(byteValue, &user)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %s", err)
	}

	fmt.Printf("User: %+v\n", user)
}

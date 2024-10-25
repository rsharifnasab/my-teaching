package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const filePath = "./a.json"

func main() {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %s, error: %v", filePath, err)
	}
	defer file.Close()

	var result map[string]interface{}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&result); err != nil {
		log.Fatalf("Failed to decode JSON: %v", err)
	}

	// Print the result map
	fmt.Printf("Decoded map: %+v\n", result)

	// access fields
	if val, ok := result["key1"].(string); ok {
		fmt.Printf("key1: %s\n", val)
	} else {
		println("key 1 not found")
	}
	if val, ok := result["nested1"].(map[string]interface{}); ok {
		fmt.Printf("nested 1: %s\n", val)
	} else {
		println("nested 1 not found")
	}
}

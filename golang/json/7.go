package main

import (
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

func main() {
	// Define the JSON schema
	schemaLoader := gojsonschema.NewStringLoader(`{
		"type": "object",
		"properties": {
			"name": { "type": "string" },
			"email": { "type": "string", "format": "email" },
			"age": { "type": "integer", "minimum": 18 }
		},
		"required": ["name", "email", "age"]
	}`)

	// JSON data to validate
	documentLoader := gojsonschema.NewStringLoader(`{
		"name": "John Doe",
		"email": "john@example.com",
		"age": 17
	}`)

	// Perform the validation
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if result.Valid() {
		fmt.Println("The document is valid")
	} else {
		fmt.Println("The document is not valid. See errors:")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}

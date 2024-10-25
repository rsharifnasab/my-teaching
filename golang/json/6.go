package main

import (
	"fmt"

	// new library: for validation
	"github.com/go-playground/validator/v10"
)

type User struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age" validate:"gte=18"`
}

func main() {
	validate := validator.New()

	user := User{
		Name:  "John Doe",
		Email: "john@example.com",
		Age:   17,
	}

	// Validate the user struct
	err := validate.Struct(user)
	if err != nil {
		// If there are validation errors, print them
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf("Validation failed on field '%s': %s\n", err.Field(), err.ActualTag())
		}
	} else {
		fmt.Println("User data is valid")
	}
}

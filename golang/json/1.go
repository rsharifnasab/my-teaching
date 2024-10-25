package main

import (
	"fmt"
	"reflect"
)

type Teacher struct {
	// public field
	ID string `json:"id" xml:"identifier"`

	// private field
	name string `json:"name"`

	// list with different tag
	Domains []string `json:"domains_abcd"`
}

func main() {
	teach := Teacher{}
	st := reflect.TypeOf(teach)
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		fmt.Printf("field name: %s | field tag : %s\n",
			field.Name,
			field.Tag.Get("json"),
		)
	}
}
